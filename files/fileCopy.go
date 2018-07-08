package files

import (
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
	"runtime"
)

type FileCopy struct {
	RealPath string
	Size     int64
	IsDir    bool
	Mode     os.FileMode
	Handle   *os.File
}

//copy file data
func ioCopy(srcHandle *os.File, targetPth string) (err error) {
	//log.Println("[D] ioCopy targetPth:", targetPth)
	parentDirectory := "."
	if runtime.GOOS == "windows" {
		if strings.Contains(targetPth, "\\") {
			index := strings.LastIndex(targetPth, "\\")
			parentDirectory = targetPth[0:index]
		}
	} else {
		if strings.Contains(targetPth, "/") {
			index := strings.LastIndex(targetPth, "/")
			parentDirectory = targetPth[0:index]
		}
	}
	if !IsPathExist(parentDirectory) {
		err := os.MkdirAll(parentDirectory, os.ModePerm)
		if err != nil {
			//log.Println("[D] ioCopy MkdirAll:", err)
			return err
		}
	}

	dstHandle, err := os.OpenFile(targetPth, os.O_CREATE|os.O_WRONLY, os.ModePerm)
	if err != nil {
		log.Println("[D] ioCopy parentDirectory:", parentDirectory)

	}

	defer srcHandle.Close()
	defer dstHandle.Close()

	_, err = io.Copy(dstHandle, srcHandle)
	//log.Println("[D] ioCopy Copy:", err)
	return err
}

// walk path, and send file to chan
// if want suffix file, just like go WalkPaths(mockFrom, ".txt", cpInfoChan)
// srcDir -> for make src dir
// suffix -> suffix for copy
// c -> walk chan, all most make as cpInfoChan := make(chan *FileCopy, 10)
func WalkPaths(srcDir, suffix string, c chan<- *FileCopy) {
	suffix = strings.ToUpper(suffix)

	filepath.Walk(srcDir, func(f string, fi os.FileInfo, err error) error { // walk path
		if err != nil {
			log.Println("[E] Walk", err)
			return err
		}

		fileInfo := &FileCopy{}
		if strings.HasSuffix(strings.ToUpper(fi.Name()), suffix) { // find out file
			if fh, err := os.OpenFile(f, os.O_RDONLY, os.ModePerm); err != nil {
				log.Println("[E] OpenFile", err)
				return err
			} else {
				fileInfo.Handle = fh
				fileInfo.RealPath, _ = filepath.Rel(srcDir, f) // relative path
				fileInfo.Size = fi.Size()
				fileInfo.IsDir = fi.IsDir()
				fileInfo.Mode = fi.Mode()
			}

			c <- fileInfo
		}
		return nil
	})
	//log.Println("[D]", fmt.Sprintf("len(c) %v", len(c)))
	close(c) // close chan when walk end
}

// write chan file to targetDir after WalkPaths(srcDir, suffix string, c chan<- *FileCopy)
// targetDir -> targetDir
// c -> WalkPaths return chan
// verbose -> is show log
func CopyPaths(targetDir string, c <-chan *FileCopy, verbose bool) error {
	if err := os.Chdir(targetDir); err != nil { // switch work path
		log.Fatalln("[F] Chdir", err)
	}
	for f := range c {
		if verbose {
			log.Println("[D] chan item:", f.RealPath)
		}
		if fi, err := os.Stat(f.RealPath); os.IsNotExist(err) { // target dir does not exist
			if f.IsDir {
				if err := os.MkdirAll(f.RealPath, f.Mode); err != nil {
					log.Println("[E] MkdirAll", err)
					return err
				}
				if verbose {
					log.Println("[D] folder CP:", f.RealPath)
				}
			} else {
				if err := ioCopy(f.Handle, f.RealPath); err != nil {
					log.Println("[E] ioCopy not exist ", err)
					return err
				} else {
					if verbose {
						log.Println("[D] file   CP:", f.RealPath)
					}
				}
			}
		} else if !f.IsDir { // The target exists and the source is not a directory

			if fi.IsDir() != f.IsDir { // Check file name is occupied by directory name conflict
				log.Println("[E]", "filename conflict:", f.RealPath)
			} else if fi.Size() != f.Size { // Rewrite only when the source and target sizes are inconsistent
				if err := ioCopy(f.Handle, f.RealPath); err != nil {
					log.Println("[E] ioCopy Diff file Size ", err)
					return err
				} else {
					if verbose {
						log.Println("[D] Diff file Size CP:", f.RealPath)
					}
				}
			}
		}
	}
	return nil
}

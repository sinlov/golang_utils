package files

import (
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
	"fmt"
)

type FileCopy struct {
	RelPath string
	Size    int64
	IsDir   bool
	Handle  *os.File
}

//copy file data
func ioCopy(srcHandle *os.File, dstPth string) (err error) {
	dstHandle, err := os.OpenFile(dstPth, os.O_CREATE|os.O_WRONLY, os.ModePerm)
	if err != nil {
		return err
	}

	defer srcHandle.Close()
	defer dstHandle.Close()

	_, err = io.Copy(dstHandle, srcHandle)
	return err
}

// walk path, and send file to chan
func WalkPaths(srcDir, suffix string, c chan<- *FileCopy) error {
	suffix = strings.ToUpper(suffix)

	filepath.Walk(srcDir, func(f string, fi os.FileInfo, err error) error { // walk path
		if err != nil {
			log.Println("[E]", err)
			return err
		}

		fileInfo := &FileCopy{}
		if strings.HasSuffix(strings.ToUpper(fi.Name()), suffix) { // find out file
			if fh, err := os.OpenFile(f, os.O_RDONLY, os.ModePerm); err != nil {
				log.Println("[E]", err)
				return err
			} else {
				fileInfo.Handle = fh
				fileInfo.RelPath, _ = filepath.Rel(srcDir, f) // relative path
				fileInfo.Size = fi.Size()
				fileInfo.IsDir = fi.IsDir()
			}

			c <- fileInfo
		}
		return nil
	})
	log.Println("[D]", fmt.Sprintf("len(c) %v", len(c)))
	close(c) // close chan when walk end
	return nil
}

//write chan file to targetDir
func CopyPaths(targetDir string, c <-chan *FileCopy, verbose bool) error {
	if err := os.Chdir(targetDir); err != nil { // switch work path
		log.Fatalln("[F]", err)
		return err
	}
	for f := range c {
		if fi, err := os.Stat(f.RelPath); os.IsNotExist(err) { // target dir does not exist
			if f.IsDir {
				if err := os.MkdirAll(f.RelPath, os.ModeDir); err != nil {
					log.Println("[E]", err)
					return err
				}
			} else {
				if err := ioCopy(f.Handle, f.RelPath); err != nil {
					log.Println("[E]", err)
					return err
				} else {
					if verbose {
						log.Println("[I] Dir CP:", f.RelPath)
						return nil
					}
				}
			}
		} else if !f.IsDir { // The target exists and the source is not a directory

			if fi.IsDir() != f.IsDir { // Check file name is occupied by directory name conflict
				log.Println("[E]", "filename conflict:", f.RelPath)
				return nil
			} else if fi.Size() != f.Size { // Rewrite only when the source and target sizes are inconsistent
				if err := ioCopy(f.Handle, f.RelPath); err != nil {
					log.Println("[E]", err)
				} else {
					if verbose {
						log.Println("[I] Diff Size CP:", f.RelPath)
					}
				}
			}
		}
	}
	return nil
}

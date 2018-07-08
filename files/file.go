package files

import (
	"io/ioutil"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func IsPathExist(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		if ! os.IsExist(err) {
			return false
		} else {
			return false
		}
	}
	return true
}

func IsFileExist(filePath string) bool {
	f, err := os.Open(filePath)
	if err != nil || os.IsNotExist(err) {
		return false
	}
	defer f.Close()
	return true
}

func ReadFileAsString(filePath string) (string, error) {
	b, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Printf("read file error %s\n", err)
	}
	s := string(b)
	return s, nil
}

// List DirPath all file, ignore sub folder
// dirPth -> for walk path
// suffix -> suffix want, if "" not check, ignore the case of suffix matching
func ListDirFiles(dirPath string, suffix string) (files []string, err error) {
	files = make([]string, 0, 10)

	dir, err := ioutil.ReadDir(dirPath)
	if err != nil {
		return nil, err
	}

	PthSep := string(os.PathSeparator)
	suffix = strings.ToUpper(suffix) // ignore the case of suffix matching

	for _, fi := range dir {
		if fi.IsDir() { // ignore dir
			continue
		}
		if strings.HasSuffix(strings.ToUpper(fi.Name()), suffix) { // suffix file
			files = append(files, dirPath+PthSep+fi.Name())
		}
	}

	return files, nil
}

// List DirPath all sub-folder, ignore sub Dirs
// dirPth -> for walk path
func ListSubDirs(dirPath string) (folder []string, err error) {
	folder = make([]string, 0, 10)

	dir, err := ioutil.ReadDir(dirPath)
	if err != nil {
		return nil, err
	}

	PthSep := string(os.PathSeparator)

	for _, fi := range dir {
		if !fi.IsDir() { // ignore file
			continue
		}
		folder = append(folder, dirPath+PthSep+fi.Name())
	}

	return folder, nil
}

// can get full file and in sub-folder file, ignore folder name
// dirPth -> for walk path
// suffix -> suffix want, if "" not check, ignore the case of suffix matching
func WalkDirFileAll(dirPath, suffix string) (files []string, err error) {
	files = make([]string, 0, 30)
	suffix = strings.ToUpper(suffix) // Ignore the case of suffix matching

	err = filepath.Walk(dirPath, func(name string, fi os.FileInfo, err error) error {
		//if err != nil { // ignore error?
		// return err
		//}

		if fi.IsDir() { // ignore folder
			return nil
		}

		if strings.HasSuffix(strings.ToUpper(fi.Name()), suffix) {
			files = append(files, name)
		}

		return nil
	})

	return files, err
}

// can get full folder and in sub-folder folder, ignore all file, and including itself
// dirPth -> for walk path
func WalkDirFolderAll(dirPath string) (folder []string, err error) {
	folder = make([]string, 0, 30)
	err = filepath.Walk(dirPath, func(name string, fi os.FileInfo, err error) error {
		//if err != nil { // ignore error?
		// return err
		//}

		if fi.IsDir() { // ignore file
			folder = append(folder, name)
		} else {
			return nil
		}

		return nil
	})

	return folder, err
}

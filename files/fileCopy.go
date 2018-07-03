package files

import (
	"os"
	"time"
	"path/filepath"
	"strings"
	"fmt"
	"github.com/Unknwon/com"
)

const (
	IsDirectory = iota
	IsRegular
	IsSymlink
)

type sysFile struct {
	fType  int
	fName  string
	fLink  string
	fSize  int64
	fMtime time.Time
	fPerm  os.FileMode
}

type F struct {
	files []*sysFile
}

//  _______ _________ _        _        _______
// (  ____ \\__   __/( (    /|( \      (  ___  )|\     /|
// | (    \/   ) (   |  \  ( || (      | (   ) || )   ( |
// | (_____    | |   |   \ | || |      | |   | || |   | |
// (_____  )   | |   | (\ \) || |      | |   | |( (   ) )
//       ) |   | |   | | \   || |      | |   | | \ \_/ /
// /\____) |___) (___| )  \  || (____/\| (___) |  \   /
// \_______)\_______/|/    )_)(_______/(_______)   \_/
//
func (f *F) visit(path string, file os.FileInfo, err error) error {
	if file == nil {
		return err
	}
	var tp int
	if file.IsDir() {
		tp = IsDirectory
	} else if (file.Mode() & os.ModeSymlink) > 0 {
		tp = IsSymlink
	} else {
		tp = IsRegular
	}
	inoFile := &sysFile{
		fName:  path,
		fType:  tp,
		fPerm:  file.Mode(),
		fMtime: file.ModTime(),
		fSize:  file.Size(),
	}
	f.files = append(f.files, inoFile)
	return nil
}

func (f *F) MergeFolderPath(sourceDir string, targetDir string) error {

	source := F{
		files: make([]*sysFile, 0),
	}
	err := filepath.Walk(sourceDir,
		func(path string, f os.FileInfo, err error) error {
			return source.visit(path, f, err)
		},
	)
	if err != nil {
		return err
	}
	target := F{
		files: make([]*sysFile, 0),
	}
	err = filepath.Walk(targetDir, func(path string, f os.FileInfo, err error) error {
		return target.visit(path, f, err)
	})
	if err != nil {
		return err
	}

	for _, v := range source.files {

		if com.IsFile(v.fName) == true {

			tmp1 := strings.Split(v.fName, "\\")
			sourceName := tmp1[len(tmp1)-1]
			for _, r := range target.files {
				if com.IsFile(r.fName) == true {
					tmp2 := strings.Split(r.fName, "\\")
					targetName := tmp2[len(tmp2)-1]
					if sourceName == targetName {
						fmt.Printf("the same file: %s\n", sourceName)
						com.Copy(v.fName, r.fName)
					}
				}
			}

		}
	}

	return nil
}

package cli

import (
	"os/exec"
	"os"
	"path/filepath"
	"strings"
	"github.com/sinlov/golang_utils/jstring"
)

func CommandPath() string {
	file, _ := exec.LookPath(os.Args[0])
	return file
}

func ParentDirectory(directory string) string {
	if IsSysWindows() {
		return jstring.SubString(directory, 0, strings.LastIndex(directory, "\\"))
	} else {
		return jstring.SubString(directory, 0, strings.LastIndex(directory, "/"))
	}

}

func CurrentDirectory() string {
	dir, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	if IsSysWindows() {
		return strings.Replace(dir, "\\", "/", -1)
	} else {
		return strings.Replace(dir, "/", "/", -1)
	}
}

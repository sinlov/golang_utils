package cli

import (
	"os/exec"
	"os"
	"path/filepath"
	"strings"
	"github.com/sinlov/golang_utils/jstring"
	"runtime"
	"os/user"
	"bytes"
	"errors"
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

// Home returns the home directory for the executing user.
//
// This uses an OS-specific method for discovering the home directory.
// An error is returned if a home directory cannot be detected.
func Home() (string, error) {
	userCur, err := user.Current()
	if nil == err {
		return userCur.HomeDir, nil
	}

	// cross compile support

	if "windows" == runtime.GOOS {
		return homeWindows()
	}

	// Unix-like system, so just assume Unix
	return homeUnix()
}

func homeUnix() (string, error) {
	// First prefer the HOME environmental variable
	if home := os.Getenv("HOME"); home != "" {
		return home, nil
	}

	// If that fails, try the shell
	var stdout bytes.Buffer
	cmd := exec.Command("sh", "-c", "eval echo ~$USER")
	cmd.Stdout = &stdout
	if err := cmd.Run(); err != nil {
		return "", err
	}

	result := strings.TrimSpace(stdout.String())
	if result == "" {
		return "", errors.New("blank output when reading home directory")
	}

	return result, nil
}

func homeWindows() (string, error) {
	drive := os.Getenv("HOMEDRIVE")
	path := os.Getenv("HOMEPATH")
	home := drive + path
	if drive == "" || path == "" {
		home = os.Getenv("USERPROFILE")
	}
	if home == "" {
		return "", errors.New("HOMEDRIVE, HOMEPATH, and USERPROFILE are blank")
	}

	return home, nil
}

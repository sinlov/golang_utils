package files

import (
	"io/ioutil"
	"fmt"
	"os"
)

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

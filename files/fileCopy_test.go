package files

import (
	"os"
	"strings"
	"path/filepath"
	"github.com/smartystreets/goconvey/convey"
	"testing"
	"github.com/Pallinder/go-randomdata"
	"github.com/Unknwon/com"
)

const (
	gitRepo = "golang_utils"
	gitUser = "sinlov"
	gitHost = "github.com"
)

func findOrCreateCodeBuildPath() string {
	goPathEnv := os.Getenv("GOPATH")
	goPathEnvS := strings.Split(goPathEnv, ":")
	for _, path := range goPathEnvS {
		codePath := filepath.Join(path, "src", gitHost, gitUser, gitRepo)
		if IsPathExist(codePath) {
			projectBuildPath := filepath.Join(codePath, "build")
			if IsFileExist(projectBuildPath) {
				return projectBuildPath
			} else {
				os.MkdirAll(projectBuildPath, os.ModePerm)
				return projectBuildPath
			}
		}
	}
	return ""
}

func cleanAndReCreatePath(path string) error {
	if IsPathExist(path) {
		err := os.RemoveAll(path)
		if err != nil {
			return err
		}
	}
	err := os.MkdirAll(path, os.ModePerm)
	if err != nil {
		return err
	}
	return nil
}

func mockRandomFiles(form, to string) (mockFrom, mockTo string) {
	buildPath := findOrCreateCodeBuildPath()
	mockFrom = filepath.Join(buildPath, form)
	cleanAndReCreatePath(mockFrom)
	mockTo = filepath.Join(buildPath, to)
	cleanAndReCreatePath(mockTo)
	return mockFrom, mockTo
}

func TestCopyPaths(t *testing.T) {
	convey.Convey("mock TestF_MergeFolderPath", t, func() {
		// mock
		copyFrom, copyTo := mockRandomFiles("mockFrom", "mockTo")
		randomAddressTextPath := filepath.Join(copyFrom, randomdata.RandStringRunes(10))
		randomAddress := randomdata.Address()
		err := com.WriteFile(randomAddressTextPath, []byte(randomAddress))
		if err != nil {
			t.Fatalf("error writefile %v, error is %v\n", randomAddressTextPath, err.Error())
		}
		// inner empty folder
		emptyDirPath := filepath.Join(copyFrom, "emptyDir")
		err = os.MkdirAll(emptyDirPath, os.ModePerm)
		if err != nil {
			t.Fatalf("error emptyDirPath %v, error is %v\n", emptyDirPath, err.Error())
		}
		// inner folder
		innerPath := filepath.Join(copyFrom, "inner", randomdata.RandStringRunes(10))
		err = com.WriteFile(innerPath, []byte(randomdata.IpV4Address()))
		if err != nil {
			t.Fatalf("error innerPath %v, error is %v\n", innerPath, err.Error())
		}

		convey.Convey("do TestF_MergeFolderPath", func() {
			// do
			cpInfoChan := make(chan *FileCopy, 10)
			go WalkPaths(copyFrom, "", cpInfoChan)
			copyErr := CopyPaths(copyTo, cpInfoChan, true)
			convey.Convey("verify TestF_MergeFolderPath", func() {
				// verify
				convey.So(copyErr, convey.ShouldEqual, nil)
			})
		})
	})
}

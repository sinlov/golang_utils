package files

import (
	"os"
	"strings"
	"path/filepath"
	"github.com/smartystreets/goconvey/convey"
	"testing"
	"github.com/Pallinder/go-randomdata"
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

func TestF_MergeFolderPath(t *testing.T) {
	convey.Convey("mock TestF_MergeFolderPath", t, func() {
		// mock
		buildPath := findOrCreateCodeBuildPath()
		copyFrom := filepath.Join(buildPath, "copyFrom")
		cleanAndReCreatePath(copyFrom)
		copyTo := filepath.Join(buildPath, "copyTo")
		cleanAndReCreatePath(copyTo)
		randomdata.Address()
		// TODO sinlov 2018/7/4 faker data
		convey.Convey("do TestF_MergeFolderPath", func() {
			// do
			convey.Convey("verify TestF_MergeFolderPath", func() {
				// verify
				convey.So("", convey.ShouldEqual, "")
			})
		})
	})
}

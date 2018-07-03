package files

import (
	"github.com/smartystreets/goconvey/convey"
	"testing"
	"github.com/sinlov/golang_utils/cli"
	"path/filepath"
)

func TestIsPathExist(t *testing.T) {
	convey.Convey("mock TestIsPathExist", t, func() {
		// mock
		home, err := cli.Home()
		if err != nil {
			t.Logf("find Home err %v\n", err)
		}
		notExistPath := filepath.Join(home, "zzzzz")
		convey.Convey("do TestIsPathExist", func() {
			// do
			exist := isPathExist(home)
			pathNotExist := isPathExist(notExistPath)
			convey.Convey("verify TestIsPathExist", func() {
				// verify
				convey.So(exist, convey.ShouldEqual, true)
				convey.So(pathNotExist, convey.ShouldEqual, false)
			})
		})
	})
}

func TestIsFileExist(t *testing.T) {
	convey.Convey("mock TestIsFileExist", t, func() {
		// mock
		homePath, err := cli.Home()
		if err != nil {
			t.Logf("find Home err %v\n", err)
		}
		userBashPath := filepath.Join(homePath, ".bashrc")
		userNotExistRCPath := filepath.Join(homePath, ".notExistrc")
		convey.Convey("do TestIsFileExist", func() {
			// do
			userBashExist := IsFileExist(userBashPath)
			userNotExistRC := IsFileExist(userNotExistRCPath)
			convey.Convey("verify TestIsFileExist", func() {
				// verify
				convey.So(userBashExist, convey.ShouldEqual, true)
				convey.So(userNotExistRC, convey.ShouldEqual, false)
			})
		})
	})
}

func TestReadFileAsString(t *testing.T) {
	convey.Convey("mock TestReadFileAsString", t, func() {
		// mock
		homePath, err := cli.Home()
		if err != nil {
			t.Logf("find Home err %v\n", err)
		}
		userBashRcPath := filepath.Join(homePath, ".bashrc")
		convey.Convey("do TestReadFileAsString", func() {
			// do
			fileAsString, err := ReadFileAsString(userBashRcPath)
			if err != nil {
				t.Logf("ReadFileAsString err %v\n", fileAsString)
			}
			convey.Convey("verify TestReadFileAsString", func() {
				// verify
				convey.So(fileAsString, convey.ShouldNotBeNil)
			})
		})
	})
}

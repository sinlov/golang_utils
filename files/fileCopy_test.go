package files

import (
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

const (
	gitRepo = "golang_utils"
	gitUser = "sinlov"
	gitHost = "github.com"
)

func TestCopyPaths(t *testing.T) {
	convey.Convey("mock TestF_MergeFolderPath", t, func() {
		// mock
		mockFrom, mockTo := mockRandomFiles("mockFrom", "mockTo")
		mockDirsAndFileInMockFolder(mockFrom, t)

		convey.Convey("do TestF_MergeFolderPath", func() {
			// do
			cpInfoChan := make(chan *FileCopy, 10)
			go WalkPaths(mockFrom, "", cpInfoChan)
			err := CopyPaths(mockTo, cpInfoChan, false)
			convey.Convey("verify TestF_MergeFolderPath", func() {
				// verify
				convey.So(err, convey.ShouldEqual, nil)
			})
		})
	})
}

func TestCopyPathsTxt(t *testing.T) {
	convey.Convey("mock TestCopyPathsTxt", t, func() {
		// mock
		mockFrom, mockTo := mockRandomFiles("mockFrom", "mockTo")
		mockDirsAndFileInMockFolder(mockFrom, t)

		convey.Convey("do TestCopyPathsTxt", func() {
			// do
			cpInfoChan := make(chan *FileCopy, 10)
			go WalkPaths(mockFrom, ".txt", cpInfoChan)
			err := CopyPaths(mockTo, cpInfoChan, false)
			convey.Convey("verify TestCopyPathsTxt", func() {
				// verify
				convey.So(err, convey.ShouldEqual, nil)

			})
		})
	})
}

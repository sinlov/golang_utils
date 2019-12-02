package files

import (
	"github.com/Pallinder/go-randomdata"
	"github.com/Unknwon/com"
	"github.com/sinlov/golang_utils/cli"
	"github.com/stretchr/testify/assert"
	"os"
	"path/filepath"
	"strings"
	"testing"
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

func mockDirsAndFileInMockFolder(mockFrom string, t *testing.T) {
	randomAddressTextPath := filepath.Join(mockFrom, randomdata.RandStringRunes(10))
	randomAddress := randomdata.Address()
	err := com.WriteFile(randomAddressTextPath, []byte(randomAddress))
	if err != nil {
		t.Logf("error randomAddressTextPath %v, error is %v\n", randomAddressTextPath, err.Error())
	}
	// inner empty folder
	emptyDirPath := filepath.Join(mockFrom, "emptyDir")
	err = os.MkdirAll(emptyDirPath, os.ModePerm)
	if err != nil {
		t.Fatalf("error emptyDirPath %v, error is %v\n", emptyDirPath, err.Error())
	}
	// double inner empty folder
	doubleEmptyDirPath := filepath.Join(emptyDirPath, "doubleDir")
	err = os.MkdirAll(doubleEmptyDirPath, os.ModePerm)
	if err != nil {
		t.Fatalf("error doubleEmptyDirPath %v, error is %v\n", doubleEmptyDirPath, err.Error())
	}

	// inner folder
	innerPath := filepath.Join(mockFrom, "inner", randomdata.RandStringRunes(10))
	err = com.WriteFile(innerPath, []byte(randomdata.IpV4Address()))
	if err != nil {
		t.Fatalf("error innerPath %v, error is %v\n", innerPath, err.Error())
	}
	// inner double folder
	doubleInnerPath := filepath.Join(mockFrom, "inner", "double")
	err = os.MkdirAll(doubleInnerPath, os.ModePerm)
	if err != nil {
		t.Fatalf("error doubleInnerPath %v, error is %v\n", doubleInnerPath, err.Error())
	}
	// doubler inner file path
	doubleInnerFilePath := filepath.Join(doubleInnerPath, randomdata.RandStringRunes(10))
	err = com.WriteFile(doubleInnerFilePath, []byte(randomdata.City()))
	if err != nil {
		t.Fatalf("error doubleInnerFilePath %v, error is %v\n", doubleInnerFilePath, err.Error())
	}

	// .txt file mock
	textPath := filepath.Join(mockFrom, randomdata.RandStringRunes(10)+".txt")
	err = com.WriteFile(textPath, []byte(randomdata.Address()))
	if err != nil {
		t.Logf("error textPath %v, error is %v\n", textPath, err.Error())
	}
	// inner .txt file mock
	textInnerPath := filepath.Join(mockFrom, "inner", randomdata.RandStringRunes(10)+".txt")
	err = com.WriteFile(textInnerPath, []byte(randomdata.Address()))
	if err != nil {
		t.Logf("error textInnerPath %v, error is %v\n", textInnerPath, err.Error())
	}
	// doubler inner .txt file path
	doubleInnerTxtFilePath := filepath.Join(doubleInnerPath, randomdata.RandStringRunes(10)+".txt")
	err = com.WriteFile(doubleInnerTxtFilePath, []byte(randomdata.City()))
	if err != nil {
		t.Fatalf("error doubleInnerTxtFilePath %v, error is %v\n", doubleInnerTxtFilePath, err.Error())
	}
}

func TestIsPathExist(t *testing.T) {
	// mock
	home, err := cli.Home()
	if err != nil {
		t.Logf("find Home err %v\n", err)
	}
	notExistPath := filepath.Join(home, "zzzzz")

	// do
	exist := IsPathExist(home)
	pathNotExist := IsPathExist(notExistPath)

	// verify
	assert.Equal(t, exist, true)
	assert.Equal(t, pathNotExist, false)

	//convey.Convey("mock TestIsPathExist", t, func() {
	//	// mock
	//	home, err := cli.Home()
	//	if err != nil {
	//		t.Logf("find Home err %v\n", err)
	//	}
	//	notExistPath := filepath.Join(home, "zzzzz")
	//	convey.Convey("do TestIsPathExist", func() {
	//		// do
	//		exist := IsPathExist(home)
	//		pathNotExist := IsPathExist(notExistPath)
	//		convey.Convey("verify TestIsPathExist", func() {
	//			// verify
	//			convey.So(exist, convey.ShouldEqual, true)
	//			convey.So(pathNotExist, convey.ShouldEqual, false)
	//		})
	//	})
	//})
}

func TestIsFileExist(t *testing.T) {
	// mock
	homePath, err := cli.Home()
	if err != nil {
		t.Logf("find Home err %v\n", err)
	}
	userBashPath := filepath.Join(homePath, ".bashrc")
	userNotExistRCPath := filepath.Join(homePath, ".notExistrc")

	// do
	userBashExist := IsFileExist(userBashPath)
	userNotExistRC := IsFileExist(userNotExistRCPath)

	// verify
	assert.Equal(t, userBashExist, true)
	assert.Equal(t, userNotExistRC, false)

	//convey.Convey("mock TestIsFileExist", t, func() {
	//	// mock
	//	homePath, err := cli.Home()
	//	if err != nil {
	//		t.Logf("find Home err %v\n", err)
	//	}
	//	userBashPath := filepath.Join(homePath, ".bashrc")
	//	userNotExistRCPath := filepath.Join(homePath, ".notExistrc")
	//	convey.Convey("do TestIsFileExist", func() {
	//		// do
	//		userBashExist := IsFileExist(userBashPath)
	//		userNotExistRC := IsFileExist(userNotExistRCPath)
	//		convey.Convey("verify TestIsFileExist", func() {
	//			// verify
	//			convey.So(userBashExist, convey.ShouldEqual, true)
	//			convey.So(userNotExistRC, convey.ShouldEqual, false)
	//		})
	//	})
	//})
}

func TestReadFileAsString(t *testing.T) {
	// mock
	homePath, err := cli.Home()
	if err != nil {
		t.Logf("find Home err %v\n", err)
	}
	userBashRcPath := filepath.Join(homePath, ".bashrc")

	// do
	fileAsString, err := ReadFileAsString(userBashRcPath)
	if err != nil {
		t.Logf("ReadFileAsString err %v\n", fileAsString)
	}

	// verify
	assert.NotEqual(t, nil, fileAsString)

	//convey.Convey("mock TestReadFileAsString", t, func() {
	//	// mock
	//	homePath, err := cli.Home()
	//	if err != nil {
	//		t.Logf("find Home err %v\n", err)
	//	}
	//	userBashRcPath := filepath.Join(homePath, ".bashrc")
	//	convey.Convey("do TestReadFileAsString", func() {
	//		// do
	//		fileAsString, err := ReadFileAsString(userBashRcPath)
	//		if err != nil {
	//			t.Logf("ReadFileAsString err %v\n", fileAsString)
	//		}
	//		convey.Convey("verify TestReadFileAsString", func() {
	//			// verify
	//			convey.So(fileAsString, convey.ShouldNotBeNil)
	//		})
	//	})
	//})
}

func TestListSubDirs(t *testing.T) {
	// mock
	mockFrom, mockTo := mockRandomFiles("mockFrom", "mockTo")
	mockDirsAndFileInMockFolder(mockFrom, t)
	t.Logf("mockFrom %v, mockTo %v\n", mockFrom, mockTo)

	// do
	files, err := ListSubDirs(mockFrom)
	if err != nil {
		t.Fatalf("ListSubDirs %v, error is %v\n", mockFrom, err.Error())
	}

	// verify
	for _, walk := range files {
		assert.NotEqual(t, nil, walk)
		t.Logf("ListSubDirs item %v", walk)
	}

	//convey.Convey("mock TestListSubDirs", t, func() {
	//	// mock
	//	mockFrom, mockTo := mockRandomFiles("mockFrom", "mockTo")
	//	mockDirsAndFileInMockFolder(mockFrom, t)
	//	t.Logf("mockFrom %v, mockTo %v\n", mockFrom, mockTo)
	//
	//	convey.Convey("do TestListSubDirs", func() {
	//		// do
	//		files, err := ListSubDirs(mockFrom)
	//		if err != nil {
	//			t.Fatalf("ListSubDirs %v, error is %v\n", mockFrom, err.Error())
	//		}
	//
	//		convey.Convey("verify TestListSubDirs", func() {
	//			// verify
	//			for _, walk := range files {
	//				t.Logf("ListSubDirs item %v", walk)
	//			}
	//			convey.So("", convey.ShouldEqual, "")
	//
	//		})
	//	})
	//})
}

func TestListDirFiles(t *testing.T) {
	// mock
	mockFrom, mockTo := mockRandomFiles("mockFrom", "mockTo")
	mockDirsAndFileInMockFolder(mockFrom, t)
	t.Logf("mockFrom %v, mockTo %v\n", mockFrom, mockTo)

	// do
	dirFiles, err := ListDirFiles(mockFrom, "")
	if err != nil {
		t.Fatalf("ListDirFiles %v, error is %v\n", mockFrom, err.Error())
	}

	dirTxtFiles, err := ListDirFiles(mockFrom, ".txt")
	if err != nil {
		t.Fatalf("ListDirFiles %v, error is %v\n", mockFrom, err.Error())
	}

	// verify
	for _, walk := range dirFiles {
		assert.NotEqual(t, nil, walk)
		t.Logf("ListDirFiles suffix=[ %v ] item %v", "", walk)
	}
	for _, walk := range dirTxtFiles {
		assert.NotEqual(t, nil, walk)
		t.Logf("ListDirFiles suffix=[ %v ] item %v", ".txt", walk)
	}

	//convey.Convey("mock TestListDirFiles", t, func() {
	//	// mock
	//	mockFrom, mockTo := mockRandomFiles("mockFrom", "mockTo")
	//	mockDirsAndFileInMockFolder(mockFrom, t)
	//	t.Logf("mockFrom %v, mockTo %v\n", mockFrom, mockTo)
	//
	//	convey.Convey("do TestListDirFiles", func() {
	//		// do
	//		dirFiles, err := ListDirFiles(mockFrom, "")
	//		if err != nil {
	//			t.Fatalf("ListDirFiles %v, error is %v\n", mockFrom, err.Error())
	//		}
	//
	//		dirTxtFiles, err := ListDirFiles(mockFrom, ".txt")
	//		if err != nil {
	//			t.Fatalf("ListDirFiles %v, error is %v\n", mockFrom, err.Error())
	//		}
	//
	//		convey.Convey("verify TestListDirFiles", func() {
	//			// verify
	//			for _, walk := range dirFiles {
	//				t.Logf("ListDirFiles suffix=[ %v ] item %v", "", walk)
	//			}
	//			for _, walk := range dirTxtFiles {
	//				t.Logf("ListDirFiles suffix=[ %v ] item %v", ".txt", walk)
	//			}
	//			convey.So("", convey.ShouldEqual, "")
	//		})
	//	})
	//})
}

func TestWalkDirFileAll(t *testing.T) {
	// mock
	mockFrom, mockTo := mockRandomFiles("mockFrom", "mockTo")
	mockDirsAndFileInMockFolder(mockFrom, t)

	t.Logf("mockFrom %v, mockTo %v\n", mockFrom, mockTo)

	// do
	walkDir, err := WalkDirFileAll(mockFrom, "")
	if err != nil {
		t.Fatalf("WalkDirFileAll error %v", err.Error())
	}
	txtWalkDir, err := WalkDirFileAll(mockFrom, ".txt")
	if err != nil {
		t.Fatalf("WalkDirFileAll error %v", err.Error())
	}

	// verify
	for _, walk := range walkDir {
		assert.NotEqual(t, nil, walk)
		t.Logf("WalkDirFileAll suffix=[ %v ] item %v", "", walk)
	}

	for _, walk := range txtWalkDir {
		assert.NotEqual(t, nil, walk)
		t.Logf("WalkDirFileAll suffix=[ %v ] item %v", ".txt", walk)
	}

	//convey.Convey("mock TestWalkDir", t, func() {
	//	// mock
	//	mockFrom, mockTo := mockRandomFiles("mockFrom", "mockTo")
	//	mockDirsAndFileInMockFolder(mockFrom, t)
	//
	//	t.Logf("mockFrom %v, mockTo %v\n", mockFrom, mockTo)
	//
	//	convey.Convey("do TestWalkDir", func() {
	//		// do
	//		walkDir, err := WalkDirFileAll(mockFrom, "")
	//		if err != nil {
	//			t.Fatalf("WalkDirFileAll error %v", err.Error())
	//		}
	//		txtWalkDir, err := WalkDirFileAll(mockFrom, ".txt")
	//		if err != nil {
	//			t.Fatalf("WalkDirFileAll error %v", err.Error())
	//		}
	//		convey.Convey("verify TestWalkDir", func() {
	//			// verify
	//			for _, walk := range walkDir {
	//				t.Logf("WalkDirFileAll suffix=[ %v ] item %v", "", walk)
	//			}
	//
	//			for _, walk := range txtWalkDir {
	//				t.Logf("WalkDirFileAll suffix=[ %v ] item %v", ".txt", walk)
	//			}
	//
	//			convey.So("", convey.ShouldEqual, "")
	//
	//		})
	//	})
	//})
}

func TestWalkDirFolderAll(t *testing.T) {
	// mock
	mockFrom, mockTo := mockRandomFiles("mockFrom", "mockTo")
	mockDirsAndFileInMockFolder(mockFrom, t)

	t.Logf("mockFrom %v, mockTo %v\n", mockFrom, mockTo)

	// do
	folders, err := WalkDirFolderAll(mockFrom)
	if err != nil {
		t.Fatalf("WalkDirFolderAll error %v", err.Error())
	}

	// verify
	for _, walk := range folders {
		assert.NotEqual(t, nil, walk)
		t.Logf("WalkDirFolderAll item %v", walk)
	}

	//convey.Convey("mock TestWalkDirFolderAll", t, func() {
	//	// mock
	//	mockFrom, mockTo := mockRandomFiles("mockFrom", "mockTo")
	//	mockDirsAndFileInMockFolder(mockFrom, t)
	//
	//	t.Logf("mockFrom %v, mockTo %v\n", mockFrom, mockTo)
	//
	//	convey.Convey("do TestWalkDirFolderAll", func() {
	//		// do
	//		folders, err := WalkDirFolderAll(mockFrom)
	//		if err != nil {
	//			t.Fatalf("WalkDirFolderAll error %v", err.Error())
	//		}
	//		convey.Convey("verify TestWalkDirFolderAll", func() {
	//			// verify
	//			for _, walk := range folders {
	//				t.Logf("WalkDirFolderAll item %v", walk)
	//			}
	//			convey.So("", convey.ShouldEqual, "")
	//
	//		})
	//	})
	//})
}

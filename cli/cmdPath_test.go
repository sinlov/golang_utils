package cli

import (
	"testing"
	"github.com/smartystreets/goconvey/convey"
	"os/exec"
	"os"
	"os/user"
)

func TestCommandPath(t *testing.T) {
	convey.Convey("mock TestCommandPath", t, func() {
		// mock
		want, _ := exec.LookPath(os.Args[0])
		convey.Convey("do TestCommandPath", func() {
			// do
			got := CommandPath()
			convey.Convey("verify TestCommandPath", func() {
				// verify
				convey.So(got, convey.ShouldEqual, want)
			})
		})
	})
}

func TestCurrentDirectory(t *testing.T) {
	convey.Convey("mock TestCurrentDirectory", t, func() {
		// mock
		want := ""
		convey.Convey("do TestCurrentDirectory", func() {
			// do
			got := CurrentDirectory();
			convey.Convey("verify TestCurrentDirectory", func() {
				// verify
				t.Skipf("%s %s got %s", t.Name(), "cuz got Directory not for test just skip", got)
				convey.So(got, convey.ShouldEqual, want)
			})
		})
	})
}

func TestParentDirectory(t *testing.T) {
	convey.Convey("mock TestParentDirectory", t, func() {
		// mock
		directory := CurrentDirectory()
		want := ""
		convey.Convey("do TestParentDirectory", func() {
			// do
			got := ParentDirectory(directory)
			convey.Convey("verify TestParentDirectory", func() {
				// verify
				t.Skipf("%s %s got %s", t.Name(), "cuz got Directory not for test just skip", got)
				convey.So(got, convey.ShouldEqual, want)
			})
		})
	})
}

func TestHome(t *testing.T) {
	convey.Convey("mock TestHome", t, func() {
		// mock
		userCur, _ := user.Current()
		want := userCur.HomeDir
		convey.Convey("do TestHome", func() {
			// do
			got, err := Home()
			convey.Convey("verify TestHome", func() {
				// verify
				if err != nil {
					t.Errorf("%s got User Home error, %s", t.Name(), err)
				} else {
					convey.So(got, convey.ShouldEqual, want)
				}

			})
		})
	})
}

func TestHomeUnix(t *testing.T) {
	convey.Convey("mock TestHomeUnix", t, func() {
		// mock
		userCur, _ := user.Current()
		want := userCur.HomeDir
		convey.Convey("do TestHomeUnix", func() {
			// do
			got, err := homeUnix()
			convey.Convey("verify TestHomeUnix", func() {
				// verify
				if err != nil {
					t.Errorf("%s got User Home by homeUnix() error, %s", t.Name(), err)
				} else {
					if IsSysWindows() {
						t.Skipf("%s %s", t.Name(), "now system windows so pass homeUnix!")
					} else {
						convey.So(got, convey.ShouldEqual, want)
					}
				}
			})
		})
	})
}

func TestHomeWindows(t *testing.T) {
	convey.Convey("mock TestHomeWindows", t, func() {
		// mock
		userCur, _ := user.Current()
		want := userCur.HomeDir
		convey.Convey("do TestHomeWindows", func() {
			// do
			if !IsSysWindows() {
				t.Skipf("%s %s", t.Name(), "not windows system so pass!")
			}
			got, err := homeWindows()
			convey.Convey("verify TestHomeWindows", func() {
				// verify
				if err != nil {
					t.Errorf("%s got User Home by homeWindows() error, %s", t.Name(), err)
				} else {
					convey.So(got, convey.ShouldEqual, want)
				}
			})
		})
	})
}

package cli

import (
	"testing"
	"github.com/smartystreets/goconvey/convey"
	"runtime"
)

func TestIsSysWindows(t *testing.T) {
	convey.Convey("mock TestIsSysWindows", t, func() {
		// mock
		want := runtime.GOOS == "windows"
		convey.Convey("do TestIsSysWindows", func() {
			// do
			got := IsSysWindows()
			convey.Convey("verify TestIsSysWindows", func() {
				// verify
				convey.So(got, convey.ShouldEqual, want)
			})
		})
	})
}

func TestCmdTeaRun(t *testing.T) {
	convey.Convey("mock TestCmdTeaRun", t, func() {
		// mock
		type args struct {
			chartSet string
			isPrint  bool
			cmd      []string
		}
		data := struct {
			name     string
			args     args
			want     bool
			exitInfo string
		}{
			name: "git status",
			args: args{
				chartSet: "",
				isPrint:  true,
				cmd:      []string{"git", "status "},
			},
			want:     true,
			exitInfo: "exit status 0",
		}
		convey.Convey("do TestCmdTeaRun", func() {
			// do
			cmdTea := new(CmdTea)
			cmdTea.CmdTeaInit(data.args.chartSet, data.args.isPrint, data.args.cmd...)
			hasSuccess, tea := cmdTea.CmdTeaRun()

			convey.Convey("verify TestCmdTeaRun", func() {
				// verify
				if tea.ExitState != data.exitInfo {
					t.Errorf("CmdTeaRun() got = %v, want %v", tea.ExitState, data.exitInfo)
				}
				if !hasSuccess {
					if tea.ErrorInfo != nil {
						//t.Logf("tea result\n %s\n", tea)
						t.Logf("tea CmdStrings: \n%v\n", tea.CmdStrings)
						t.Logf("tea ShellPath: \n%v\n", tea.ShellPath)
						t.Logf("tea IsSuccess: \n%v\n", tea.IsSuccess)
						t.Logf("tea ErrorInfo: \n%v\n", tea.ErrorInfo)
						t.Logf("tea Pid: \n%v\n", tea.Pid)
						t.Logf("tea Out: \n%v\n", tea.Out)
						t.Logf("tea Err: \n%v\n", tea.Err)
						t.Logf("tea ExitState: \n%v\n", tea.ExitState)
					}

				} else {
					//t.Logf("tea result\n %s\n", tea)
					t.Logf("tea IsSuccess: \n%v\n", tea.IsSuccess)
					t.Logf("tea ErrorInfo: \n%v\n", tea.ErrorInfo)
					t.Logf("tea Pid: \n%v\n", tea.Pid)
					t.Logf("tea Out: \n%v\n", tea.Out)
					t.Logf("tea Err: \n%v\n", tea.Err)
					t.Logf("tea ExitState: \n%v\n", tea.ExitState)
				}
				convey.So(hasSuccess, convey.ShouldEqual, data.want)
			})
		})
	})
}

func TestCmdTeaRunError(t *testing.T) {
	convey.Convey("mock TestCmdTeaRunError", t, func() {
		// mock
		type args struct {
			chartSet string
			isPrint  bool
			cmd      []string
		}
		data := struct {
			name     string
			args     args
			want     bool
			exitInfo string
		}{
			name: "cat this.log",
			args: args{
				chartSet: "",
				isPrint:  true,
				cmd:      []string{"cat", "this.log"},
			},
			want:     false,
			exitInfo: "exit status 1",
		}

		convey.Convey("do TestCmdTeaRunError", func() {
			// do
			cmdTea := new(CmdTea)
			cmdTea.CmdTeaInit(data.args.chartSet, data.args.isPrint, data.args.cmd...)
			hasSuccess, tea := cmdTea.CmdTeaRun()
			convey.Convey("verify TestCmdTeaRunError", func() {
				// verify
				// verify
				if tea.ExitState != data.exitInfo {
					t.Errorf("CmdTeaRun() got = %v, want %v", tea.ExitState, data.exitInfo)
				}
				if !hasSuccess {
					if tea.ErrorInfo != nil {
						//t.Logf("tea result\n %s\n", tea)
						t.Logf("tea CmdStrings: \n%v\n", tea.CmdStrings)
						t.Logf("tea ShellPath: \n%v\n", tea.ShellPath)
						t.Logf("tea IsSuccess: \n%v\n", tea.IsSuccess)
						t.Logf("tea ErrorInfo: \n%v\n", tea.ErrorInfo)
						t.Logf("tea Pid: \n%v\n", tea.Pid)
						t.Logf("tea Out: \n%v\n", tea.Out)
						t.Logf("tea Err: \n%v\n", tea.Err)
						t.Logf("tea ExitState: \n%v\n", tea.ExitState)
					}

				} else {
					//t.Logf("tea result\n %s\n", tea)
					t.Logf("tea IsSuccess: \n%v\n", tea.IsSuccess)
					t.Logf("tea ErrorInfo: \n%v\n", tea.ErrorInfo)
					t.Logf("tea Pid: \n%v\n", tea.Pid)
					t.Logf("tea Out: \n%v\n", tea.Out)
					t.Logf("tea Err: \n%v\n", tea.Err)
					t.Logf("tea ExitState: \n%v\n", tea.ExitState)
				}
				convey.So(hasSuccess, convey.ShouldEqual, data.want)
			})
		})
	})
}

func TestCmdTeaExec(t *testing.T) {
	convey.Convey("mock TestCmdTeaExec", t, func() {
		// mock
		type args struct {
			chartSet string
			cmd      []string
		}
		data := struct {
			name      string
			args      args
			want      bool
			code      int
			outStatus string
			outInfo   string
		}{
			name: "git status",
			args: args{
				chartSet: "",
				cmd:      []string{"git", "status"},
			},
			want:      true,
			code:      0,
			outStatus: "exit status 0",
		}
		convey.Convey("do TestCmdTeaExec", func() {
			// do
			isSuccess, code, out, outInfo := CmdExec(data.args.chartSet, data.args.cmd...)
			convey.Convey("verify TestCmdTeaExec", func() {
				// verify
				t.Logf("isSuccess=> %v\n code=> %v\noutStatus=> %v\noutInfo=> %v", isSuccess, code, out, outInfo)
				convey.So(isSuccess, convey.ShouldEqual, data.want)
				convey.So(code, convey.ShouldEqual, data.code)
				convey.So(out, convey.ShouldEqual, data.outStatus)
				convey.So(outInfo, convey.ShouldNotBeNil)
			})
		})
	})
}

func TestCmdTeaExecError(t *testing.T) {
	convey.Convey("mock TestCmdTeaExecError", t, func() {
		// mock
		type args struct {
			chartSet string
			cmd      []string
		}
		data := struct {
			name      string
			args      args
			want      bool
			code      int
			outStatus string
			outInfo   string
		}{
			name: "cat this.log",
			args: args{
				chartSet: "",
				cmd:      []string{"cat", "this.log"},
			},
			want:      false,
			code:      0,
			outStatus: "exit status 1",
			outInfo:   "cat: this.log: No such file or directory\n",
		}
		convey.Convey("do TestCmdTeaExecError", func() {
			// do
			isSuccess, code, out, outInfo := CmdExec(data.args.chartSet, data.args.cmd...)
			convey.Convey("verify TestCmdTeaExecError", func() {
				// verify
				t.Logf("isSuccess=> %v\ncode=> %v\noutStatus=> %v\noutInfo=> %v", isSuccess, code, out, outInfo)
				convey.So(isSuccess, convey.ShouldEqual, data.want)
				convey.So(code, convey.ShouldEqual, data.code)
				convey.So(out, convey.ShouldEqual, data.outStatus)
				convey.So(outInfo, convey.ShouldEqual, data.outInfo)
			})
		})
	})
}

func TestCmdRun(t *testing.T) {
	convey.Convey("mock TestCmdRun", t, func() {
		// mock
		type args struct {
			chartSet string
			cmd      []string
		}
		data := struct {
			name string
			args args
			want bool
		}{
			name: "git status",
			args: args{
				chartSet: "",
				cmd:      []string{"git", "status"},
			},
			want: true,
		}
		convey.Convey("do TestCmdRun", func() {
			// do
			isSuccess, err := CmdRun(data.args.chartSet, data.args.cmd...)
			convey.Convey("verify TestCmdRun", func() {
				// verify
				if err != nil {
					t.Errorf("%s test run error %s", t.Name(), err)
				} else {
					convey.So(isSuccess, convey.ShouldEqual, data.want)
				}
			})
		})
	})
}

func TestCmdRunError(t *testing.T) {
	convey.Convey("mock TestCmdRunError", t, func() {
		// mock
		type args struct {
			chartSet string
			cmd      []string
		}
		data := struct {
			name string
			args args
			want bool
		}{
			name: "cat this.log",
			args: args{
				chartSet: "",
				cmd:      []string{"cat", "this.log"},
			},
			want: false,
		}
		convey.Convey("do TestCmdRunError", func() {
			// do
			isSuccess, err := CmdRun(data.args.chartSet, data.args.cmd...)
			convey.Convey("verify TestCmdRunError", func() {
				// verify
				convey.So(isSuccess, convey.ShouldEqual, data.want)
				t.Logf("isSuccess %v\nError info %s", isSuccess, err)
			})
		})
	})
}

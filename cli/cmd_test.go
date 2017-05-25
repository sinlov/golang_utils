package cli

import (
	"testing"
	"fmt"
)

func TestIsSysWindows(t *testing.T) {
	tests := []struct {
		name string
		want bool
	}{
		{
			name: "mac test",
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsSysWindows(); got != tt.want {
				t.Errorf("IsSysWindows() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCmdExec(t *testing.T) {
	type args struct {
		chartSet string
		cmd      []string
	}
	tests := []struct {
		name  string
		args  args
		want  bool
		want1 int
		want2 string
		want3 string
	}{
		{
			name: "git status",
			args: args{
				chartSet: "",
				cmd:      []string{"git", "status"},
			},
			want:  true,
			want2: "exit status 0",
		},
		{
			name: "cat this.log",
			args: args{
				chartSet: "",
				cmd:      []string{"cat", "this.log"},
			},
			want:  false,
			want2: "exit status 1",
			want3: "cat: this.log: No such file or directory",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, got2, got3 := CmdExec(tt.args.chartSet, tt.args.cmd...)
			if got != tt.want {
				t.Errorf("CmdExec() got = %v, want %v", got, tt.want)
			}
			if got2 != tt.want2 {
				t.Errorf("CmdExec() got2 = %v, want %v", got2, tt.want2)
			}
			//if got3 != tt.want3 {
			//  t.Errorf("CmdExec() got3 = %v, want %v", got3, tt.want3)
			//}
			fmt.Printf("Success %v\nProcess PID: %v\nProcess Status: %v\nOut: %v", got, got1, got2, got3)
			//if got2 != tt.want2 {
			//  t.Errorf("CmdExec() got2 = %v, want %v", got2, tt.want2)
			//}
		})
	}
}

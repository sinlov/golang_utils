package cli

import (
	"testing"
	"fmt"
)

func TestCmdExec(t *testing.T) {
	type args struct {
		cmd    string
		system string
	}
	tests := []struct {
		name  string
		args  args
		want  bool
		want1 string
		want2 string
	}{
		{
			name: "git status",
			args: args{
				cmd: "git status",
			},
			want:  true,
			want1: "exit status 0",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, got2 := CmdExec(tt.args.cmd, tt.args.system)
			if got != tt.want {
				t.Errorf("CmdExec() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("CmdExec() got1 = %v, want %v", got1, tt.want1)
			}
			fmt.Printf("Success %v\nProcess Status: %v\nOut: %v", got, got1, got2)
			//if got2 != tt.want2 {
			//	t.Errorf("CmdExec() got2 = %v, want %v", got2, tt.want2)
			//}
		})
	}
}

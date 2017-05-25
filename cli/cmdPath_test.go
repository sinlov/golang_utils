package cli

import (
	"testing"
	"fmt"
)

func TestCommandPath(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{
			name: "CommandPath",
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CommandPath(); got != tt.want {
				fmt.Printf("got: %v\n", got)
				//t.Errorf("CommandPath() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParentDirectory(t *testing.T) {
	type args struct {
		directory string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "ParentDirectory",
			args: args{
				directory: CurrentDirectory(),
			},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ParentDirectory(tt.args.directory); got != tt.want {
				fmt.Printf("got: %v\n", got)
				//t.Errorf("ParentDirectory() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCurrentDirectory(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{
			name: "CurrentDirectory",
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CurrentDirectory(); got != tt.want {
				fmt.Printf("got: %v\n", got)
				//t.Errorf("CurrentDirectory() = %v, want %v", got, tt.want)
			}
		})
	}
}

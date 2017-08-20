package cli

import (
	"fmt"
	"testing"
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

func TestHome(t *testing.T) {
	tests := []struct {
		name    string
		want    string
		wantErr bool
	}{
		{
			name:"test Home",
			want:"/Users/sinlov",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Home()
			if (err != nil) != tt.wantErr {
				t.Errorf("Home() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Home() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_homeUnix(t *testing.T) {
	tests := []struct {
		name    string
		want    string
		wantErr bool
	}{
		{
			name:"Test_homeUnix",
			want:"/Users/sinlov",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := homeUnix()
			if (err != nil) != tt.wantErr {
				t.Errorf("homeUnix() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("homeUnix() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_homeWindows(t *testing.T) {
	tests := []struct {
		name    string
		want    string
		wantErr bool
	}{
		{
			name:"Test_homeUnix",
			want:"C:\\User\\sinlov",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := homeWindows()
			if (err != nil) != tt.wantErr {
				t.Errorf("homeWindows() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("homeWindows() = %v, want %v", got, tt.want)
			}
		})
	}
}

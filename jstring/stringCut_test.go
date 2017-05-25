package jstring

import (
	"testing"
	"fmt"
)

func TestSubString(t *testing.T) {
	type args struct {
		s      string
		pos    int
		length int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "base",
			args: args{
				s:      "/Users/sinlov/goPath/src/github.com/sinlov/golang_utils/jstring",
				pos:    0,
				length: 10,
			},
			want: "/Users/sin",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SubString(tt.args.s, tt.args.pos, tt.args.length); got != tt.want {
				fmt.Printf("got: %v\n", got)
				t.Errorf("SubString() = %v, want %v", got, tt.want)
			}
		})
	}
}

package format

import (
	"testing"
	"fmt"
	"time"
)

func TestLayoutNowTime(t *testing.T) {
	type args struct {
		layout string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "base test",
			args: args{
				layout: "2006-01-02 15:04:05 .000",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			got := LayoutNowTime(tt.args.layout)
			fmt.Printf("got : %v local %v\n", got, time.Now().Local().String())
			if got != tt.want {
				t.Errorf("LayoutNowTime() = %v, want %v", got, tt.want)
			}

		})
	}
}

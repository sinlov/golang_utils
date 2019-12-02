package format

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
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
			want: time.Now().Format("2006-01-02 15:04:05 .000"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			got := LayoutNowTime(tt.args.layout)
			fmt.Printf("got : %v local %v\n", got, time.Now().Local().String())
			assert.Equal(t, tt.want, got)
		})
	}
}

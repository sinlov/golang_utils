package jstring

import "testing"

func TestStringStartWith(t *testing.T) {
	type args struct {
		s     string
		start string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "StringStartWith match test",
			args: args{
				s:     "123456",
				start: "1",
			},
			want: true,
		},
		{
			name: "StringStartWith match more",
			args: args{
				s:     "123456",
				start: "123",
			},
			want: true,
		},
		{
			name: "StringStartWith not match test",
			args: args{
				s:     "123456",
				start: "2",
			},
			want: false,
		},
		{
			name: "StringStartWith s is empty",
			args: args{
				s:     "",
				start: "my",
			},
			want: false,
		},
		{
			name: "StringStartWith start is empty",
			args: args{
				s:     "start",
				start: "",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := StringStartWith(tt.args.s, tt.args.start); got != tt.want {
				t.Errorf("StringStartWith() = %v, want %v, s = [ %s ] => start = [ %s ]", got, tt.want, tt.args.s, tt.args.start)
			}
		})
	}
}

func TestStringEndWith(t *testing.T) {
	type args struct {
		s   string
		end string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "TestStringEndWith match test",
			args: args{
				s:   "123456",
				end: "6",
			},
			want: true,
		},
		{
			name: "TestStringEndWith match more",
			args: args{
				s:   "123456",
				end: "456",
			},
			want: true,
		},
		{
			name: "TestStringEndWith not match test",
			args: args{
				s:   "123456",
				end: "2",
			},
			want: false,
		},
		{
			name: "TestStringEndWith s is empty",
			args: args{
				s:   "",
				end: "my",
			},
			want: false,
		},
		{
			name: "TestStringEndWith end is empty",
			args: args{
				s:   "end",
				end: "",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := StringEndWith(tt.args.s, tt.args.end); got != tt.want {
				t.Errorf("StringEndWith() = %v, want %v, s = [ %s ] => end = [ %s ]", got, tt.want, tt.args.s, tt.args.end)
			}
		})
	}
}

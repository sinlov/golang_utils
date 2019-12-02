package randomplus

import (
	"testing"
)

func TestTimeSeed(t *testing.T) {
	random, sed := TimeSeed(999)
	t.Logf("random %v, time sed %v", random, sed.Unix())
}

func TestPositive(t *testing.T) {
	type args struct {
		size int
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{
			name: "postive",
			args: args{
				size: 7,
			},
		},
	}
	for _, tt := range tests {
		got, err := Positive(tt.args.size)
		if err != nil {
			t.Errorf("%q. Postive() error = %v", tt.name, err)

		}
		t.Logf("%q. Postive got %v", tt.name, got)
		//if (err != nil) != tt.wantErr {
		//	t.Errorf("%q. Positive() error = %v, wantErr %v", tt.name, err, tt.wantErr)
		//	continue
		//}
		//if got != tt.want {
		//	t.Errorf("%q. Positive() = %v, want %v", tt.name, got, tt.want)
		//}
	}
}

func TestPositiveMore(t *testing.T) {
	testTimes := 1000
	for i := 1; i < testTimes; i++ {
		got, err := Positive(6)
		if err != nil {
			t.Errorf("Postive() error = %v", err)

		}
		if got < 99999 && got > 999999 {
			t.Errorf("Postive got number error %v", got)
		} else {
			t.Logf("Postive got %v", got)
		}
	}
}

func BenchmarkPositive(b *testing.B) {
	testTimes := 10
	for i := 1; i < testTimes; i++ {
		got, err := Positive(8)
		if err != nil {
			b.Errorf("Postive() error = %v", err)

		}
		if got < 9999999 && got > 99999999 {
			b.Errorf("Postive got number error %v", got)
		} else {
			//b.Logf("Postive got %v", got)
		}
	}
}

func TestPositiveNegative(t *testing.T) {
	type args struct {
		size int
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{
			name: "PositiveNegative",
			args: args{
				size: 7,
			},
		},
	}
	for _, tt := range tests {
		got, err := PositiveNegative(tt.args.size)
		if err != nil {
			t.Errorf("%q. PositiveNegative() error = %v", tt.name, err)

		}
		t.Logf("%q. PositiveNegative got %v", tt.name, got)
	}
}

func TestPositiveNegativeMore(t *testing.T) {
	testTimes := 1000
	for i := 1; i < testTimes; i++ {
		got, err := PositiveNegative(8)
		if err != nil {
			t.Errorf("PositiveNegative() error = %v", err)

		}
		if got < 0 {
			t.Logf("change PositiveNegative %v", got)
			got = got * -1
		}
		if got < 9999999 && got > 99999999 {
			t.Errorf("PositiveNegative got size error %v", got)
		} else {
			t.Logf("PositiveNegative got %v", got)
		}
	}
}

func BenchmarkPositiveNegative(b *testing.B) {
	testTimes := 10
	for i := 1; i < testTimes; i++ {
		got, err := PositiveNegative(8)
		if err != nil {
			b.Errorf("PositiveNegative() error = %v", err)

		}
		if got < 0 {
			//b.Logf("change PositiveNegative %v", got)
			got = got * -1
		}
		if got < 9999999 && got > 99999999 {
			b.Errorf("PositiveNegative got size error %v", got)
		} else {
			//b.Logf("PositiveNegative got %v", got)
		}
	}
}

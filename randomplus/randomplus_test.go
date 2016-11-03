package randomplus

import "testing"

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
			name:"postive",
			args:args{
				size:7,
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
	test_times := 10000
	for i := 1; i < test_times; i ++ {
		got, err := Positive(8)
		if err != nil {
			t.Errorf("Postive() error = %v", err)

		}
		if got < 99999999 && got > 9999999 {
			t.Logf("Postive got %v", got)
		} else {
			t.Errorf("Postive got number error %v", got)
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
			name:"PositiveNegative",
			args:args{
				size:7,
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
	test_times := 10000
	for i := 1; i < test_times; i ++ {
		got, err := PositiveNegative(8)
		if err != nil {
			t.Errorf("PositiveNegative() error = %v", err)

		}
		if got < 0 {
			t.Logf("change PositiveNegative %v", got)
			got = got * -1
		}
		if got < 99999999 && got > 9999999 {
			t.Logf("PositiveNegative got %v", got)
		} else {
			t.Errorf("PositiveNegative got size error %v", got)
		}
	}
}

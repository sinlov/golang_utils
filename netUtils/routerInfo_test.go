package netUtils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetGateway(t *testing.T) {
	tests := []struct {
		name    string
		want    [4]byte
		wantErr bool
	}{
		{
			name: "base",
			want: [4]byte{0, 0, 0, 0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetGateway()
			if (err != nil) != tt.wantErr {
				t.Errorf("GetGateway() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			t.Logf("got: %v, err %s", got, err)
			assert.NotEqual(t, tt.want, got)
			//if !reflect.DeepEqual(got, tt.want) {
			//	t.Errorf("GetGateway() = %v, want %v", got, tt.want)
			//}
		})
	}
}

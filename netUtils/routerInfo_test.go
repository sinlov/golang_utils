package netUtils

import (
	"reflect"
	"testing"
	"fmt"
)

func TestGetGateway(t *testing.T) {
	tests := []struct {
		name    string
		want    [4]byte
		wantErr bool
	}{
		{
			name: "base",

		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetGateway()
			if (err != nil) != tt.wantErr {
				t.Errorf("GetGateway() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			fmt.Printf("got: %v, err %s", got, err)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetGateway() = %v, want %v", got, tt.want)
			}
		})
	}
}

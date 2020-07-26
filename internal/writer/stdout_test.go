package writer

import (
	"reflect"
	"testing"
)

func TestNewStdOut(t *testing.T) {
	tests := []struct {
		name    string
		want    Writer
		wantErr bool
	}{
		{
			name:    "NewStdOutWriter",
			want:    &StdOut{},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewStdOut()
			if (err != nil) != tt.wantErr {
				t.Errorf("NewStdOut() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewStdOut() got = %v, want %v", got, tt.want)
			}
		})
	}
}

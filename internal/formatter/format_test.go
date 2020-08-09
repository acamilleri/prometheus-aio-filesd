package formatter

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestListFormatterAvailable(t *testing.T) {
	tests := []struct {
		name     string
		register map[string]func() Formatter
		want     []string
	}{
		{
			name:     "ListFormatters",
			register: registered,
			want: []string{
				"json",
				"yaml",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ListFormatterAvailable(); !assert.EqualValues(t, got, tt.want) {
				t.Errorf("ListFormatterAvailable() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNew(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name        string
		args        args
		want        Formatter
		expectedErr error
		wantErr     bool
	}{
		{
			name: "NewJSONFormatter",
			args: args{
				name: "json",
			},
			want:    JSONFormatter{},
			wantErr: false,
		},
		{
			name: "NewYAMLFormatter",
			args: args{
				name: "yaml",
			},
			want:    YAMLFormatter{},
			wantErr: false,
		},
		{
			name: "NewUnknownFormatter",
			args: args{
				name: "unknown",
			},
			want:        nil,
			expectedErr: ErrFormatterNotFound,
			wantErr:     true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := New(tt.args.name)
			if (err != nil) != tt.wantErr {
				t.Errorf("New() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() got = %v, want %v", got, tt.want)
			}

			if tt.wantErr {
				assert.Equal(t, tt.expectedErr, err)
			}
		})
	}
}

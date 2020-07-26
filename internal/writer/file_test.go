package writer

import (
	"errors"
	"os"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

type envVars struct {
	Key   string
	Value string
}

func TestNewFile(t *testing.T) {
	tests := []struct {
		name        string
		vars        []envVars
		want        Writer
		expectedErr error
		wantErr     bool
	}{
		{
			name: "NewFileWriter",
			vars: []envVars{
				{
					Key:   "FILESD_WRITER_FILE_DEST",
					Value: "/tmp/file",
				},
			},
			want: &File{
				Dest: "/tmp/file",
			},
			wantErr: false,
		},
		{
			name:        "NewFileWriterErr",
			vars:        []envVars{},
			want:        nil,
			expectedErr: errors.New("required key FILESD_WRITER_FILE_DEST missing value"),
			wantErr:     true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer func() {
				for _, envVar := range tt.vars {
					_ = os.Unsetenv(envVar.Key)
				}
			}()
			for _, envVar := range tt.vars {
				_ = os.Setenv(envVar.Key, envVar.Value)
			}
			got, err := NewFile()
			if (err != nil) != tt.wantErr {
				t.Errorf("NewFile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewFile() got = %v, want %v", got, tt.want)
			}

			if tt.wantErr {
				assert.Equal(t, tt.expectedErr, err)
			}
		})
	}
}

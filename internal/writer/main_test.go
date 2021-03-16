package writer

import (
	"os"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestListAvailableWriters(t *testing.T) {
	tests := []struct {
		name     string
		register map[string]func() (Writer, error)
		want     []string
	}{
		{
			name:     "ListWriters",
			register: registered,
			want: []string{
				"file",
				"stdout",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ListAvailableWriters(); !assert.ElementsMatch(t, got, tt.want) {
				t.Errorf("ListAvailableWriters() = %v, want %v", got, tt.want)
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
		vars        []envVars
		want        Writer
		expectedErr error
		wantErr     bool
	}{
		{
			name:    "NewStdoutWriter",
			args:    args{name: "stdout"},
			want:    &StdOut{},
			wantErr: false,
		},
		{
			name: "NewFileWriter",
			args: args{name: "file"},
			vars: []envVars{
				{
					Key:   "FILESD_WRITER_FILE_DEST",
					Value: "/tmp/file",
				},
			},
			want:    &File{Dest: "/tmp/file"},
			wantErr: false,
		},
		{
			name:        "NewUnknownWriter",
			args:        args{name: "unknown"},
			want:        nil,
			expectedErr: ErrWriterNotFound,
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

type fooBar struct{}

func (f *fooBar) Write(data []byte) error {
	return nil
}

func TestRegister(t *testing.T) {
	loadFn := func() (Writer, error) {
		return &fooBar{}, nil
	}

	type args struct {
		name   string
		loadFn func() (Writer, error)
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "RegisterFooBarWriter",
			args: args{
				name:   "FooBar",
				loadFn: loadFn,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Register(tt.args.name, tt.args.loadFn)
			assert.Contains(t, registered, "foobar")
		})
	}
}

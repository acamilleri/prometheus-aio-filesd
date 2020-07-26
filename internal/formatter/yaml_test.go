package formatter

import (
	"reflect"
	"testing"

	"github.com/sirupsen/logrus"
)

func TestYAMLFormatter_Format(t *testing.T) {
	type args struct {
		in interface{}
	}
	tests := []struct {
		name    string
		args    args
		want    []byte
		wantErr bool
	}{
		{
			name: "TestFormat",
			args: args{in: struct {
				Foo string `yaml:"foo"`
			}{
				Foo: "bar",
			}},
			want: []byte("Foo: bar\n"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			y := YAMLFormatter{}
			got, err := y.Format(tt.args.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("Format() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				logrus.Info(string(got))
				t.Errorf("Format() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewYAMLFormatter(t *testing.T) {
	tests := []struct {
		name string
		want Formatter
	}{
		{
			name: "NewYAMLFormatter",
			want: YAMLFormatter{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewYAMLFormatter(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewYAMLFormatter() = %v, want %v", got, tt.want)
			}
		})
	}
}

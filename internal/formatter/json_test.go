package formatter

import (
	"reflect"
	"testing"
)

func TestJSONFormatter_Format(t *testing.T) {
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
				Foo string `json:"foo"`
			}{
				Foo: "bar",
			}},
			want: []byte(`{"foo":"bar"}`),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			y := JSONFormatter{}
			got, err := y.Format(tt.args.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("Format() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Format() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewJSONFormatter(t *testing.T) {
	tests := []struct {
		name string
		want Formatter
	}{
		{
			name: "NewJSONFormatter",
			want: JSONFormatter{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewJSONFormatter(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewJSONFormatter() = %v, want %v", got, tt.want)
			}
		})
	}
}

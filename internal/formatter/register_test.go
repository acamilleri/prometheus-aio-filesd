package formatter

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type fooBar struct{}

func (f fooBar) Format(in interface{}) ([]byte, error) {
	return nil, nil
}

func TestRegister(t *testing.T) {
	loadFn := func() Formatter {
		return fooBar{}
	}

	type args struct {
		name   string
		loadFn func() Formatter
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "RegisterFooBarFormatter",
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

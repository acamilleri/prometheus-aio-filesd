package core

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/acamilleri/prometheus-aio-filesd/internal/models"
)

func TestListAvailableProviders(t *testing.T) {
	Register("foobar", func() (Provider, error) {
		return &fooBar{}, nil
	})

	tests := []struct {
		name     string
		register map[string]func() (Provider, error)
		want     []string
	}{
		{
			name:     "ListProviders",
			register: registered,
			want: []string{
				"foobar",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ListAvailableProviders(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ListAvailableProviders() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNew(t *testing.T) {
	Register("foobar", func() (Provider, error) {
		return &fooBar{}, nil
	})

	type args struct {
		name string
	}
	tests := []struct {
		name        string
		args        args
		want        Provider
		expectedErr error
		wantErr     bool
	}{
		{
			name: "NewProvider",
			args: args{name: "foobar"},
			want: &fooBar{},
		},
		{
			name:        "NewProviderErr",
			args:        args{name: "unknown"},
			want:        nil,
			expectedErr: ErrProviderNotFound,
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

func TestRegister(t *testing.T) {
	loadFn := func() (Provider, error) {
		return &fooBar{}, nil
	}

	type args struct {
		name   string
		loadFn func() (Provider, error)
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "RegisterFooBarProvider",
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

type fooBar struct{}

func (f *fooBar) Name() string {
	return "foobar_provider"
}

func (f *fooBar) ListTargets() ([]models.Target, error) {
	return []models.Target{
		{
			Name:        "target1",
			Host:        "127.0.0.1",
			Port:        9090,
			MetricsPath: "/metrics",
			Scheme:      "https",
			Labels: models.Label{
				"foo":      "bar",
				"provider": "fake_provider",
				"source":   "prometheus-aio-filesd",
			},
		},
	}, nil
}

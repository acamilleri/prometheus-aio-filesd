package adapter

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/acamilleri/prometheus-aio-filesd/internal/formatter"
	"github.com/acamilleri/prometheus-aio-filesd/internal/models"
	"github.com/acamilleri/prometheus-aio-filesd/internal/provider/core"
	"github.com/acamilleri/prometheus-aio-filesd/internal/writer"
)

func TestNew(t *testing.T) {
	type args struct {
		provider  core.Provider
		writer    writer.Writer
		formatter formatter.Formatter
	}
	tests := []struct {
		name        string
		args        args
		want        Adapter
		expectedErr error
		wantErr     bool
	}{
		{
			name: "NewAdapterProviderErr",
			args: args{
				formatter: &formatter.JSONFormatter{},
				writer:    &writer.StdOut{},
				provider:  nil,
			},
			want:        nil,
			expectedErr: ErrAdapterProviderEmpty,
			wantErr:     true,
		},
		{
			name: "NewAdapterFormatterErr",
			args: args{
				provider:  &fakeProvider{},
				writer:    &writer.StdOut{},
				formatter: nil,
			},
			want:        nil,
			expectedErr: ErrAdapterFormatterEmpty,
			wantErr:     true,
		},
		{
			name: "NewAdapterWriterErr",
			args: args{
				provider:  &fakeProvider{},
				formatter: &formatter.JSONFormatter{},
				writer:    nil,
			},
			want:        nil,
			expectedErr: ErrAdapterWriterEmpty,
			wantErr:     true,
		},
		{
			name: "NewAdapter",
			args: args{
				provider:  &fakeProvider{},
				formatter: &formatter.JSONFormatter{},
				writer:    &writer.StdOut{},
			},
			want: &adapter{
				provider:  &fakeProvider{},
				writer:    &writer.StdOut{},
				formatter: &formatter.JSONFormatter{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := New(tt.args.provider, tt.args.writer, tt.args.formatter)
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

type fakeProvider struct{}

func (f fakeProvider) Name() string {
	return "fake_provider"
}

func (f fakeProvider) ListTargets() ([]models.Target, error) {
	return nil, nil
}

func Test_adapter_generateTargetGroups(t *testing.T) {
	type fields struct {
		provider  core.Provider
		writer    writer.Writer
		formatter formatter.Formatter
	}
	type args struct {
		targets []models.Target
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []models.Group
		wantErr bool
	}{
		{
			name: "TestGenerateTargetGroups",
			fields: fields{
				provider:  &fakeProvider{},
				writer:    &writer.StdOut{},
				formatter: &formatter.JSONFormatter{},
			},
			args: args{
				targets: []models.Target{
					{
						Name:        "target1",
						Host:        "127.0.0.1",
						Port:        9090,
						MetricsPath: "/metrics",
						Scheme:      "https",
						Labels: models.Label{
							"foo": "bar",
						},
					},
				},
			},
			want: []models.Group{
				{
					Targets: []models.Target{
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
					},
					Labels: models.Label{
						"foo":      "bar",
						"provider": "fake_provider",
						"source":   "prometheus-aio-filesd",
					},
					Source: "fake_provider",
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &adapter{
				provider:  tt.fields.provider,
				writer:    tt.fields.writer,
				formatter: tt.fields.formatter,
			}
			got, err := a.generateTargetGroups(tt.args.targets)
			if (err != nil) != tt.wantErr {
				t.Errorf("generateTargetGroups() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("generateTargetGroups() got = %v, want %v", got, tt.want)
			}
		})
	}
}

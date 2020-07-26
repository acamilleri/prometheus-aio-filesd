package adapter

import (
	promModel "github.com/prometheus/common/model"

	"github.com/acamilleri/prometheus-aio-filesd/internal/formatter"
	"github.com/acamilleri/prometheus-aio-filesd/internal/models"
	provider "github.com/acamilleri/prometheus-aio-filesd/internal/provider/core"
	"github.com/acamilleri/prometheus-aio-filesd/internal/writer"
)

// Adapter ...
type Adapter interface {
	Run() error
}

type adapter struct {
	provider  provider.Provider
	writer    writer.Writer
	formatter formatter.Formatter
}

// New Initialize the Adapter struct
func New(provider provider.Provider, writer writer.Writer, formatter formatter.Formatter) (Adapter, error) {
	if provider == nil {
		return nil, ErrAdapterProviderEmpty
	}

	if writer == nil {
		return nil, ErrAdapterWriterEmpty
	}

	if formatter == nil {
		return nil, ErrAdapterFormatterEmpty
	}

	return &adapter{
		provider:  provider,
		writer:    writer,
		formatter: formatter,
	}, nil
}

// Run Adapter fetch targets from Provider(s) and
// export a Prometheus FileSD TargetGroups
// https://prometheus.io/docs/prometheus/latest/configuration/configuration/#file_sd_config
func (a *adapter) Run() error {
	targets, err := a.provider.ListTargets()
	if err != nil {
		return err
	}

	tgs, err := a.generateTargetGroups(targets)
	if err != nil {
		return err
	}

	return a.write(tgs)
}

func (a *adapter) generateTargetGroups(targets []models.Target) ([]models.Group, error) {
	var tgs []models.Group

	// require project labels
	sourceLabel := "source"
	providerLabel := "provider"

	for _, target := range targets {
		var group models.Group
		group.Source = a.provider.Name()
		group.Targets = append(group.Targets, target)

		labels := target.Labels
		if labels == nil {
			labels = make(models.Label)
		}
		labels[promModel.LabelName(sourceLabel)] = promModel.LabelValue("prometheus-aio-filesd")
		labels[promModel.LabelName(providerLabel)] = promModel.LabelValue(a.provider.Name())
		group.Labels = labels

		tgs = append(tgs, group)
	}
	return tgs, nil
}

func (a *adapter) write(tgs []models.Group) error {
	data, err := a.formatter.Format(tgs)
	if err != nil {
		return err
	}

	return a.writer.Write(data)
}

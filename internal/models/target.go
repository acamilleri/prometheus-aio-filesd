package models

import (
	"encoding/json"
	"fmt"

	"github.com/ghodss/yaml"
	"github.com/prometheus/common/model"
)

var (
	// DefaultScrapeLabel label use to define is the target
	// must be scrape by Prometheus
	DefaultScrapeLabel = "prometheus.io/scrape"
	// DefaultScrapePortLabel label use to define the port
	// of the metrics handler of the target
	DefaultScrapePortLabel = "prometheus.io/port"
	// DefaultScrapeHostLabel label use to define the host
	// of the target
	DefaultScrapeHostLabel = "prometheus.io/host"
	// DefaultScrapePathLabel label use to define the metrics
	// path of the target
	DefaultScrapePathLabel = "prometheus.io/path"
	// DefaultScrapeSchemeLabel label use to define the scheme
	// of the target
	DefaultScrapeSchemeLabel = "prometheus.io/scheme"
)

// Label ...
type Label model.LabelSet

// Target Object
type Target struct {
	Name        string
	Host        string
	Port        int
	MetricsPath string
	Scheme      string
	Labels      Label
}

// MarshalJSON ...
func (t Target) MarshalJSON() ([]byte, error) {
	addr := t.Host

	if t.Host == "" {
		return nil, fmt.Errorf("no valid host for target %s", t.Name)
	}

	if t.Port != 0 {
		addr = fmt.Sprintf("%s:%d", addr, t.Port)
	}

	return json.Marshal(addr)
}

// MarshalYAML ...
func (t Target) MarshalYAML() (interface{}, error) {
	addr := t.Host

	if t.Host == "" {
		return nil, fmt.Errorf("no valid host for target %s", t.Name)
	}

	if t.Port != 0 {
		addr = fmt.Sprintf("%s:%d", addr, t.Port)
	}

	return yaml.Marshal(addr)
}

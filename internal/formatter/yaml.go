package formatter

import "github.com/ghodss/yaml"

func init() {
	Register("yaml", NewYAMLFormatter)
}

// YAMLFormatter ...
type YAMLFormatter struct{}

// NewYAMLFormatter ...
func NewYAMLFormatter() Formatter {
	return YAMLFormatter{}
}

// Format marshalling input data as YAML
func (y YAMLFormatter) Format(in interface{}) ([]byte, error) {
	return yaml.Marshal(in)
}

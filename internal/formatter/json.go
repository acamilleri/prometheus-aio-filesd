package formatter

import "encoding/json"

func init() {
	Register("json", NewJSONFormatter)
}

// JSONFormatter ...
type JSONFormatter struct{}

// NewJSONFormatter ...
func NewJSONFormatter() Formatter {
	return JSONFormatter{}
}

// Format marshalling input data as JSON
func (y JSONFormatter) Format(in interface{}) ([]byte, error) {
	return json.Marshal(in)
}

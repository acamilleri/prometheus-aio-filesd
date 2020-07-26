package writer

import (
	"fmt"
	"os"

	"github.com/kelseyhightower/envconfig"
)

func init() {
	Register("stdout", NewStdOut)
}

// StdOut ...
type StdOut struct{}

// NewStdOut StdOut writer
func NewStdOut() (Writer, error) {
	var std StdOut
	err := envconfig.Process(fmt.Sprintf("%s_%s", DefaultEnvVarsPrefix, "stdout"), &std)
	if err != nil {
		return nil, err
	}

	return &std, nil
}

// Write ...
func (std *StdOut) Write(data []byte) error {
	_, err := fmt.Fprintf(os.Stdout, "%s\n", string(data))
	return err
}

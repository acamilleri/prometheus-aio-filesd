package writer

import "strings"

var registered = make(map[string]func() (Writer, error))

// Writer ...
type Writer interface {
	Write(data []byte) error
}

// New ...
func New(name string) (Writer, error) {
	if writer, ok := registered[name]; ok {
		return writer()
	}
	return nil, ErrWriterNotFound
}

// Register Register a writer
func Register(name string, loadFn func() (Writer, error)) {
	name = strings.ToLower(name)
	registered[name] = loadFn
}

// ListAvailableWriters List writer registered
func ListAvailableWriters() []string {
	var writers []string

	for writerName := range registered {
		writers = append(writers, writerName)
	}
	return writers
}

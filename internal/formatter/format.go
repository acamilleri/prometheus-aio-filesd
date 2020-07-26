package formatter

// Formatter ...
type Formatter interface {
	Format(in interface{}) ([]byte, error)
}

// ListFormatterAvailable formatter available
func ListFormatterAvailable() []string {
	var fs []string

	for formatterName := range registered {
		fs = append(fs, formatterName)
	}

	return fs
}

// New ...
func New(name string) (Formatter, error) {
	if formatter, ok := registered[name]; ok {
		return formatter(), nil
	}
	return nil, ErrFormatterNotFound
}

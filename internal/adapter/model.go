package adapter

import "errors"

var (
	// ErrAdapterProviderEmpty Error occurred when the provider is nil
	ErrAdapterProviderEmpty = errors.New("provider is empty")
	// ErrAdapterWriterEmpty Error occurred when the writer is nil
	ErrAdapterWriterEmpty = errors.New("writer is empty")
	// ErrAdapterFormatterEmpty Error occurred when the formatter is nil
	ErrAdapterFormatterEmpty = errors.New("formatter is empty")
)

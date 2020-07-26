package adapter

import "errors"

var (
	// ErrAdapterProviderEmpty Error occured when the provider is nil
	ErrAdapterProviderEmpty = errors.New("provider is empty")
	// ErrAdapterWriterEmpty Error occured when the writer is nil
	ErrAdapterWriterEmpty = errors.New("writer is empty")
	// ErrAdapterFormatterEmpty Error occured when the formatter is nil
	ErrAdapterFormatterEmpty = errors.New("formatter is empty")
)

package writer

import "errors"

var (
	// DefaultEnvVarsPrefix Env config var prefix
	DefaultEnvVarsPrefix = "filesd_writer"

	// ErrWriterNotFound Error occured when the writer is not found
	ErrWriterNotFound = errors.New("writer not found")
)



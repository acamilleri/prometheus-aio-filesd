package core

import "errors"

var (
	// DefaultEnvVarsPrefix Env config var prefix
	DefaultEnvVarsPrefix = "filesd_provider"

	// ErrProviderNotFound Error occurred when the provider is not found
	ErrProviderNotFound = errors.New("provider not found")
)



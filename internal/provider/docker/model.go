package docker

import "errors"

var (
	// ErrDockerListFailed Error occured when docker failed to list containers
	ErrDockerListFailed = errors.New("failed to list containers")
)

package core

import (
	"strings"

	"github.com/acamilleri/prometheus-aio-filesd/internal/models"
)

var registered = make(map[string]func() (Provider, error))

// Provider Interface to fetch Target(s)
type Provider interface {
	Name() string
	ListTargets() ([]models.Target, error)
}

// Register Register a provider
func Register(name string, loadFn func() (Provider, error)) {
	name = strings.ToLower(name)
	registered[name] = loadFn
}

// ListAvailableProviders List provider registered
func ListAvailableProviders() []string {
	var providers []string

	for providerName := range registered {
		providers = append(providers, providerName)
	}
	return providers
}

// New Initialize a registered provider
func New(name string) (Provider, error) {
	if provider, ok := registered[name]; ok {
		return provider()
	}
	return nil, ErrProviderNotFound
}

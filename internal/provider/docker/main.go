package docker

import (
	"net/http"
	"path/filepath"

	docker "github.com/docker/docker/client"
	"github.com/docker/go-connections/tlsconfig"
	"github.com/kelseyhightower/envconfig"

	provider "github.com/acamilleri/prometheus-aio-filesd/internal/provider/core"
)

func init() {
	provider.Register("docker", newDockerClient)
}

type clientOptions struct {
	Host       string `envconfig:"DOCKER_HOST" default:"unix:///var/run/docker.sock"`
	APIVersion string `envconfig:"DOCKER_API_VERSION" default:"1.25" split_words:"true"`
	TLS        clientTLSOptions
}

type clientTLSOptions struct {
	CertificatePath    string `envconfig:"DOCKER_CERT_PATH"`
	InsecureSkipVerify bool   `envconfig:"DOCKER_TLS_VERIFY" split_words:"true"`
}

func loadOptionFromEnv() (clientOptions, error) {
	var options clientOptions
	err := envconfig.Process(provider.DefaultEnvVarsPrefix, &options)
	if err != nil {
		return clientOptions{}, err
	}
	return options, nil
}

func newDockerClient() (provider.Provider, error) {
	options, err := loadOptionFromEnv()
	if err != nil {
		return nil, err
	}

	var httpClient *http.Client
	if options.TLS != (clientTLSOptions{}) {
		options := tlsconfig.Options{
			CAFile:             filepath.Join(options.TLS.CertificatePath, "ca.pem"),
			CertFile:           filepath.Join(options.TLS.CertificatePath, "cert.pem"),
			KeyFile:            filepath.Join(options.TLS.CertificatePath, "key.pem"),
			InsecureSkipVerify: options.TLS.InsecureSkipVerify,
		}

		tlsc, err := tlsconfig.Client(options)
		if err != nil {
			return nil, err
		}

		httpClient = &http.Client{
			Transport: &http.Transport{
				TLSClientConfig: tlsc,
			},
		}
	}

	client, err := docker.NewClient(options.Host, options.APIVersion, httpClient, nil)
	if err != nil {
		return nil, err
	}
	return &manager{client: client}, nil
}

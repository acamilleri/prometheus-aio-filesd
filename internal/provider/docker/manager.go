package docker

import (
	"context"
	"fmt"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/filters"
	"github.com/docker/docker/client"

	"github.com/acamilleri/prometheus-aio-filesd/internal/models"
)

// Manager Implement the Provider interface
type manager struct {
	// Docker client
	client *client.Client
}

// Name return the Provider name
func (m *manager) Name() string {
	return "docker"
}

// ListTargets Manager fetch Docker container(s)
// containing prometheus.io labels like Kubernetes
// https://www.weave.works/docs/cloud/latest/tasks/monitor/configuration-k8s/#per-pod-prometheus-annotations
func (m *manager) ListTargets() ([]models.Target, error) {
	ctx := context.Background()

	label := fmt.Sprintf("%s=%t", models.DefaultScrapeLabel, true)
	f := filters.NewArgs()
	f.Add("label", label)

	containers, err := m.client.ContainerList(ctx, types.ContainerListOptions{
		Filters: f,
	})
	if err != nil {
		return nil, ErrDockerListFailed
	}

	var targets []models.Target
	for _, container := range containers {
		container, err := m.client.ContainerInspect(ctx, container.ID)
		if err != nil {
			continue
		}
		targets = append(targets, ContainerJSONObjectToTargetObject(Container(container)))
	}
	return targets, nil
}

package docker

import (
	"fmt"
	"net"
	"strconv"
	"strings"

	"github.com/docker/docker/api/types"
	"github.com/prometheus/common/model"

	"github.com/acamilleri/prometheus-aio-filesd/internal/models"
)

const labelPrefix = "container"

// Container Object to wrap the Docker
// ContainerJSON Object
type Container types.ContainerJSON

// ContainerJSONObjectToTargetObject Convert Container Object
// to Target Object
// Add labels to target
func ContainerJSONObjectToTargetObject(container Container) models.Target {
	idLabel := fmt.Sprintf("%s_id", labelPrefix)
	imageNameLabel := fmt.Sprintf("%s_image_name", labelPrefix)
	imageVersionLabel := fmt.Sprintf("%s_image_version", labelPrefix)

	return models.Target{
		Name:        container.GetName(),
		Host:        container.GetHost(),
		Port:        container.GetPort(),
		MetricsPath: container.GetMetricsPath(),
		Labels: models.Label{
			model.LabelName(idLabel):                model.LabelValue(container.GetID()),
			model.LabelName(imageNameLabel):         model.LabelValue(container.GetImage().GetName()),
			model.LabelName(imageVersionLabel):      model.LabelValue(container.GetImage().GetVersion()),
			model.LabelName(model.MetricsPathLabel): model.LabelValue(container.GetMetricsPath()),
			model.LabelName(model.SchemeLabel):      model.LabelValue(container.GetScheme()),
		},
	}
}

// GetID Return a short id
func (c Container) GetID() string {
	if len(c.ID) >= 12 {
		// return only first twelve characters
		return c.ID[0:12]
	}

	return c.ID
}

// GetName Return the name of container
func (c Container) GetName() string {
	if strings.HasPrefix(c.Name, "/") {
		return strings.Trim(c.Name, "/")
	}
	return c.Name
}

// ContainerImage Object
type ContainerImage struct {
	Name    string
	Version string
}

// GetImage compute and return an ContainerImage Object
// from Container config
func (c Container) GetImage() ContainerImage {
	var name, version string
	if strings.Contains(c.Config.Image, ":") {
		i := strings.Split(c.Config.Image, ":")
		name = i[0]
		version = i[1]
	} else {
		name = c.Config.Image
		version = "latest"
	}

	return ContainerImage{
		Name:    name,
		Version: version,
	}
}

// GetName return the image name of Container
func (i ContainerImage) GetName() string {
	return i.Name
}

// GetVersion return the image version of Container
func (i ContainerImage) GetVersion() string {
	return i.Version
}

// GetHost Return the IP Addr or Hostname of container
// Precedence value:
//	- From the 'prometheus.io/host' label.
//  - From the IP Address field of NetworkSettings struct.
//  - From the first IP Address available when loop on Networks map
//    of NetworkSettings struct.
func (c Container) GetHost() string {
	var host string

	if hostValue, ok := c.Config.Labels[models.DefaultScrapeHostLabel]; ok {
		host = hostValue
	} else if c.NetworkSettings.IPAddress != "" {
		host = net.ParseIP(c.NetworkSettings.IPAddress).String()
	} else if c.NetworkSettings.Networks != nil {
		for _, network := range c.NetworkSettings.Networks {
			if network.IPAddress != "" {
				host = net.ParseIP(network.IPAddress).String()
				break
			}
		}
	}
	return host
}

// GetPort Return the port of container
// Precedence value:
//	- From the 'prometheus.io/port' label.
func (c Container) GetPort() int {
	var port int

	if portValue, ok := c.Config.Labels[models.DefaultScrapePortLabel]; ok {
		port, _ = strconv.Atoi(portValue)
	}
	return port
}

// GetMetricsPath Return the port of container
// Precedence value:
//	- From the 'prometheus.io/path' label.
//  - Default value: /metrics
func (c Container) GetMetricsPath() string {
	var metricsPath = "/metrics"

	if pathValue, ok := c.Config.Labels[models.DefaultScrapePathLabel]; ok {
		metricsPath = pathValue
	}
	return metricsPath
}

// GetScheme return the scheme of the metrics path url of Container
func (c Container) GetScheme() string {
	var scheme = "http"

	if pathValue, ok := c.Config.Labels[models.DefaultScrapeSchemeLabel]; ok {
		scheme = pathValue
	}

	return scheme
}

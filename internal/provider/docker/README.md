# Docker Provider

This provider fetches all containers with the label `prometheus.io/scrape=true`.

## Configuration

All options can be configured from environment variables:

| Name                               | Alternative Name   | Required           | Default                     |
|------------------------------------|--------------------|--------------------|-----------------------------|
| FILESD_PROVIDER_DOCKER_HOST        | DOCKER_HOST        | Yes                | unix:///var/run/docker.sock |
| FILESD_PROVIDER_DOCKER_API_VERSION | DOCKER_API_VERSION | Yes                | 1.25                        |
| FILESD_PROVIDER_DOCKER_CERT_PATH   | DOCKER_CERT_PATH   | No                 | ""                          |
| FILESD_PROVIDER_DOCKER_TLS_VERIFY  | DOCKER_TLS_VERIFY  | No                 | False                       |

## Containers Labels

| Name                 | Default      | Description                              |
|----------------------|--------------|------------------------------------------|
| prometheus.io/scrape | False        | Define if the container must be scrape   |
| prometheus.io/host   | [Container IP](https://github.com/acamilleri/prometheus-aio-filesd/blob/master/internal/provider/docker/converts.go#L96) | Define the host of the metrics handler   |
| prometheus.io/port   | 80           | Define the port of the metrics handler   |
| prometheus.io/path   | /metrics     | Define the path of the metrics handler   |
| prometheus.io/scheme | http         | Define the scheme of the metrics handler |

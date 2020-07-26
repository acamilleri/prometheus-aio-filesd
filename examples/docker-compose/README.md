# Docker Compose example

This example is based on the Docker provider.

We will run some containers and automatically adding as targets in Prometheus.

# Usage

Run
```
$ docker-compose up -d
```

This will create three Docker containers:
- Prometheus
- Node exporter
- Filesd

## Prometheus

Visit your [Prometheus](http://localhost:9090/targets) previously started with the docker-compose file.

You must have two targets:
 - The Node exporter container with the port 9100.
 - The Prometheus container with the port 9090.
 
# How it works ?

The Filesd container is configure to query the [Docker API](https://docs.docker.com/engine/api/) by the [Docker socket](https://docs.docker.com/engine/reference/commandline/dockerd/#daemon-socket-option) 
linked into the container.
The Filesd container will fetch all containers with the label `prometheus.io/scrape=true` and write an 
`/file_sd/docker.json` file every 5 min, the duration value can be increase or decrease in the configuration.

The `/file_sd` directory it's a Docker volume mounted in the Filesd container and the Prometheus container, but for the 
Prometheus container, the directory was mount into the `/etc/prometheus/file_sd` directory.

Prometheus will simply look this repository `/etc/prometheus/file_sd/*.json` and update targets when changes occurred.

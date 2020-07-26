# Prometheus all in one file_sd integration

[![Go Report Card](https://goreportcard.com/badge/github.com/acamilleri/prometheus-aio-filesd)](https://goreportcard.com/report/github.com/acamilleri/prometheus-aio-filesd)

This project provides a Provider interface to easily integrate any provider
to fetch some targets to export a compatible Prometheus file_sd format.

**This project works, but is still ongoing**

# Usage

## Docker:

You can find an example [here](./examples/docker-compose).

# Providers

List of available providers:
- Docker

Feel free to contribute to adding new providers or [create a provider integration request with an issue](https://github.com/acamilleri/prometheus-aio-filesd/issues/new?assignees=&labels=provider%2Frequest&template=provider-request.md&title=%5BProvider%5D+Add+...) :)

# Build from source

clone the project
```
$ git clone git@github.com:acamilleri/prometheus-aio-filesd.git
```

build
```
$ make build
```

# TODO:
- Improve code (tests!, logger, comments, ..)
- Adding metrics
- Add CI
- Add more providers

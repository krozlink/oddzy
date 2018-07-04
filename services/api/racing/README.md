# Racing Service

This is the Racing service

Generated with

```
micro new github.com/krozlink/oddzy/services/api/racing --namespace=go.micro --type=api
```

## Getting Started

- [Configuration](#configuration)
- [Dependencies](#dependencies)
- [Usage](#usage)

## Configuration

- FQDN: go.micro.api.racing
- Type: api
- Alias: racing

## Dependencies

Micro services depend on service discovery. The default is consul.

```
# install consul
brew install consul

# run consul
consul agent -dev
```

## Usage

A Makefile is included for convenience

Build the binary

```
make build
```

Run the service
```
./racing-api
```

Build a docker image
```
make docker
```
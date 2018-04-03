# go-primogen

[![License](https://img.shields.io/badge/license-Apache%20License%202.0-blue.svg?style=flat)][license]
[![Build Status](https://travis-ci.org/steenzout/go-primogen.svg?branch=master)](https://travis-ci.org/steenzout/go-primogen/)
[![Coverage Status](https://coveralls.io/repos/steenzout/go-primogen/badge.svg?branch=master&service=github)](https://coveralls.io/github/steenzout/go-primogen?branch=master)
[![Go Report Card](https://goreportcard.com/badge/github.com/steenzout/go-primogen)](https://goreportcard.com/report/github.com/steenzout/go-primogen)

[![Project](https://www.openhub.net/p/go-steenzout-primogen/widgets/project_thin_badge.gif)][project]

## Prerequisites

Be sure you have [protoc](https://github.com/google/protobuf/releases) installed.

Also:

```bash
$ go get -u google.golang.org/grpc

$ go get -u github.com/golang/protobuf/protoc-gen-go
```

## Structure

- PokéAPI gRPC service: simple RPC
  - `pokeapi/service.proto`: contains the service definition and the message request and reply.
  - `pokeapi/service/service.pb.go`: the generated Go code for the service.
  - `pokeapi/server/`:.
  - `pokeapi/client/`:.
- : server-side streaming RPC
- : client-side streaming RPC
- : bidirectional streaming RPC

## Links

- [google/protobuf releases](https://github.com/google/protobuf/releases)
- [Protocol Buffes > Language Guide (proto3)](https://developers.google.com/protocol-buffers/docs/proto3)
- [gRPC Quick Start](https://grpc.io/docs/quickstart/go.html)
- [PokéAPI](https://pokeapi.co/)

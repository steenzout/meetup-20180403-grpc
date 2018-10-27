# 2018.04.03 Utah Golang Meetup gRPC

[![License](https://img.shields.io/badge/license-Apache%20License%202.0-blue.svg?style=flat)][license]
[![Build Status](https://travis-ci.org/steenzout/meetup-20180403-grpc.svg?branch=master)](https://travis-ci.org/steenzout/meetup-20180403-grpc/)
[![Coverage Status](https://coveralls.io/repos/steenzout/meetup-20180403-grpc/badge.svg?branch=master&service=github)](https://coveralls.io/github/steenzout/meetup-20180403-grpc?branch=master)
[![Go Report Card](https://goreportcard.com/badge/github.com/steenzout/meetup-20180403-grpc)](https://goreportcard.com/report/github.com/steenzout/meetup-20180403-grpc)

## Prerequisites

Be sure you have [protoc](https://github.com/google/protobuf/releases) installed.

Also:

```bash
$ go get -u google.golang.org/grpc

$ go get -u github.com/golang/protobuf/protoc-gen-go

$ dep ensure
```

## Structure

- PokéAPI gRPC service: simple RPC
  - `pokeapi/getrequest.go`: helpers for the `pokeapi.GetRequest` struct.
  - `pokeapi/service.proto`: contains the service definition and the message request and reply.
  - `pokeapi/service.pb.go`: (generated) Go code for the service.
  - `pokeapi/cmd/`: PokéAPI command-line interface package.
  - `pokeapi/pokeapi-cli`: binary to run the PokéAPI command-line interface (OSX).
  - `pokeapi/pokeapi-cli.linux`: binary to run the PokéAPI command-line interface (Linux).
- : server-side streaming RPC
- : client-side streaming RPC
- : bidirectional streaming RPC

## Build

```bash
$ make
```

## Usage

```bash
$ pokeapi-cli
NAME:
   pokeapi-cli - pokeapi-cli is a command-line interface to the PokéAPI.

USAGE:
   pokeapi-cli [global options] command [command options] [arguments...]

VERSION:
   0.0.0

COMMANDS:
     get      get pokemon information
     start    start pokeapi gRPC server
     help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h     show help
   --version, -v  print the version
```

## Links

- [google/protobuf releases](https://github.com/google/protobuf/releases)
- [Protocol Buffers > Language Guide (proto3)](https://developers.google.com/protocol-buffers/docs/proto3)
- [gRPC Quick Start](https://grpc.io/docs/quickstart/go.html)
- [PokéAPI](https://pokeapi.co/)

[license]:  https://raw.githubusercontent.com/steenzout/meetup-20180403-grpc/master/LICENSE   "Apache License 2.0"

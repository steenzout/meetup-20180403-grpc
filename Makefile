default: pokeapi

clean:
	rm pokeapi/service.pb.go
	rm pokeapi/pokeapi-cli
	rm pokeapi/pokeapi-cli.linux

pokeapi: pokeapi_pb pokeapi_build

pokeapi_build: pokeapi/pokeapi-cli # pokeapi/pokeapi-cli.linux
	chmod 755 pokeapi/pokeapi-cli*

pokeapi_pb: pokeapi/service.pb.go

pokeapi/pokeapi-cli:
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -ldflags "-s" -a -installsuffix cgo -o pokeapi/pokeapi-cli ./pokeapi/cmd

pokeapi/pokeapi-cli.linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "-s" -a -installsuffix cgo -o pokeapi/pokeapi-cli.linux ./pokeapi/cmd

pokeapi/service.pb.go:
	protoc \
		-I . \
		--go_out="plugins=grpc:${GOPATH}/src" \
		pokeapi/*.proto

vendor:
	dep ensure -v

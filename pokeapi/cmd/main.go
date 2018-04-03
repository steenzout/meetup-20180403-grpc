package main

import (
	"log"
	"os"

	cli "gopkg.in/urfave/cli.v1"

	"github.com/steenzout/meetup-20180403-grpc/pokeapi/cmd/internal/get"
	"github.com/steenzout/meetup-20180403-grpc/pokeapi/cmd/internal/start"
)

func main() {
	app := cli.NewApp()

	app.Usage = "pokeapi-cli is a command-line interface to the PokéAPI."
	app.Commands = []cli.Command{
		{
			Name:   "get",
			Usage:  "get pokemon information",
			Before: get.ValidateArg,
			Action: get.Run,
		},
		{
			Name:   "start",
			Usage:  "start pokeapi gRPC server",
			Action: start.Run,
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

package get

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/steenzout/meetup-20180403-grpc/pokeapi"

	"google.golang.org/grpc"
	cli "gopkg.in/urfave/cli.v1"

	pb "github.com/steenzout/meetup-20180403-grpc/pokeapi"
)

const (
	address = "localhost:50051"
)

var pokeid uint64

// ValidateArg verifies that the pokemon identifier is not empty.
func ValidateArg(c *cli.Context) error {
	var err error

	arg := c.Args().Get(0)
	if arg == "" {
		return cli.NewExitError("missing pokemon identifier", 86)
	}

	argint, err := strconv.Atoi(arg)
	if err != nil {
		return cli.NewExitError(err.Error(), 86)
	}

	if argint < 0 {
		return cli.NewExitError(fmt.Sprintf("invalid pokemon identifier: %d", argint), 86)
	}

	pokeid = uint64(argint)

	return nil
}

// Run get pokemon information.
func Run(cctx *cli.Context) error {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		return cli.NewExitError(fmt.Errorf("did not connect: %v", err), 86)

	}
	defer conn.Close()

	client := pb.NewPokeapiClient(conn)
	req, err := pokeapi.NewGetRequest(pokeid)
	if err != nil {
		return cli.NewExitError(err.Error(), 86)
	}

	// contact the server and print out the pokemon information
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	reply, err := client.GetPokemon(ctx, req)
	if err != nil {
		return cli.NewExitError(fmt.Errorf("could not GetPokemon: %v", err), 86)
	}
	if reply == nil {
		return cli.NewExitError(fmt.Errorf("could not find pokemon: %d", pokeid), 86)
	}

	fmt.Printf("%s\t%s\t%s\n", reply.Id, reply.Name, reply.Weight)

	return nil
}

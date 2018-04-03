package start

import (
	"context"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	cli "gopkg.in/urfave/cli.v1"

	pb "github.com/steenzout/meetup-20180403-grpc/pokeapi"
)

const (
	port = ":50051"
)

// server pokeapi.PokeapiServer implementation.
type server struct{}

// GetPokemon get information about a pokemon.
func (s *server) GetPokemon(ctx context.Context, in *pb.GetRequest) (*pb.GetReply, error) {
	return &pb.GetReply{
		Name: "test",
	}, nil
}

// Run start gRPC server.
func Run(cctx *cli.Context) error {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterPokeapiServer(s, &server{})
	// Register reflection service on gRPC server.
	reflection.Register(s)

	err = s.Serve(lis)
	if err != nil {
		return cli.NewExitError(fmt.Errorf("failed to serve: %v", err), 86)
	}

	return nil
}

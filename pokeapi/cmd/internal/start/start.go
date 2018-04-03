package start

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net"
	"net/http"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	cli "gopkg.in/urfave/cli.v1"

	pb "github.com/steenzout/meetup-20180403-grpc/pokeapi"
)

const (
	port = ":50051"
)

var (
	errInternal = fmt.Errorf("internal error")
)

// server pokeapi.PokeapiServer implementation.
type server struct{}

// GetPokemon get information about a pokemon.
func (s *server) GetPokemon(ctx context.Context, in *pb.GetRequest) (*pb.GetReply, error) {
	pokemon, err := getPokemon(in.Id)
	if err != nil {
		log.Printf("[ERROR] GetPokemon() %s", err.Error())
		return nil, errInternal
	}
	log.Printf("[INFO] found pokemon %d %v", in.Id, pokemon)

	return &pb.GetReply{
		Id:     int64(pokemon.ID),
		Name:   pokemon.Name,
		Weight: pokemon.Weight,
	}, nil
}

func getPokemon(id uint64) (*Pokemon, error) {
	url := fmt.Sprintf("http://pokeapi.co/api/v2/pokemon/%d/", id)

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New(resp.Status)
	}

	pokemon := new(Pokemon)
	err = json.NewDecoder(resp.Body).Decode(pokemon)
	if err != nil {
		return nil, err
	}

	return pokemon, nil
}

// Run start gRPC server.
func Run(cctx *cli.Context) error {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		return cli.NewExitError(fmt.Errorf("failed to listen: %v", err), 86)
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

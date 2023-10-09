package modules

import (
	"context"
	"log"

	superherov1 "buf.build/gen/go/teddy-lourson/superhero/protocolbuffers/go/superhero/v1"
	"connectrpc.com/connect"
)

// SuperheroServer represents the gRPC server that provides Superhero services.
type SuperheroServer struct {
}

// NewSuperheroServer creates a new SuperheroServer.
func NewSuperheroServer() *SuperheroServer {
	return &SuperheroServer{}
}

// GetSuperheroes is a unary RPC that returns a list of all supeheroes.
func (s *SuperheroServer) GetSuperheroes(
	ctx context.Context,
	req *connect.Request[superherov1.GetSuperheroesRequest],
) (*connect.Response[superherov1.GetSuperheroesResponse], error) {
	log.Printf("Received a GetSuperheroes RPC request")
	res := []*superherov1.Superhero{
		{
			Id:        "34cb7019-f957-43e3-a343-f7688bd83a1e",
			Name:      "Iron Man",
			FirstName: "Tony",
			LastName:  "Stark",
		},
		{
			Id:        "8d44168a-fcdc-4b5b-aa27-346990b32dee",
			Name:      "Spiderman",
			FirstName: "Peter",
			LastName:  "Parker",
		},
		{
			Id:        "dfccd601-410e-4c66-8ef0-ba05097bc208",
			Name:      "Captain America",
			FirstName: "Steve",
			LastName:  "Rogers",
		},
		{
			Id:        "6ab9a81e-2e0d-4f7a-8588-159d06387cd2",
			Name:      "Black Widow",
			FirstName: "Natasha",
			LastName:  "Romanoff",
		},
	}
	return connect.NewResponse(&superherov1.GetSuperheroesResponse{
		Superheroes: res,
	}), nil
}

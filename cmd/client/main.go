package main

import (
	"context"
	"log"
	"net/http"

	"buf.build/gen/go/teddy-lourson/superhero/connectrpc/go/superhero/v1/superherov1connect"
	superherov1 "buf.build/gen/go/teddy-lourson/superhero/protocolbuffers/go/superhero/v1"
	"connectrpc.com/connect"
	"google.golang.org/protobuf/encoding/protojson"
)

func main() {
	client := superherov1connect.NewSuperheroServiceClient(http.DefaultClient, "http://localhost:3000")
	res, err := client.GetSuperheroes(context.Background(), connect.NewRequest(&superherov1.GetSuperheroesRequest{}))
	if err != nil {
		log.Println(err)
		return
	}
	jsm := &protojson.MarshalOptions{
		Indent: "\t",
		// EmitUnpopulated: true,
	}
	toPrint, err := jsm.Marshal(res.Msg)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println(string(toPrint))
}

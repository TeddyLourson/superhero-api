package main

import (
	"context"
	"log"
	"net/http"

	"buf.build/gen/go/digibear/digibear/connectrpc/go/usr/v1/usrv1connect"
	usrv1 "buf.build/gen/go/digibear/digibear/protocolbuffers/go/usr/v1"
	"connectrpc.com/connect"
	"google.golang.org/protobuf/encoding/protojson"
)

func main() {
	client := usrv1connect.NewUsrServiceClient(http.DefaultClient, "http://localhost:3000")
	res, err := client.GetAccountForProfileSelection(context.Background(), connect.NewRequest(&usrv1.GetAccountForProfileSelectionRequest{Id: "b03db4de-402f-479d-882a-2475090a4e6d"}))
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

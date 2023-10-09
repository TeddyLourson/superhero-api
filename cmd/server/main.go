package main

import (
	"net/http"

	"buf.build/gen/go/teddy-lourson/superhero/connectrpc/go/superhero/v1/superherov1connect"
	"github.com/digibearapp/digibear-api/modules"
	"github.com/rs/cors"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

func main() {
	superheroServer := modules.NewSuperheroServer()
	mux := http.NewServeMux()
	path, handler := superherov1connect.NewSuperheroServiceHandler(superheroServer)
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*", "*"},
		AllowedMethods:   []string{"HEAD", "GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: true,
		// AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token", "X-Custom-Header", "Connect-Protocol-Version"},
		// ExposedHeaders:   []string{"Link"},
		// AllowCredentials: false,
	})
	handler = c.Handler(handler)
	mux.Handle(path, handler)

	http.ListenAndServe(
		"localhost:3000",
		h2c.NewHandler(mux, &http2.Server{}),
	)
}

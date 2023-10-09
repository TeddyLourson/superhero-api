package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"buf.build/gen/go/digibear/digibear/connectrpc/go/usr/v1/usrv1connect"
	"github.com/digibearapp/digibear-api/modules"
	pkgstore "github.com/digibearapp/digibear-api/pkg/store"
	"github.com/joho/godotenv"
	"github.com/rs/cors"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	store, err := pkgstore.NewStorePgx(context.Background(), os.Getenv("USR_DATABASE_URL"))
	if err != nil {
		log.Fatalf("Error creating store: %v", err)
	}
	usrStore := modules.NewUsrStorePgx(*store)
	usrServer := modules.NewUsrServer(usrStore)
	mux := http.NewServeMux()
	path, handler := usrv1connect.NewUsrServiceHandler(usrServer)
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

gen:
	@buf generate
clean:
	@buf clean
server:
	@go run cmd/server/main.go
client:
	@go run cmd/client/main.go

update-connect:
	@go get buf.build/gen/go/digibear/digibear/connectrpc/go
update-proto:
	@go get buf.build/gen/go/digibear/digibear/protocolbuffers/go
update-buf: update-connect update-proto
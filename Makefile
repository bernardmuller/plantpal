run-plants-service:
	go run services/plants/main.go

gen-plants:
	@protoc \
		--proto_path=services "services/plants/plantspb/plants.proto" \
		--go_out=services --go_opt=paths=source_relative \
  		--go-grpc_out=services --go-grpc_opt=paths=source_relative
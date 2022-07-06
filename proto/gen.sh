protoc -I=. --go_out=gen/go --go_opt=paths=source_relative \
    --go-grpc_out=gen/go --go-grpc_opt=require_unimplemented_servers=false,paths=source_relative \
    trip.proto
protoc -I=. --grpc-gateway_out=paths=source_relative,grpc_api_configuration=trip.yaml:gen/go trip.proto
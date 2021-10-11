protoc -I=. --go_out=paths=source_relative:gen/go user.proto
protoc -I=. --go-grpc_out=paths=source_relative:gen/go user.proto
protoc -I=. --grpc-gateway_out=paths=source_relative,grpc_api_configuration=user.yaml:gen/go user.proto

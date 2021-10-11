protoc -I=. --go_out=paths=source_relative:gen/go paste.proto
protoc -I=. --go-grpc_out=paths=source_relative:gen/go paste.proto
protoc -I=. --grpc-gateway_out=paths=source_relative,grpc_api_configuration=paste.yaml:gen/go paste.proto

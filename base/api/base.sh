protoc -I=. --go_out=paths=source_relative:gen/go base.proto
protoc -I=. --go-grpc_out=paths=source_relative:gen/go base.proto

package main

import (
	"github.com/namsral/flag"
	"google.golang.org/grpc"
	"log"
	"pasteProject/server"
	userpb "pasteProject/user/api/gen/go"
	"pasteProject/user/user"
)

var userAddr = flag.String("user_addr", "localhost:8081", "address for user service")

func main() {
	flag.Parse()
	log.Fatal(server.RunGRPCServer(&server.GRPCConfig{
		Name: "user",
		Addr: *userAddr,
		RegisterFunc: func(s *grpc.Server) {
			userpb.RegisterUserServiceServer(s, &user.Service{})
		},
	}))
}

package main

import (
	"github.com/namsral/flag"
	"google.golang.org/grpc"
	"log"
	pastepb "pasteProject/paste/api/gen/go"
	"pasteProject/paste/paste"
	"pasteProject/server"
)

var pasteAddr = flag.String("paste_addr", "localhost:8082", "address for paste service")

func main() {
	flag.Parse()
	log.Fatal(server.RunGRPCServer(&server.GRPCConfig{
		Name: "Paste",
		Addr: *pasteAddr,
		RegisterFunc: func(s *grpc.Server) {
			pastepb.RegisterPasteServiceServer(s, &paste.Service{})
		},
	}))
}

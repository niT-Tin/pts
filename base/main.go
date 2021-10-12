package main

import (
	"github.com/namsral/flag"
	"google.golang.org/grpc"
	"log"
	"net"
	basepb "pasteProject/base/api/gen/go"
	"pasteProject/base/base"
)

var (
	envPath = flag.String("env", "/home/liuzehao/GolandProjects/pasteProject/base/conf/envs/db.env", "env file")
	//envAddr = flag.String("env_addr", "localhost:8084", "address for env loader")
)

func main() {
	flag.Parse()
	base.Load(*envPath)
	//fmt.Println(unix.Getpid() + unix.Getppid())
	lis, err := net.Listen("tcp", "localhost:8084")
	if err != nil {
		log.Fatalf("env starting: %v", err)
	}
	s := grpc.NewServer()
	basepb.RegisterBaseServiceServer(s, &base.Service{})
	err = s.Serve(lis)
	if err != nil {
		log.Fatalf("env servering: %v", err)
	}
}

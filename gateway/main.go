package main

import (
	"context"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/namsral/flag"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/encoding/protojson"
	"log"
	"net/http"
	"pasteProject/errs"
	pastepb "pasteProject/paste/api/gen/go"
	userpb "pasteProject/user/api/gen/go"
	"time"
)

var (
	addr      = flag.String("addr", ":8080", "address to listen")
	userAddr  = flag.String("user_addr", "localhost:8081", "address for user service")
	pasteAddr = flag.String("paste_addr", "localhost:8082", "address for paste service")
)

type serverConfig struct {
	name         string
	addr         string
	registerFunc func(ctx context.Context, mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) (err error)
}

func main() {
	flag.Parse()
	newErrs := errs.NewErrs(errs.GetDB())
	c := context.Background()
	c, cancel := context.WithCancel(c)
	defer cancel()
	mux := runtime.NewServeMux(runtime.WithMarshalerOption(
		runtime.MIMEWildcard, &runtime.JSONPb{
			MarshalOptions: protojson.MarshalOptions{
				UseEnumNumbers: true, // 将枚举常量转换为数字类型
			},
		},
	))
	var sfs = []serverConfig{
		{
			"user",
			*userAddr,
			userpb.RegisterUserServiceHandlerFromEndpoint,
		},
		{
			"paste",
			*pasteAddr,
			pastepb.RegisterPasteServiceHandlerFromEndpoint,
		},
	}

	for _, s := range sfs {
		err := s.registerFunc(c, mux, s.addr, []grpc.DialOption{grpc.WithInsecure()})
		if err != nil {
			var e = errs.Err{Message: "error registering service",
				When: time.Now(), Where: "gateway"}
			newErrs.ReciteErrors(e)
			log.Fatal(e)
		}
	}
	err := http.ListenAndServe(*addr, mux)
	if err != nil {
		var e = errs.Err{Message: "error starting gateway",
			When: time.Now(), Where: "gateway"}
		newErrs.ReciteErrors(e)
		log.Fatal(e)
	}
}

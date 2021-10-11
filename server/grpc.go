package server

import (
	"google.golang.org/grpc"
	"net"
	"pasteProject/errs"
	"time"
)

// GRPCConfig grpc配置结构
type GRPCConfig struct {
	Name         string                    // grpc服务名称
	Addr         string                    // grpc服务地址
	RegisterFunc func(server *grpc.Server) //注册grpc服务函数
}

func RunGRPCServer(c *GRPCConfig) error {
	handleErr := errs.NewErrs(errs.GetDB())
	lis, err := net.Listen("tcp", c.Addr)
	if err != nil {
		handleErr.ReciteErrors(errs.Err{Message: "Listen error on" + c.Addr, When: time.Now(), Where: "grpcRunner"})
		return err
	}
	// 此处还可以使用Interceptor 目前还不是很明白暂且不用
	s := grpc.NewServer()
	c.RegisterFunc(s)
	return s.Serve(lis)
}

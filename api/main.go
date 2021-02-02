package main

import (
	"github.com/asim/go-micro/plugins/registry/consul/v3"
	"github.com/asim/go-micro/plugins/server/http/v3"
	"github.com/asim/go-micro/v3"
	"github.com/asim/go-micro/v3/registry"
	"github.com/asim/go-micro/v3/server"
	"github.com/cloudjjcc/demo-services/api/handler"
	"github.com/cloudjjcc/demo-services/api/middleware"
	"time"
)

func main() {
	//etcdRegistry := etcd.NewRegistry(
	//	registry.Addrs("192.168.99.11:2379"),
	//)
	consulRegistry := consul.NewRegistry(
		registry.Addrs("j1900.home:8500"),
		registry.Timeout(1*time.Minute),
	)
	httpServer := http.NewServer(
		server.Address(":6001"),
		server.Name("api"),
		server.Registry(consulRegistry),
	)
	srv := micro.NewService(
		micro.Registry(consulRegistry),
		micro.RegisterTTL(30*time.Second),
		micro.RegisterInterval(10*time.Second),
		micro.Server(httpServer),
		micro.WrapClient(middleware.NewLogWrapper),
	)
	if err := httpServer.Handle(httpServer.NewHandler(handler.NewApiHandler(srv))); err != nil {
		panic(err)
	}
	srv.Init()
	if err := srv.Run(); err != nil {
		panic(err)
	}
}

package main

import (
	"github.com/asim/go-micro/plugins/registry/etcd/v3"
	"github.com/asim/go-micro/v3"
	"github.com/asim/go-micro/v3/registry"
)

func main() {
	etcdRegistry := etcd.NewRegistry(
		registry.Addrs("localhost:2379"),
	)
	srv := micro.NewService(
		micro.Name("demo-productor"),
		micro.Registry(etcdRegistry),
	)
	srv.Init()
	if err := srv.Run(); err != nil {
		panic(err)
	}
}

module github.com/cloudjjcc/demo-services/api

go 1.14

require (
	github.com/asim/go-micro/plugins/registry/consul/v3 v3.0.0-20210130181356-ca3dfc4580fa
	github.com/asim/go-micro/plugins/registry/etcd/v3 v3.0.0-20210125075117-bba3107ae13f
	github.com/asim/go-micro/plugins/server/http/v3 v3.0.0-20210130181356-ca3dfc4580fa
	github.com/asim/go-micro/v3 v3.0.1
	github.com/bmizerany/assert v0.0.0-20160611221934-b7ed37b82869 // indirect
	github.com/cloudjjcc/demo-services/product v0.0.0
	github.com/coreos/bbolt v1.3.3 // indirect
	github.com/coreos/go-systemd v0.0.0-20190719114852-fd7a80b32e1f // indirect
	github.com/cpuguy83/go-md2man/v2 v2.0.0 // indirect
	github.com/gin-gonic/gin v1.6.3
	github.com/golang/groupcache v0.0.0-20200121045136-8c9f03a8e57e // indirect
	github.com/golang/protobuf v1.4.3 // indirect
	github.com/google/uuid v1.2.0 // indirect
	github.com/gorilla/websocket v1.4.2 // indirect
	github.com/grpc-ecosystem/go-grpc-middleware v1.1.0 // indirect
	github.com/grpc-ecosystem/go-grpc-prometheus v1.2.0 // indirect
	github.com/grpc-ecosystem/grpc-gateway v1.9.5 // indirect
	github.com/jonboulle/clockwork v0.1.0 // indirect
	github.com/json-iterator/go v1.1.10 // indirect
	github.com/kr/pretty v0.2.0 // indirect
	github.com/miekg/dns v1.1.35 // indirect
	github.com/soheilhy/cmux v0.1.4 // indirect
	github.com/stretchr/testify v1.6.1 // indirect
	github.com/tmc/grpc-websocket-proxy v0.0.0-20200122045848-3419fae592fc // indirect
	github.com/xiang90/probing v0.0.0-20190116061207-43a291ad63a2 // indirect
	go.etcd.io/bbolt v1.3.5 // indirect
	golang.org/x/crypto v0.0.0-20201221181555-eec23a3978ad // indirect
	golang.org/x/net v0.0.0-20210119194325-5f4716e94777 // indirect
	golang.org/x/sys v0.0.0-20210124154548-22da62e12c0c // indirect
	golang.org/x/text v0.3.5 // indirect
	golang.org/x/time v0.0.0-20191024005414-555d28b269f0 // indirect
	google.golang.org/grpc v1.25.1 // indirect
	google.golang.org/protobuf v1.23.0 // indirect
	sigs.k8s.io/yaml v1.1.0 // indirect
)

replace (
	github.com/asim/go-micro/plugins/registry/memory/v3 v3.0.0-00010101000000-000000000000 => github.com/asim/go-micro/plugins/registry/memory/v3 v3.0.0-20210120213617-60010e82e2c7
	github.com/cloudjjcc/demo-services/product => ../product
)

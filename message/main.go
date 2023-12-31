package main

import (
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/registry/etcd"
	"message/conf"
	"message/core"
	"message/core/to_relation"
	proto "message/service"
	to_relation_proto "message/service/to_relation"
)

func main() {

	conf.Init()

	// etcd注册件
	etcdReg := etcd.NewRegistry(
		registry.Addrs("127.0.0.1:2379"),
	)
	// 得到一个微服务实例
	microService := micro.NewService(
		micro.Name("rpcMessageService"), // 微服务名字
		micro.Address("127.0.0.1:8091"),
		micro.Registry(etcdReg), // etcd注册件
		micro.Metadata(map[string]string{"protocol": "http"}),
	)

	// 结构命令行参数，初始化
	microService.Init()
	// 服务注册
	_ = proto.RegisterMessageServiceHandler(microService.Server(), new(core.MessageService))
	_ = to_relation_proto.RegisterToRelationServiceHandler(microService.Server(), new(to_relation.ToRelationService))
	// 启动微服务

	_ = microService.Run()

}

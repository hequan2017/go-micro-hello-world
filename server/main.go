package main

import (
	"context"
	"fmt"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-plugins/registry/etcdv3"
	"go-admin-micro/proto"
)

type Say struct {}

func (s *Say) Hello(ctx context.Context, req *model.SayParam, rsp *model.SayResponse) error {
	fmt.Println("received", req.Msg)
	rsp.Msg = "hello world"
	return nil
}


func main() {
	// 我这里用的etcd 做为服务发现，如果使用consul可以去掉
	reg := etcdv3.NewRegistry(func(op *registry.Options){
		op.Addrs = []string{
			"http://127.0.0.1:2379",
		}
	})
	// 初始化服务
	service := micro.NewService(
		micro.Name("lp.srv.eg1"),
		micro.Registry(reg),
	)

	// 如果你用的是consul把上面的注释掉用下面的
	/*
	// 初始化服务
	service := micro.NewService(
		micro.Name("lp.srv.eg1"),
	)
	 */
	service.Init()

	// 注册 Handler
	_ = model.RegisterSayHandler(service.Server(), new(Say))

	// run server
	if err := service.Run(); err != nil {
		panic(err)
	}
}
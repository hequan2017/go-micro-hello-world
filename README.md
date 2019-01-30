# hello world　go 微服务例子

## 软件
```
https://github.com/protocolbuffers/protobuf/releases
https://github.com/etcd-io/etcd/releases


根据系统环境 下载  protoc 和 etcd

启动etcd
```

## 生成  pb.go和micro.go文件

> bin\protoc.exe  --micro_out=. --go_out=. proto/common.proto

## 运行

> go run server/main.go


> go run client/main.go
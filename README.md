# go-crypto-exchange

基于 go-zero 微服务的加密货币交易所

## 项目结构

- grpc-common // grpc 公共服务
- ucenter // 用户中心服务
- ucenter-api // 用户中心 API 网关服务
- mscoin-common // 公共模块

## 项目内容

1、创建用户的 RPC 服务

```shell
# 生成 RPC 服务代码
goctl rpc protoc register.proto --go_out=./types --go-grpc_out=./types --zrpc_out=./register --style go_zero

protoc register.proto --go_out=./types --go-grpc_out=./types

protoc login.proto --go_out=./types --go-grpc_out=./types

# 创建用户中心服务 api 服务
goctl api new ucenterapi --style go_zero

# 同步数据库
goctl model mysql datasource --url="root:root@tcp(127.0.0.1:3306)/mscoin" --table="member" -c --dir .

# 注册
# 登录
```

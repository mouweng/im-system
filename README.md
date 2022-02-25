# IM-System即时通讯系统
[Golang](https://www.bilibili.com/video/BV1gf4y1r79E)
## 版本V1-构建基础Server
- 构建服务
```shell
go build -o server main.go server.go
./server
```
- 测试
```shell
nc 127.0.0.1 8888
```

## 版本V2-用户上线功能
- 构建服务
```shell
go build -o server main.go server.go user.go
./server
```
- 测试
```shell
nc 127.0.0.1 8888
```
```shell
nc 127.0.0.1 8888
```
```shell
nc 127.0.0.1 8888
```

## 版本V3-用户消息广播功能
- 构建服务
```shell
go build -o server main.go server.go user.go
./server
```
- 测试
```shell
nc 127.0.0.1 8888
hello
```
```shell
nc 127.0.0.1 8888
nihao
```
## 版本V4-用户业务封装
- 构建服务
```shell
go build -o server main.go server.go user.go
./server
```
- 测试
```shell
nc 127.0.0.1 8888
hello
```
```shell
nc 127.0.0.1 8888
nihao
```

## 版本V5-在线用户查询
- 构建服务
```shell
go build -o server main.go server.go user.go
./server
```
- 测试
```shell
nc 127.0.0.1 8888
who
```
```shell
nc 127.0.0.1 8888
who
```


## 版本V6-修改用户名
- 构建服务
```shell
go build -o server main.go server.go user.go
./server
```
- 测试
```shell
nc 127.0.0.1 8888
who
rename|yifanweng
who
```

## 版本V7-超时强踢功能
- 构建服务
```shell
go build -o server main.go server.go user.go
./server
```
- 测试
```shell
nc 127.0.0.1 8888

[127.0.0.1:52282]127.0.0.1:52282:已上线
你已超时，程序关闭...
```

## 版本V8-私聊功能
- 构建服务
```shell
go build -o server main.go server.go user.go
./server
```
- 测试
```shell
nc 127.0.0.1 8888
rename|yifanweng
to|wyf|hello
```
```shell
nc 127.0.0.1 8888
rename|wyf
to|yifanweng|hi
```

## 版本V9-新增客户端功能
- 构建服务
```shell
go build -o server main.go server.go user.go
./server
```
```shell
go build -o client client.go
./client -ip 127.0.0.1 -port 8888
```

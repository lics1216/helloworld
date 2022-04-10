## 第三次作业
按照自己的构想，写一个项目满足基本的目录结构和工程，代码需要包含对数据层、业务层、API 注册，以及 main 函数对于服务的注册和启动，信号处理，使用 Wire 构建依赖。可以使用自己熟悉的框架。

## 利用kratos 写一个增删改查的例子
增加一个demo 的模块，利用postman 发起post 请求 http://localhost:8000/demo，传入参数name。就会在数据库demo 表增加一条记录

```
post http://localhost:8000/demo
{
    "id": "34",
    "name": " lll"
}
```
暂时没操作数据库，和加入错误处理，@todo 
* 操作数据库，把增删改查其他接口补齐
* 错误处理


```
# Install Kratos
go install github.com/go-kratos/kratos/cmd/kratos/v2@latest
kratos new helloworld
go generate ./...
go build -o ./bin/ ./...

# 增加 demo.proto
kratos proto add api/helloworld/demo.proto
kratos proto client api/helloworld/demo.proto
kratos proto server api/helloworld/demo.proto -t internal/service
```

## 运行
```
go generate ./...
kratos run 
```



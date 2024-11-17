# 简介

作者：京城郭少

# 关于项目

> 微服务Demo程序

# 特点

* 参数解析：使用flag包
* 配置文件：使用toml格式
* 引入日志库：https://github.com/GuoFlight/glog
* 引入错误库：https://github.com/GuoFlight/gerror
* http服务器：采用Gin框架
* 包版本管理：go mod
* 接口鉴权
* 优雅退出
* 规范的TraceID
* 规范的目录结构

# 启动

```shell
go run main.go
go run main.go -c ./config.toml
```
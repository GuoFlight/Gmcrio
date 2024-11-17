# 简介

作者：京城郭少

# 关于项目

> 微服务Demo程序

# 特点

* 参数解析：使用flag包
* 配置文件：使用toml格式
* 日志引入glog库：https://github.com/GuoFlight/glog
* http服务器：采用Gin框架
* 包版本管理：go mod
* 接口鉴权
* 优雅退出
* 规范的TraceID
* 规范的目录结构

# 启动

```shell
go run main.go
go run main.go -c ./configDev.toml
```

# API

## 鉴权

```shell
# 获取Token
export umc_server="http://127.0.0.1:8080"
token=$(curl -s -XPOST ${umc_server}/v1/login -d "{\"username\":\"admin\",\"password\":\"admin\"}" -H "Content-Type: application/json" | jq -r .data)
# 访问需要鉴权的接口
curl -H "Authorization: ${token}" ${umc_server}/xxx
```

<br>
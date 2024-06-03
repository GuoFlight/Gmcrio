# 简介

作者：京城郭少

# 关于项目

> 微服务Demo程序

# 特点

配置文件

* 使用toml格式

日志

* 引入glog库：https://github.com/GuoFlight/glog

http服务器

* 采用Gin框架。

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
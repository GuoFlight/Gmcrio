前置操作
```shell
server="http://127.0.0.1:1216"
```

健康检查：

```shell
# 无需鉴权
curl ${server}/api/v1/health
```

# 鉴权

```shell
# 获取token
curl -X POST ${server}/api/v1/jwt/login -H 'Content-Type: application/json' -d '{"username":"user1","password":"123456"}'

# 访问需鉴权的接口
curl -H "Authorization: ${token}" ${server}/api/v1/admin/testAuth
```

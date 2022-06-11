# ucs-go-sdk

[![build](https://github.com/Macrow/ucs-go-sdk/actions/workflows/build.yml/badge.svg)](https://github.com/Macrow/ucs-go-sdk/actions/workflows/build.yml)
[![Go Reference](https://pkg.go.dev/badge/github.com/Macrow/ucs-go-sdk/v4.svg)](https://pkg.go.dev/github.com/Macrow/ucs-go-sdk)

用于集成```ucs```的开发包

## 快速开始

### 安装
```
go get -u github.com/Macrow/ucs-go-sdk
```

### 创建连接UCS的客户端
```
client := NewHttpClient("http://your.domain.com:port", yourAccessCode) // Http客户端
// client := NewHttpClient("https://your.domain.com:port", yourAccessCode) // Https客户端
client.SetUserToken(token)
client.SetClientIdAndSecret(clientId, clientSecret)
```

### UCS服务端验证Jwt
```
jwtUser, err := client.UserValidateJwt()
```

### UCS服务端验证操作码
```
err := client.UserValidatePermByOperation("UCS_O_CODE")
```

### UCS服务端验证接口
```
err := client.UserValidatePermByAction("ucs", "/api/v1/ucs/users", "get")
```

### UCS服务端验证用户是否拥有机构权限
```
err := client.UserValidatePermOrgById("org_id_is_here")
```

### 向UCS服务端发起应用级调用
```
res, err := client.ClientRequest("POST", "/api/v1/ucs/client/validate", nil)
```

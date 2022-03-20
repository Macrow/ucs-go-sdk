# ucs-go-sdk

[![build](https://github.com/Macrow/ucs-go-sdk/actions/workflows/build.yml/badge.svg)](https://github.com/Macrow/ucs-go-sdk/actions/workflows/build.yml)
[![Go Reference](https://pkg.go.dev/badge/github.com/Macrow/ucs-go-sdk/v4.svg)](https://pkg.go.dev/github.com/Macrow/ucs-go-sdk)

用于集成```ucs```的开发包

## 快速开始

### 安装
```
go get -u github.com/Macrow/ucs-go-sdk
```

### 验证Jwt
```
validator := NewValidator(rsaPublicKey)
jwtUser, err := validator.ValidateJwt(token)
```

### 创建连接UCS的客户端
```
client := NewRpcClient("your.domain.com:port") // Rpc客户端
// client := NewTLSRpcClient(certFile, "your.domain.com:port") // TLS连接下的Rpc客户端，需要UCS服务也同时开启
// client := NewHttpClient("https://your.domain.com:port", yourAccessCode) // Http客户端
client.SetToken(token)
```

### 如果令牌过期了，重新申请令牌
```
token, err := client.RenewToken()
```

### UCS服务端验证Jwt
```
err := client.ValidateJwt()
```

### UCS服务端验证操作码
```
err := client.ValidatePermOperationByCode("UCS_O_CODE")
```

### UCS服务端验证接口
```
err := client.ValidatePermAction("ucs", "/api/v1/ucs/users", "get")
```

### UCS服务端验证用户是否拥有机构权限
```
err := client.ValidatePermOrgById("org_id_is_here")
```

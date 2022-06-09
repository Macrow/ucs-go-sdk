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
client.SetToken(token)
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

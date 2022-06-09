package ucs

type RawJwtUser struct {
	Id         string  `json:"id"`   // 用户id
	Name       string  `json:"name"` // 用户登录名
	DeviceId   string  `json:"did"`  // 设备id
	DeviceName string  `json:"dn"`   // 设备名
	Issuer     string  `json:"iss"`  // 签发者
	IssueAt    float64 `json:"iat"`  // 签发时间
	ExpireAt   float64 `json:"exp"`  // 过期时间
}

type JwtUser struct {
	RawJwtUser
	Token string `json:"userToken"` // 令牌字符串
}

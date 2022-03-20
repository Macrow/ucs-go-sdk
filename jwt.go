package ucs

import (
	"errors"
	"github.com/golang-jwt/jwt/v4"
	"strings"
)

const (
	JwtTokenClaimsId         = "id"
	JwtTokenClaimsName       = "name"
	JwtTokenClaimsDeviceId   = "did"
	JwtTokenClaimsDeviceName = "dn"
	JwtTokenClaimsIssuer     = "iss"
	JwtTokenClaimsIssueAt    = "iat"
	JwtTokenClaimsExpireAt   = "exp"
)

const (
	JwtKeySplitter = "_"
	JwtErrInternal = "内部错误"
	JwtErrFormat   = "令牌格式错误"
	JwtErrVersion  = "令牌版本错误"
)

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
	Token string `json:"token"` // 令牌字符串
}

// GetUserJwtCacheKey 获取Jwt存储的key
func GetUserJwtCacheKey(prefix, userId, deviceId string) string {
	return strings.Join([]string{prefix, userId, deviceId}, JwtKeySplitter)
}

func GenerateJwt(privateKey []byte, id, username, deviceId, deviceName, issuer string, issueAt float64, expireAt float64) (jwtUser *JwtUser, err error) {
	key, err := jwt.ParseRSAPrivateKeyFromPEM(privateKey)
	if err != nil {
		return nil, err
	}

	rawToken := jwt.New(jwt.SigningMethodRS256)
	claims := rawToken.Claims.(jwt.MapClaims)
	claims[JwtTokenClaimsId] = id
	claims[JwtTokenClaimsName] = username
	claims[JwtTokenClaimsDeviceId] = deviceId
	claims[JwtTokenClaimsDeviceName] = deviceName
	claims[JwtTokenClaimsIssuer] = issuer
	claims[JwtTokenClaimsIssueAt] = issueAt
	claims[JwtTokenClaimsExpireAt] = expireAt

	token, err := rawToken.SignedString(key)
	if err != nil {
		return nil, err
	}
	jwtUser = &JwtUser{
		RawJwtUser: RawJwtUser{
			Id:         id,
			Name:       username,
			DeviceId:   deviceId,
			DeviceName: deviceName,
			Issuer:     issuer,
			IssueAt:    float64(issueAt),
			ExpireAt:   float64(expireAt),
		},
		Token: token,
	}

	return jwtUser, nil
}

func ValidateJwt(publicKey []byte, tokenString string) (*JwtUser, error) {
	key, err := jwt.ParseRSAPublicKeyFromPEM(publicKey)
	if err != nil {
		return nil, errors.New(JwtErrInternal)
	}
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, errors.New(JwtErrFormat)
		}
		return key, nil
	})
	if err != nil || token == nil || !token.Valid {
		return nil, errors.New(JwtErrFormat)
	}
	// 解析令牌并存储为JwtUser格式
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, errors.New(JwtErrFormat)
	}

	if claims[JwtTokenClaimsId] == nil ||
		claims[JwtTokenClaimsName] == nil ||
		claims[JwtTokenClaimsDeviceId] == nil ||
		claims[JwtTokenClaimsDeviceName] == nil ||
		claims[JwtTokenClaimsIssuer] == nil ||
		claims[JwtTokenClaimsIssueAt] == nil ||
		claims[JwtTokenClaimsExpireAt] == nil {
		return nil, errors.New(JwtErrVersion)
	}

	jwtUser := &JwtUser{
		RawJwtUser: RawJwtUser{
			Id:         claims[JwtTokenClaimsId].(string),
			Name:       claims[JwtTokenClaimsName].(string),
			DeviceId:   claims[JwtTokenClaimsDeviceId].(string),
			DeviceName: claims[JwtTokenClaimsDeviceName].(string),
			Issuer:     claims[JwtTokenClaimsIssuer].(string),
			IssueAt:    claims[JwtTokenClaimsIssueAt].(float64),
			ExpireAt:   claims[JwtTokenClaimsExpireAt].(float64),
		},
		Token: tokenString,
	}

	return jwtUser, nil
}

func IsSame(u1 *RawJwtUser, u2 *RawJwtUser) bool {
	if strings.Compare(u1.Id, u2.Id) == 0 &&
		strings.Compare(u1.Name, u2.Name) == 0 &&
		strings.Compare(u1.DeviceId, u2.DeviceId) == 0 &&
		strings.Compare(u1.DeviceName, u2.DeviceName) == 0 &&
		strings.Compare(u1.Issuer, u2.Issuer) == 0 &&
		u1.IssueAt == u2.IssueAt &&
		u1.ExpireAt == u2.ExpireAt {
		return true
	}
	return false
}

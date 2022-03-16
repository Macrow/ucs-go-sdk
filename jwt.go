package ucs

import (
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
)

const (
	JwtTokenClaimsId       = "id"
	JwtTokenClaimsName     = "name"
	JwtTokenClaimsDeviceId = "did"
	JwtTokenClaimsIssuer   = "iss"
	JwtTokenClaimsIssueAt  = "iat"
	JwtTokenClaimsExpireAt = "exp"
)

type JwtUser struct {
	Id       string  `json:"id"`   // 用户id
	Name     string  `json:"name"` // 用户登录名
	DeviceId string  `json:"did"`  // 用户登录名
	Issuer   string  `json:"iss"`  // 签发者
	IssueAt  float64 `json:"iat"`  // 签发时间
	ExpireAt float64 `json:"exp"`  // 过期时间
}

func ValidateJwt(publicKey []byte, tokenString string) (token *jwt.Token, user *JwtUser, err error) {
	key, err := jwt.ParseRSAPublicKeyFromPEM(publicKey)
	if err != nil {
		return
	}
	token, err = jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("unexpected sign method: %v", t.Method.Alg())
		}
		return key, nil
	})
	if err != nil || token == nil || !token.Valid {
		return
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		err = errors.New("jwt format error")
		return
	}
	user = &JwtUser{
		Id:       claims[JwtTokenClaimsId].(string),
		Name:     claims[JwtTokenClaimsName].(string),
		DeviceId: claims[JwtTokenClaimsDeviceId].(string),
		Issuer:   claims[JwtTokenClaimsIssuer].(string),
		IssueAt:  claims[JwtTokenClaimsIssueAt].(float64),
		ExpireAt: claims[JwtTokenClaimsExpireAt].(float64),
	}
	return
}

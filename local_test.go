package ucs

import (
	"fmt"
	"testing"
)

const token = `eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJkaWQiOiJhZG1pbl93ZWIiLCJkbiI6IkNocm9tZSIsImV4cCI6MTY4NzA1ODI1OSwiaWF0IjoxNjU1NTIyMjU5LCJpZCI6ImNhaDFrOHV2OW1jNnU1dTdmaWNnIiwiaXNzIjoidWNzIiwibmFtZSI6InJvb3QifQ.m2uOt7IlZpfng_UhBM2aeVETjhABp0sreAeqgJRT6QejXhaogNY3qXjr-ANi_oXqsVkA0Tof3z2qCMwl0mrHc5WEHXPvCRr_gOJ184z10Lf1z6cxaaQ4gt1R3TlCHst3DIlyl4iRAstLjfnlmm3aTWYZMjK-d3FXKA6i2yWZAXMInEoijpNMlYFGaojFfEZjlTPTp_Lmj4Spus7s8f_AjvckUJfYcymvRJHR9M7YEgRq2Lu_E-y4IsCGt9PphDah12JFv8-qg6UWFheiNIgg5rcQ0KKZcal73wpm9tmVEpJbn8SBsRV_tMfIOvjC8Vvbfh_-DoYWD3ZNtivrd8VMbg`
const clientToken = "d3NUREp6Z0FLZ0AxMjM0NTY="
const clientId = "wsTDJzgAKg"
const clientSecret = "123456"

func testByClient(client Client) {
	jwtUser, err := client.UserValidateJwt()
	if err != nil {
		fmt.Println(err)
	}
	if jwtUser != nil {
		fmt.Println(jwtUser)
	}
	ok, err := client.ClientValidate(ClientAuthKindToken)
	if ok {
		fmt.Println("应用端鉴权成功")
	} else {
		fmt.Println(err)
	}
	ok, err = client.ClientValidate(ClientAuthKindIdAndSecret)
	if ok {
		fmt.Println("应用端鉴权成功")
	} else {
		fmt.Println(err)
	}

	res, err := client.UserValidatePermByOperation("不存在的操作", true)
	fmt.Println(res.User)
	fmt.Println(res.Permit)
	fmt.Println(err)

	res, err = client.UserValidatePermByOperation("UCS_USER_LIST", true)
	fmt.Println(res.User)
	fmt.Println(res.Permit)
	fmt.Println(err)

	res, err = client.UserValidatePermByAction("ucs", "GET", "/api/v1/ucs/users", true)
	fmt.Println(res.User)
	fmt.Println(res.Permit)
	fmt.Println(err)

	userRes, err := client.UserRequest("GET", "/api/v1/ucs/users?pageSize=1", nil)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(userRes)
	}

	_, err = client.ClientRequest("GET", "/api/v1/ucs/client/validate", nil, ClientAuthKindToken)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("应用鉴权成功")
	}

	_, err = client.ClientRequest("GET", "/api/v1/ucs/client/validate", nil, ClientAuthKindIdAndSecret)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("应用鉴权成功")
	}
}

func TestHttp(t *testing.T) {
	client := NewHttpClient("http://localhost:8019", "1A2B3C4D")
	client.SetUserToken(token)
	client.SetClientToken(clientToken)
	client.SetClientIdAndSecret(clientId, clientSecret)
	testByClient(client)
}

func TestTlsHttp(t *testing.T) {
	client := NewHttpClient("https://localhost:8019", "1A2B3C4D")
	client.SetUserToken(token)
	client.SetClientToken(clientToken)
	client.SetClientIdAndSecret(clientId, clientSecret)
	testByClient(client)
}

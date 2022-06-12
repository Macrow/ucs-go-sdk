package ucs

import (
	"fmt"
	"testing"
)

const token = `eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJkaWQiOiJhZG1pbl93ZWIiLCJkbiI6IkNocm9tZSIsImV4cCI6MTY4NjQyMzE3MCwiaWF0IjoxNjU0ODg3MTcwLCJpZCI6ImNhaDFrOHV2OW1jNnU1dTdmaWNnIiwiaXNzIjoidWNzIiwibmFtZSI6InJvb3QifQ.IhgvqpWe9TJvSm1x39HH0LSiKwoZp1ge6GQgDOSKKcbAzArUEFaKJfpoJQUCJVJeq-I8TpUVSEjdwRh8Hty03L0G79POlqb87u-hzh29RmfP9tFNPY565Zm9GyB0kybiWA68ZQriDiTZaUEk1K2N4sq85HIpArV04haSvE9lJ46v2wrNprcVRxjWFWWxAt1qeBZFPuUtFk93A1OIWn2PbxE_fmlE1qVjqwukpanIKR9y3O2geC4F4-ed9qA8VZl0N8IHjMLABE-oIPa0Tlvt9tVoJ1sx0LqlA5GphZHXARDzgr2hdytuE_OxJeyULkadKVvqMZgeNRnwL404DoSx-Q`
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

	clientRes, err := client.ClientRequest("POST", "/api/v1/ucs/client/validate", nil)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(clientRes)
	}
}

func TestHttp(t *testing.T) {
	client := NewHttpClient("http://localhost:8019", "1A2B3C4D")
	client.SetUserToken(token)
	client.SetClientIdAndSecret(clientId, clientSecret)
	testByClient(client)
}

func TestTlsHttp(t *testing.T) {
	client := NewHttpClient("https://localhost:8019", "1A2B3C4D")
	client.SetUserToken(token)
	client.SetClientIdAndSecret(clientId, clientSecret)
	testByClient(client)
}

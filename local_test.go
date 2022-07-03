package ucs

import (
	"fmt"
	"testing"
)

const token = `eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJkaWQiOiJhZG1pbl93ZWIiLCJkbiI6IkNocm9tZSIsImV4cCI6MTY4ODM3NTYxMiwiaWF0IjoxNjU2ODM5NjEyLCJpZCI6ImNiMGxyYnV2OW1jNzE5NzY0ZXBnIiwiaXNzIjoidWNzIiwibmFtZSI6InJvb3QifQ.4IQ5Ewy6FCB8cs2gWulS57iSC7AVUr5B4klNXOSYRof0yX3V4UktrVV1SX9mlhv3oc3Js_tLY9CtPizX8f5yGlWlkjyRZYrg0ueKOFnquRrsF3n7SwqIMCVDRxD9ale1vxxn4aSL8H-ZH3yXzXoqIy-dJPYqqmlO362L0WfxT6jyRLzVTr7pis9MZuircPJVnC5HTvKL4_Qb2V_7zvC3mly1s6lEnlXQ4waTsRrCzsh57px19YluWDJY_jlXFKTBdu6Mot11CNz_n0UpRSsUaQYr9BuhGpEauvjpAPqYr3JTDzDxl02OJ-OrCjTV_E4N0lxRlIS57uqiFSXTzNlXtA`
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

	res, err := client.UserValidatePermByOperation("不存在的操作", true, true)
	fmt.Println(res.User)
	fmt.Println(res.Permit)
	fmt.Println(res.OrgIds)
	fmt.Println(err)

	res, err = client.UserValidatePermByOperation("UCS_USER_LIST", true, true)
	fmt.Println(res.User)
	fmt.Println(res.Permit)
	fmt.Println(res.OrgIds)
	fmt.Println(err)

	res, err = client.UserValidatePermByAction("ucs", "GET", "/api/v1/ucs/users", true, true)
	fmt.Println(res.User)
	fmt.Println(res.Permit)
	fmt.Println(res.OrgIds)
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

package ucs

import (
	"fmt"
	"testing"
)

const token = `eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJkaWQiOiJhZG1pbl93ZWIiLCJkbiI6IkNocm9tZSIsImV4cCI6MTY4NjMyNjcwMywiaWF0IjoxNjU0NzkwNzAzLCJpZCI6ImNhaDFrOHV2OW1jNnU1dTdmaWNnIiwiaXNzIjoidWNzIiwibmFtZSI6InJvb3QifQ.GtGvfltbGmV79SWoxaPX6dYrTyaHGLak_Zg3D7PfujJWMDBi5R8s0POS2TRm7LNFZxUeqRancjj9EGPnKdWsw9oH_nCPBjVhF_YY0U9CqtMnI6WAIrtIt9ouOJxfIW_TmJQumHzaqrclULAoL-_-LgoKJFiLuHhcOtsuinK0eH0UHsF7ruW0YY1a1E3pg6gKVjom17Y1V1RaLjDQpirqojfQtqZjgpaaa2IRBMSvmtLXfG1BFAIQd_SSr3EqugxQItkp5rQdMJzNWHhHn041pmE2dHXA1n7UmBha9z1q2jo8u4EkPzUYD2AxRJVs3k1GiesByGR_UOyrVbmS1ojYFQ`
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

	err = client.UserValidatePermOperationByCode("不存在的操作")
	fmt.Println(err)

	err = client.UserValidatePermOperationByCode("UCS_USER_LIST")
	fmt.Println(err)

	err = client.UserValidatePermAction("ucs", "/api/v1/ucs/users", "get")
	fmt.Println(err)

	err = client.UserValidatePermOrgById("rererwerw")
	fmt.Println(err)

	err = client.UserValidatePermOrgById("c8fjca649b3hbmov5n60")
	fmt.Println(err)

	err = client.UserValidatePermActionWithOrgId("ucs", "/api/v1/ucs/users", "get", "c8fjca649b3hbmov5n60")
	fmt.Println(err)

	err = client.UserValidatePermActionWithOrgId("ucs", "/api/v1/ucs/users", "get", "234sdfsdja")
	fmt.Println(err)

	res, err := client.UserQueryOrgIdsByAction("ucs", "/api/v1/ucs/users", "get")
	fmt.Println(err)
	fmt.Println(res.orgPermissionType)
	fmt.Println(res.orgIds)

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

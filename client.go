package ucs

type Client interface {
	SetTimeout(timeout int) Client
	SetUserToken(userToken string) Client
	SetClientIdAndSecret(clientId, clientSecret string) Client
	SetHttpHeaderNames(accessCodeHeader, randomKeyHeader, userTokenHeader, clientTokenHeader string) Client

	UserValidateJwt() (*JwtUser, error)
	UserValidatePermOperationByCode(operationCode string) error
	UserValidatePermAction(service, path, method string) error
	UserValidatePermOrgById(orgId string) error
	UserValidatePermActionWithOrgId(service, path, method, orgId string) error
	UserQueryOrgIdsByAction(service, path, method string) (*ActionOrgIds, error)

	ClientRequest(method, url string, data map[string]string) (interface{}, error)
}

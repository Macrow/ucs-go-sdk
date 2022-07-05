package ucs

type Client interface {
	SetTimeout(timeout int) Client
	SetBaseUrl(baseUrl string) Client
	SetAccessCode(accessCode string) Client
	SetRandomKey(randomKey string) Client
	SetUserToken(userToken string) Client
	SetClientToken(clientToken string) Client
	SetClientIdAndSecret(clientId, clientSecret string) Client
	SetHttpHeaderNames(accessCodeHeader, randomKeyHeader, userTokenHeader, clientTokenHeader string) Client

	UserValidateJwt() (*JwtUser, error)
	ClientValidate(clientAuthKind ClientAuthKind) (bool, error)
	UserValidatePermByOperation(code string, fulfillJwt bool, fulfillOrgIds bool) (*PermitResult, error)
	UserValidatePermByAction(service, method, path string, fulfillJwt bool, fulfillOrgIds bool) (*PermitResult, error)

	UserRequest(method, url string, data map[string]string) (interface{}, error)
	ClientRequest(method, url string, data map[string]string, clientAuthKind ClientAuthKind) (interface{}, error)
}

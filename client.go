package ucs

type Client interface {
	SetTimeout(timeout int) Client
	SetToken(token string) Client
	SetHttpHeaderNames(accessCodeHeader, randomKeyHeader string) Client
	ValidateJwt() error
	ValidatePermOperationByCode(operationCode string) error
	ValidatePermAction(service, path, method string) error
	ValidatePermOrgById(orgId string) error
}

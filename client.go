package ucs

type Client interface {
	SetTimeout(timeout int) Client
	SetToken(token string) Client
	SetHttpHeaderNames(accessCodeHeader, randomKeyHeader string) Client

	ValidateJwt() error
	ValidatePermOperationByCode(operationCode string) error
	ValidatePermAction(service, path, method string) error
	ValidatePermOrgById(orgId string) error
	ValidatePermActionWithOrgId(service, path, method, orgId string) error
	QueryOrgIdsByAction(service, path, method string) (*ActionOrgIds, error)

	OAuth2TokenByAuthorizationCode(code, clientId, clientSecret, deviceId, deviceName string) (string, error)
	OAuth2TokenByPassword(username, password, deviceId, deviceName string) (string, error)
}

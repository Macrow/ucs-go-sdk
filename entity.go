package ucs

type NormalHttpResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Result  interface{} `json:"result"`
}

type ValidateJwtHttpResponse struct {
	Code    int        `json:"code"`
	Message string     `json:"message"`
	Result  RawJwtUser `json:"result"`
}

type CommonPermitResult struct {
	Permit bool `json:"permit"`
}

type PermitHttpResponse struct {
	Code    int                `json:"code"`
	Message string             `json:"message"`
	Result  CommonPermitResult `json:"result"`
}

type CurrentQueryActionOrgIdsRes struct {
	OrgPermissionType string   `json:"orgPermissionType"`
	OrgIds            []string `json:"orgIds"`
}

type QueryActionOrgIdsHttpResponse struct {
	Code    int                         `json:"code"`
	Message string                      `json:"message"`
	Result  CurrentQueryActionOrgIdsRes `json:"result"`
}

type OAuth2TokenRes struct {
	AccessToken string `json:"accessToken"`
}

type OAuth2TokenResponse struct {
	Code    int            `json:"code"`
	Message string         `json:"message"`
	Result  OAuth2TokenRes `json:"result"`
}

type CommonRenewTokenResult struct {
	Token string `json:"userToken"`
}

type RenewTokenHttpResponse struct {
	Code    int                    `json:"code"`
	Message string                 `json:"message"`
	Result  CommonRenewTokenResult `json:"result"`
}

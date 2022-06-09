package ucs

type OrgPermissionType string

type ActionOrgIds struct {
	orgPermissionType OrgPermissionType
	orgIds            []string
}

const (
	DefaultHeaderRandomKey   = "Random-Key"
	DefaultHeaderAccessCode  = "Access-Code"
	DefaultHeaderUserToken   = "Authorization"
	DefaultHeaderClientToken = "Client-Authorization"
	DefaultNoPermMsg         = "权限不足"
	DefaultTimeout           = 3

	ValidateJwtURL                 = "/api/v1/ucs/current/jwt"
	ValidatePermOperationByCodeURL = "/api/v1/ucs/current/check-operation"
	ValidatePermActionURL          = "/api/v1/ucs/current/check-action"
	ValidatePermOrgByIdURL         = "/api/v1/ucs/current/check-org"
	ValidatePermActionWithOrgIdURL = "/api/v1/ucs/current/check-action-with-org-id"
	QueryOrgIdsByActionURL         = "/api/v1/ucs/current/query-action-org-ids"
)

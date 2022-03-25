package ucs

type OrgPermissionType string

type ActionOrgIds struct {
	orgPermissionType OrgPermissionType
	orgIds            []string
}

const (
	OrgPermissionTypeTree   OrgPermissionType = "tree"
	OrgPermissionTypeSelf   OrgPermissionType = "self"
	OrgPermissionTypeNone   OrgPermissionType = "none"
	OrgPermissionTypeAll    OrgPermissionType = "all"
	OrgPermissionTypeCustom OrgPermissionType = "custom"

	DefaultHeaderRandomKey  = "Random-Key"
	DefaultHeaderAccessCode = "Access-Code"
	DefaultNoPermMsg        = "权限不足"
	DefaultTimeout          = 3

	ValidateJwtURL                 = "/api/v1/ucs/current/blank"
	ValidatePermOperationByCodeURL = "/api/v1/ucs/current/check-operation"
	ValidatePermActionURL          = "/api/v1/ucs/current/check-action"
	ValidatePermOrgByIdURL         = "/api/v1/ucs/current/check-org"
	ValidatePermActionWithOrgIdURL = "/api/v1/ucs/current/check-action-with-org-id"
	QueryOrgIdsByActionURL         = "/api/v1/ucs/current/query-action-org-ids"
)

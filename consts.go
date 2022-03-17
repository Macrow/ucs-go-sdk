package ucs

const (
	DefaultHeaderRandomKey  = "Random-Key"
	DefaultHeaderAccessCode = "Access-Code"
	DefaultNoPermMsg        = "权限不足"
	DefaultTimeout          = 3

	ValidateJwtURL                 = "/api/v1/ucs/current/blank"
	ValidatePermOperationByCodeURL = "/api/v1/ucs/current/check-operation"
	ValidatePermActionURL          = "/api/v1/ucs/current/check-action"
	ValidatePermOrgByIdURL         = "/api/v1/ucs/current/check-org"
)

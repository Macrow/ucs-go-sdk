package ucs

type RequestKind string

const (
	USER                     RequestKind = "USER"
	CLIENT                   RequestKind = "CLIENT"
	DefaultHeaderRandomKey               = "Random-Key"
	DefaultHeaderAccessCode              = "Access-Code"
	DefaultHeaderUserToken               = "Authorization"
	DefaultHeaderClientToken             = "Client-Authorization"
	DefaultNoPermMsg                     = "权限不足"
	DefaultTimeout                       = 3

	ValidateJwtURL             = "/api/v1/ucs/current/jwt"
	ValidatePermByOperationURL = "/api/v1/ucs/current/check-operation"
	ValidatePermByActionURL    = "/api/v1/ucs/current/check-action"
)

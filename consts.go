package ucs

type RequestKind string
type ClientAuthKind string

const (
	RequestKindUser           RequestKind    = "RequestKindUser"
	RequestKindClient         RequestKind    = "RequestKindClient"
	ClientAuthKindIdAndSecret ClientAuthKind = "ClientAuthKindIdAndSecret"
	ClientAuthKindToken       ClientAuthKind = "ClientAuthKindToken"
	ClientAuthKindNone        ClientAuthKind = "ClientAuthKindNone"
	DefaultHeaderRandomKey                   = "Random-Key"
	DefaultHeaderAccessCode                  = "Access-Code"
	DefaultHeaderUserToken                   = "Authorization"
	DefaultHeaderClientToken                 = "Client-Authorization"
	DefaultNoPermMsg                         = "权限不足"
	DefaultTimeout                           = 3

	ValidateJwtURL             = "/api/v1/ucs/current/jwt"
	ValidatePermByOperationURL = "/api/v1/ucs/current/check-operation"
	ValidatePermByActionURL    = "/api/v1/ucs/current/check-action"
)

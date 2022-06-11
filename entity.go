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

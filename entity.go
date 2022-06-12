package ucs

type HttpResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Result  interface{} `json:"result"`
}

type ValidateJwtHttpResponse struct {
	Code    int        `json:"code"`
	Message string     `json:"message"`
	Result  RawJwtUser `json:"result"`
}

type PermitResult struct {
	Permit bool     `json:"permit"`
	User   *JwtUser `json:"user"`
}

type PermitHttpResponse struct {
	Code    int          `json:"code"`
	Message string       `json:"message"`
	Result  PermitResult `json:"result"`
}

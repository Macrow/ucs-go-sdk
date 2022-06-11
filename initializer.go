package ucs

func NewHttpClient(baseUrl string, accessCode string) Client {
	client := &HttpUcsClient{
		baseUrl:           baseUrl,
		accessCode:        accessCode,
		timeout:           DefaultTimeout,
		accessCodeHeader:  DefaultHeaderAccessCode,
		randomKeyHeader:   DefaultHeaderRandomKey,
		userTokenHeader:   DefaultHeaderUserToken,
		clientTokenHeader: DefaultHeaderClientToken,
	}
	return client
}

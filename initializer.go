package ucs

func NewHttpClient(baseUrl string, accessCode string) Client {
	client := &HttpClient{
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

package ucs

import (
	"encoding/base64"
	"errors"
	"fmt"
	"github.com/imroc/req/v3"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type HttpUcsClient struct {
	timeout           int
	baseUrl           string
	accessCode        string
	randomKey         string
	userToken         string
	clientToken       string
	clientId          string
	clientSecret      string
	accessCodeHeader  string
	randomKeyHeader   string
	userTokenHeader   string
	clientTokenHeader string
	agent             *req.Client
}

func (c *HttpUcsClient) SetTimeout(timeout int) Client {
	if timeout > 0 {
		c.timeout = timeout
	}
	return c
}

func (c *HttpUcsClient) SetBaseUrl(baseUrl string) Client {
	c.baseUrl = baseUrl
	return c
}

func (c *HttpUcsClient) SetAccessCode(accessCode string) Client {
	c.accessCode = accessCode
	return c
}

func (c *HttpUcsClient) SetRandomKey(randomKey string) Client {
	c.randomKey = randomKey
	return c
}

func (c *HttpUcsClient) SetUserToken(userToken string) Client {
	c.userToken = userToken
	return c
}

func (c *HttpUcsClient) SetClientToken(clientToken string) Client {
	c.clientToken = clientToken
	return c
}

func (c *HttpUcsClient) SetClientIdAndSecret(clientId, clientSecret string) Client {
	c.clientId = clientId
	c.clientSecret = clientSecret
	return c
}

func (c *HttpUcsClient) SetHttpHeaderNames(accessCodeHeader, randomKeyHeader, userTokenHeader, clientTokenHeader string) Client {
	if len(accessCodeHeader) > 0 {
		c.accessCodeHeader = accessCodeHeader
	}
	if len(randomKeyHeader) > 0 {
		c.randomKeyHeader = randomKeyHeader
	}
	if len(userTokenHeader) > 0 {
		c.userTokenHeader = userTokenHeader
	}
	if len(userTokenHeader) > 0 {
		c.clientTokenHeader = clientTokenHeader
	}
	return c
}

func (c *HttpUcsClient) UserValidateJwt() (*JwtUser, error) {
	a, err := c.getUserAgent()
	if err != nil {
		return nil, err
	}
	result := &ValidateJwtHttpResponse{}
	res, err := a.R().SetResult(result).Get(ValidateJwtURL)
	if err != nil {
		return nil, err
	}
	if !res.IsSuccess() {
		return nil, fmt.Errorf("error: %v", res.Error())
	}
	if result.Code == 0 {
		return &JwtUser{RawJwtUser: result.Result, Token: c.userToken}, nil
	}
	return nil, errors.New(result.Message)
}

func (c *HttpUcsClient) ClientValidate(clientAuthKind ClientAuthKind) (bool, error) {
	a, err := c.getClientAgent(clientAuthKind)
	if err != nil {
		return false, err
	}
	result := &HttpResponse{}
	res, err := a.R().SetResult(result).Get(ValidateClientURL)
	if err != nil {
		return false, err
	}
	if !res.IsSuccess() {
		return false, fmt.Errorf("error: %v", res.Error())
	}
	if result.Code == 0 {
		return true, nil
	}
	return false, errors.New(result.Message)
}

func (c *HttpUcsClient) UserValidatePermByOperation(operationCode string, fulfillJwt bool, fulfillOrgIds bool) (*PermitResult, error) {
	fulfillJwtParam := "0"
	if fulfillJwt {
		fulfillJwtParam = "1"
	}
	fulfillOrgIdsParam := "0"
	if fulfillOrgIds {
		fulfillOrgIdsParam = "1"
	}
	return c.permitPost(ValidatePermByOperationURL, map[string]string{
		"code":          operationCode,
		"fulfillJwt":    fulfillJwtParam,
		"fulfillOrgIds": fulfillOrgIdsParam,
	})
}

func (c *HttpUcsClient) UserValidatePermByAction(service, method, path string, fulfillJwt bool, fulfillOrgIds bool) (*PermitResult, error) {
	fulfillJwtParam := "0"
	if fulfillJwt {
		fulfillJwtParam = "1"
	}
	fulfillOrgIdsParam := "0"
	if fulfillOrgIds {
		fulfillOrgIdsParam = "1"
	}
	return c.permitPost(ValidatePermByActionURL, map[string]string{
		"service":       service,
		"method":        method,
		"path":          path,
		"fulfillJwt":    fulfillJwtParam,
		"fulfillOrgIds": fulfillOrgIdsParam,
	})
}

func (c *HttpUcsClient) UserRequest(method, url string, data map[string]string) (interface{}, error) {
	return c.genericRequest(RequestKindUser, method, url, data, ClientAuthKindNone)
}

func (c *HttpUcsClient) ClientRequest(method, url string, data map[string]string, clientAuthKind ClientAuthKind) (interface{}, error) {
	return c.genericRequest(RequestKindClient, method, url, data, clientAuthKind)
}

func (c *HttpUcsClient) genericRequest(kind RequestKind, method, url string, data map[string]string, clientAuthKind ClientAuthKind) (interface{}, error) {
	method = strings.ToUpper(method)
	switch method {
	case http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete, http.MethodPatch:
	default:
		return nil, fmt.Errorf("unsupport method: %v", method)
	}
	var a *req.Client
	var err error
	switch kind {
	case RequestKindUser:
		a, err = c.getUserAgent()
	case RequestKindClient:
		a, err = c.getClientAgent(clientAuthKind)
	default:
		return nil, fmt.Errorf("unsupport request kind: %v", kind)
	}
	if err != nil {
		return nil, err
	}
	result := &HttpResponse{}
	request := a.R().SetResult(result)
	if data != nil {
		request = request.SetFormData(data)
	}
	res, err := request.Send(method, url)
	if err != nil {
		return nil, err
	}
	if !res.IsSuccess() {
		return nil, fmt.Errorf("error: %v", res.Error())
	}
	if result.Code != 0 {
		return nil, fmt.Errorf(result.Message)
	}
	return result.Result, nil
}

func (c *HttpUcsClient) permitPost(url string, data map[string]string) (*PermitResult, error) {
	a, err := c.getUserAgent()
	if err != nil {
		return nil, err
	}
	result := &PermitHttpResponse{}
	res, err := a.R().SetResult(result).SetFormData(data).Post(url)
	if err != nil {
		return nil, err
	}
	if !res.IsSuccess() {
		return nil, fmt.Errorf("error: %v", res.Error())
	}
	if result.Result.User != nil {
		result.Result.User.Token = c.userToken
	}
	return &result.Result, nil
}

func (c *HttpUcsClient) initAgent() {
	if c.agent == nil {
		c.agent = req.C().
			DisableAutoDecode().
			SetAutoDecodeContentType("application/json").
			SetBaseURL(c.baseUrl)
	}
	c.agent.
		SetTimeout(time.Duration(c.timeout)*time.Second).
		SetCommonHeader(c.accessCodeHeader, c.accessCode).
		SetCommonHeader(c.randomKeyHeader, c.randomKey)
}

func (c *HttpUcsClient) getUserAgent() (*req.Client, error) {
	c.initAgent()
	if len(c.userToken) == 0 {
		return nil, errors.New("please provide userToken")
	}
	c.agent.SetCommonHeader(c.userTokenHeader, "Bearer "+c.userToken)
	return c.agent, nil
}

func (c *HttpUcsClient) getClientAgent(clientAuthKind ClientAuthKind) (*req.Client, error) {
	c.initAgent()
	var clientToken string
	switch clientAuthKind {
	case ClientAuthKindToken:
		if len(c.clientToken) == 0 {
			return nil, errors.New("请提供客户端令牌")
		}
		clientToken = c.clientToken
	case ClientAuthKindIdAndSecret:
		if len(c.clientId) == 0 || len(c.clientSecret) == 0 {
			return nil, errors.New("please provide clientId and clientSecret")
		}
		clientToken = base64.StdEncoding.EncodeToString([]byte(c.clientId + "@" + c.clientSecret))
	default:
		return nil, errors.New("客户端认证方式[" + string(clientAuthKind) + "]错误")
	}

	c.agent.SetCommonHeader(c.clientTokenHeader, "Bearer "+clientToken)
	return c.agent, nil
}

func GenerateRandomKey() string {
	rand.Seed(time.Now().UnixNano())
	output := ""
	for i := 0; i < 6; i++ {
		output += strconv.Itoa(rand.Intn(10))
	}
	return output
}

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

type HttpClient struct {
	baseUrl           string
	timeout           int
	userToken         string
	clientId          string
	clientSecret      string
	accessCode        string
	accessCodeHeader  string
	randomKeyHeader   string
	userTokenHeader   string
	clientTokenHeader string
	agent             *req.Client
}

func (c *HttpClient) SetTimeout(timeout int) Client {
	if timeout > 0 {
		c.timeout = timeout
	}
	return c
}

func (c *HttpClient) SetUserToken(userToken string) Client {
	c.userToken = userToken
	return c
}

func (c *HttpClient) SetClientIdAndSecret(clientId, clientSecret string) Client {
	c.clientId = clientId
	c.clientSecret = clientSecret
	return c
}

func (c *HttpClient) SetHttpHeaderNames(accessCodeHeader, randomKeyHeader, userTokenHeader, clientTokenHeader string) Client {
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

func (c *HttpClient) UserValidateJwt() (*JwtUser, error) {
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
		return nil, fmt.Errorf("error: %v", res)
	}
	if result.Code == 0 {
		return &JwtUser{RawJwtUser: result.Result, Token: c.userToken}, nil
	}
	return nil, errors.New(result.Message)
}

func (c *HttpClient) UserValidatePermOperationByCode(operationCode string) error {
	return c.permitPost(ValidatePermOperationByCodeURL, map[string]string{
		"code": operationCode,
	})
}

func (c *HttpClient) UserValidatePermAction(service, path, method string) error {
	return c.permitPost(ValidatePermActionURL, map[string]string{
		"service": service,
		"path":    path,
		"method":  method,
	})
}

func (c *HttpClient) UserValidatePermOrgById(orgId string) error {
	return c.permitPost(ValidatePermOrgByIdURL, map[string]string{
		"id": orgId,
	})
}

func (c *HttpClient) UserValidatePermActionWithOrgId(service, path, method, orgId string) error {
	return c.permitPost(ValidatePermActionWithOrgIdURL, map[string]string{
		"service": service,
		"path":    path,
		"method":  method,
		"orgId":   orgId,
	})
}

func (c *HttpClient) UserQueryOrgIdsByAction(service, path, method string) (*ActionOrgIds, error) {
	a, err := c.getUserAgent()
	if err != nil {
		return nil, err
	}
	result := &QueryActionOrgIdsHttpResponse{}
	res, err := a.R().SetResult(result).
		SetFormData(map[string]string{
			"service": service,
			"path":    path,
			"method":  method,
		}).
		Post(QueryOrgIdsByActionURL)
	if err != nil {
		return nil, err
	}
	if !res.IsSuccess() {
		return nil, fmt.Errorf("error: %v", res)
	}
	return &ActionOrgIds{
		orgPermissionType: OrgPermissionType(result.Result.OrgPermissionType),
		orgIds:            result.Result.OrgIds,
	}, nil
}

func (c *HttpClient) ClientRequest(method, url string, data map[string]string) (interface{}, error) {
	method = strings.ToUpper(method)
	switch method {
	case http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete, http.MethodPatch:
	default:
		return nil, fmt.Errorf("unsupport method: %v", method)
	}
	a, err := c.getClientAgent()
	if err != nil {
		return nil, err
	}
	result := &NormalHttpResponse{}
	request := a.R().SetResult(result)
	if data != nil {
		request = request.SetFormData(data)
	}
	res, err := request.Send(method, url)
	if err != nil {
		return nil, err
	}
	if !res.IsSuccess() {
		return nil, fmt.Errorf("error: %v", res)
	}
	if result.Code != 0 {
		return nil, fmt.Errorf(result.Message)
	}
	return result.Result, nil
}

func (c *HttpClient) permitPost(url string, data map[string]string) error {
	a, err := c.getUserAgent()
	if err != nil {
		return err
	}
	result := &PermitHttpResponse{}
	res, err := a.R().SetResult(result).SetFormData(data).Post(url)
	if err != nil {
		return err
	}
	if !res.IsSuccess() {
		return fmt.Errorf("error: %v", res)
	}
	if result.Result.Permit {
		return nil
	}
	if len(result.Message) > 0 {
		return errors.New(result.Message)
	}
	return errors.New(DefaultNoPermMsg)
}

func (c *HttpClient) initAgent() {
	if c.agent == nil {
		c.agent = req.C().
			DisableAutoDecode().
			SetAutoDecodeContentType("application/json").
			SetBaseURL(c.baseUrl)
	}
	c.agent.
		SetTimeout(time.Duration(c.timeout)*time.Second).
		SetCommonHeader(c.accessCodeHeader, c.accessCode).
		SetCommonHeader(c.randomKeyHeader, getRandomNumberString(6))
}

func (c *HttpClient) getUserAgent() (*req.Client, error) {
	c.initAgent()
	if len(c.userToken) == 0 {
		return nil, errors.New("please provide userToken")
	}
	c.agent.SetCommonHeader(c.userTokenHeader, "Bearer "+c.userToken)
	return c.agent, nil
}

func (c *HttpClient) getClientAgent() (*req.Client, error) {
	c.initAgent()
	if len(c.clientId) == 0 || len(c.clientSecret) == 0 {
		return nil, errors.New("please provide clientId and clientSecret")
	}
	c.agent.SetCommonHeader(c.clientTokenHeader, base64.StdEncoding.EncodeToString([]byte(c.clientId+"@"+c.clientSecret)))
	return c.agent, nil
}

func getRandomNumberString(length int) string {
	rand.Seed(time.Now().UnixNano())
	output := ""
	for i := 0; i < length; i++ {
		output += strconv.Itoa(rand.Intn(10))
	}
	return output
}

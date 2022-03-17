package ucs

import (
	"errors"
	"fmt"
	"github.com/imroc/req/v3"
	"math/rand"
	"strconv"
	"time"
)

type HttpClient struct {
	addr             string
	port             int
	ssl              bool
	timeout          int
	token            string
	accessCodeHeader string
	accessCode       string
	randomKeyHeader  string
	agent            *req.Client
}

type NormalHttpResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Result  interface{} `json:"result"`
}

type CommonPermitResult struct {
	Permit bool `json:"permit"`
}

type PermitHttpResponse struct {
	Code    int                `json:"code"`
	Message string             `json:"message"`
	Result  CommonPermitResult `json:"result"`
}

func (c *HttpClient) SetTimeout(timeout int) Client {
	if timeout > 0 {
		c.timeout = timeout
	}
	return c
}

func (c *HttpClient) SetToken(token string) Client {
	c.token = token
	return c
}

func (c *HttpClient) SetHttpHeaderNames(accessCodeHeader, randomKeyHeader string) Client {
	c.accessCodeHeader = accessCodeHeader
	c.randomKeyHeader = randomKeyHeader
	return c
}

func (c *HttpClient) ValidateJwt() error {
	a, err := c.getAgent()
	if err != nil {
		return err
	}
	result := &NormalHttpResponse{}
	res, err := a.R().SetResult(result).
		Get(ValidateJwtURL)
	if err != nil {
		return err
	}
	if !res.IsSuccess() {
		return fmt.Errorf("error: %v", res)
	}
	if result.Code == 0 {
		return nil
	}
	return errors.New(result.Message)
}

func (c *HttpClient) ValidatePermOperationByCode(operationCode string) error {
	return c.permitPost(ValidatePermOperationByCodeURL, map[string]string{
		"code": operationCode,
	})
}

func (c *HttpClient) ValidatePermAction(service, path, method string) error {
	return c.permitPost(ValidatePermActionURL, map[string]string{
		"service": service,
		"path":    path,
		"method":  method,
	})
}

func (c *HttpClient) ValidatePermOrgById(orgId string) error {
	return c.permitPost(ValidatePermOrgByIdURL, map[string]string{
		"id": orgId,
	})
}

func (c *HttpClient) permitPost(url string, data map[string]string) error {
	a, err := c.getAgent()
	if err != nil {
		return err
	}
	result := &PermitHttpResponse{}
	res, err := a.R().SetResult(result).
		SetFormData(data).
		Post(url)
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

func (c *HttpClient) getAgent() (*req.Client, error) {
	if len(c.token) == 0 {
		return nil, errors.New("please provide token")
	}
	if c.agent == nil {
		ssl := ""
		if c.ssl {
			ssl = "s"
		}
		c.agent = req.C().
			DisableAutoDecode().
			SetAutoDecodeContentType("application/json").
			SetBaseURL(fmt.Sprintf("http%v://%v:%d", ssl, c.addr, c.port))
	}
	c.agent.
		SetTimeout(time.Duration(c.timeout)*time.Second).
		SetCommonBearerAuthToken(c.token).
		SetCommonHeader(c.accessCodeHeader, c.accessCode).
		SetCommonHeader(c.randomKeyHeader, getRandomNumberString(6))
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

package ucs

import (
	"context"
	"fmt"
	"github.com/Macrow/ucs-go-sdk/pkg/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"time"
)

const (
	DefaultTimeout = 3
)

type Client struct {
	service string
	addr    string
	port    int
	timeout int

	options []grpc.DialOption
	md      metadata.MD
}

func (c *Client) SetTimeout(timeout int) {
	c.timeout = timeout
}

func (c *Client) SetToken(token string) {
	c.md = metadata.Pairs("authorization", token)
}

func (c *Client) ValidateJwt() (bool, error) {
	return authentication(c)
}

func (c *Client) CheckOperationByCode(operationCode string) (bool, error) {
	return authorization(c, &pb.AuthorizationRequest{
		Payload: &pb.AuthorizationRequest_OperationCode{
			OperationCode: operationCode,
		},
	})
}

func (c *Client) CheckAction(service, path, method string) (bool, error) {
	return authorization(c, &pb.AuthorizationRequest{
		Payload: &pb.AuthorizationRequest_Action{
			Action: &pb.Action{
				Service: service,
				Path:    path,
				Method:  method,
			},
		},
	})
}

func (c *Client) CheckOrgById(orgId string) (bool, error) {
	return authorization(c, &pb.AuthorizationRequest{
		Payload: &pb.AuthorizationRequest_OrgId{
			OrgId: orgId,
		},
	})
}

func authentication(c *Client) (bool, error) {
	if c.options == nil {
		return false, fmt.Errorf("please create client instance and set token first")
	}

	conn, err := grpc.Dial(fmt.Sprintf("%v:%d", c.addr, c.port), c.options...)
	defer func() {
		if conn != nil {
			_ = conn.Close()
		}
	}()
	if err != nil {
		return false, fmt.Errorf(err.Error())
	}

	service := pb.NewAuthServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(c.timeout))
	defer cancel()
	ctx = metadata.NewOutgoingContext(ctx, c.md)

	res, err := service.Authentication(ctx, &pb.AuthenticationRequest{})
	if err != nil {
		return false, fmt.Errorf(err.Error())
	}
	if !res.Success {
		return false, fmt.Errorf(res.Error.Reason)
	}
	return true, nil
}

func authorization(c *Client, req *pb.AuthorizationRequest) (bool, error) {
	if c.options == nil {
		return false, fmt.Errorf("please create client instance first")
	}

	conn, err := grpc.Dial(fmt.Sprintf("%v:%d", c.addr, c.port), c.options...)
	defer func() {
		if conn != nil {
			_ = conn.Close()
		}
	}()

	if err != nil {
		return false, fmt.Errorf(err.Error())
	}
	service := pb.NewAuthServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(c.timeout))
	defer cancel()
	ctx = metadata.NewOutgoingContext(ctx, c.md)

	res, err := service.Authorization(ctx, req)
	if err != nil {
		return false, fmt.Errorf(err.Error())
	}
	if !res.Success {
		return false, fmt.Errorf(res.Error.Reason)
	}
	return true, nil
}

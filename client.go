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

func (c *Client) ValidateJwt() error {
	return authentication(c)
}

func (c *Client) ValidatePermOperationByCode(operationCode string) error {
	return authorization(c, &pb.AuthorizationRequest{
		Payload: &pb.AuthorizationRequest_OperationCode{
			OperationCode: operationCode,
		},
	})
}

func (c *Client) ValidatePermAction(service, path, method string) error {
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

func (c *Client) ValidatePermOrgById(orgId string) error {
	return authorization(c, &pb.AuthorizationRequest{
		Payload: &pb.AuthorizationRequest_OrgId{
			OrgId: orgId,
		},
	})
}

func authentication(c *Client) error {
	if c.options == nil {
		return fmt.Errorf("please create client instance and set token first")
	}

	conn, err := grpc.Dial(fmt.Sprintf("%v:%d", c.addr, c.port), c.options...)
	defer func() {
		if conn != nil {
			_ = conn.Close()
		}
	}()
	if err != nil {
		return fmt.Errorf(err.Error())
	}

	service := pb.NewAuthServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(c.timeout))
	defer cancel()
	ctx = metadata.NewOutgoingContext(ctx, c.md)

	res, err := service.Authentication(ctx, &pb.AuthenticationRequest{})
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	if !res.Success {
		return fmt.Errorf(res.Error.Reason)
	}
	return nil
}

func authorization(c *Client, req *pb.AuthorizationRequest) error {
	if c.options == nil {
		return fmt.Errorf("please create client instance first")
	}

	conn, err := grpc.Dial(fmt.Sprintf("%v:%d", c.addr, c.port), c.options...)
	defer func() {
		if conn != nil {
			_ = conn.Close()
		}
	}()

	if err != nil {
		return fmt.Errorf(err.Error())
	}
	service := pb.NewAuthServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(c.timeout))
	defer cancel()
	ctx = metadata.NewOutgoingContext(ctx, c.md)

	res, err := service.Authorization(ctx, req)
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	if !res.Success {
		return fmt.Errorf(res.Error.Reason)
	}
	return nil
}

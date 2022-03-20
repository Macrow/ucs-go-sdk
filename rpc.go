package ucs

import (
	"context"
	"fmt"
	"github.com/Macrow/ucs-go-sdk/pkg/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"time"
)

type RpcClient struct {
	addr    string
	timeout int

	options []grpc.DialOption
	md      metadata.MD
}

func (c *RpcClient) SetTimeout(timeout int) Client {
	if timeout > 0 {
		c.timeout = timeout
	}
	return c
}

func (c *RpcClient) SetToken(token string) Client {
	c.md = metadata.Pairs("authorization", token)
	return c
}

func (c *RpcClient) SetHttpHeaderNames(accessCodeHeader, randomKeyHeader string) Client {
	return c
}

func (c *RpcClient) ValidateJwt() error {
	return authentication(c)
}

func (c *RpcClient) ValidatePermOperationByCode(operationCode string) error {
	return authorization(c, &pb.AuthorizationRequest{
		Payload: &pb.AuthorizationRequest_OperationCode{
			OperationCode: operationCode,
		},
	})
}

func (c *RpcClient) ValidatePermAction(service, path, method string) error {
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

func (c *RpcClient) ValidatePermOrgById(orgId string) error {
	return authorization(c, &pb.AuthorizationRequest{
		Payload: &pb.AuthorizationRequest_OrgId{
			OrgId: orgId,
		},
	})
}

func (c *RpcClient) RenewToken() (string, error) {
	if c.options == nil {
		return "", fmt.Errorf("please create client instance and set token first")
	}

	conn, err := grpc.Dial(c.addr, c.options...)
	defer func() {
		if conn != nil {
			_ = conn.Close()
		}
	}()
	if err != nil {
		return "", fmt.Errorf(err.Error())
	}

	service := pb.NewAuthServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(c.timeout))
	defer cancel()
	ctx = metadata.NewOutgoingContext(ctx, c.md)

	res, err := service.RenewToken(ctx, &pb.RenewTokenRequest{})
	if err != nil {
		return "", fmt.Errorf(err.Error())
	}
	if !res.Success {
		return "", fmt.Errorf(res.Message)
	}
	return res.Token, nil
}

func authentication(c *RpcClient) error {
	if c.options == nil {
		return fmt.Errorf("please create client instance and set token first")
	}

	conn, err := grpc.Dial(c.addr, c.options...)
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
		return fmt.Errorf(res.Message)
	}
	return nil
}

func authorization(c *RpcClient, req *pb.AuthorizationRequest) error {
	if c.options == nil {
		return fmt.Errorf("please create client instance first")
	}

	conn, err := grpc.Dial(c.addr, c.options...)
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
		return fmt.Errorf(res.Message)
	}
	return nil
}

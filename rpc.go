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

func (c *RpcClient) ValidatePermActionWithOrgId(service, path, method, orgId string) error {
	return authorization(c, &pb.AuthorizationRequest{
		Payload: &pb.AuthorizationRequest_ActionWithOrgId{
			ActionWithOrgId: &pb.ActionWithOrgId{
				Action: &pb.Action{
					Service: service,
					Path:    path,
					Method:  method,
				},
				OrgId: orgId,
			},
		},
	})
}

func (c *RpcClient) QueryOrgIdsByAction(service, path, method string) (*ActionOrgIds, error) {
	if c.options == nil {
		return nil, fmt.Errorf("please create client instance first")
	}

	conn, err := grpc.Dial(c.addr, c.options...)
	defer func() {
		if conn != nil {
			_ = conn.Close()
		}
	}()

	if err != nil {
		return nil, fmt.Errorf(err.Error())
	}
	client := pb.NewAuthServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(c.timeout))
	defer cancel()
	ctx = metadata.NewOutgoingContext(ctx, c.md)

	res, err := client.Authorization(ctx, &pb.AuthorizationRequest{
		Payload: &pb.AuthorizationRequest_OrgIdsByAction{
			OrgIdsByAction: &pb.Action{
				Service: service,
				Path:    path,
				Method:  method,
			},
		},
	})
	if err != nil {
		return nil, fmt.Errorf(err.Error())
	}
	return &ActionOrgIds{
		orgPermissionType: OrgPermissionType(res.GetOrgIds().GetOrgPermissionType()),
		orgIds:            res.GetOrgIds().GetOrgIds(),
	}, nil
}

func (c *RpcClient) OAuth2TokenByAuthorizationCode(code, clientId, clientSecret, deviceId, deviceName string) (string, error) {
	return c.oAuth2Token(&pb.OAuth2TokenRequest{
		Payload: &pb.OAuth2TokenRequest_AuthorizationCode{
			AuthorizationCode: &pb.AuthorizationCode{
				Code:         code,
				ClientId:     clientId,
				ClientSecret: clientSecret,
				DeviceId:     deviceId,
				DeviceName:   deviceName,
			},
		},
	})
}

func (c *RpcClient) OAuth2TokenByPassword(username, password, deviceId, deviceName string) (string, error) {
	return c.oAuth2Token(&pb.OAuth2TokenRequest{
		Payload: &pb.OAuth2TokenRequest_PasswordCredentials{
			PasswordCredentials: &pb.PasswordCredentials{
				Username:   username,
				Password:   password,
				DeviceId:   deviceId,
				DeviceName: deviceName,
			},
		},
	})
}

func (c *RpcClient) oAuth2Token(request *pb.OAuth2TokenRequest) (string, error) {
	if c.options == nil {
		return "", fmt.Errorf("please create client instance first")
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
	client := pb.NewAuthServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(c.timeout))
	defer cancel()
	ctx = metadata.NewOutgoingContext(ctx, c.md)

	res, err := client.OAuth2Token(ctx, request)
	if err != nil {
		return "", fmt.Errorf(err.Error())
	}
	return res.AccessToken, nil
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

	client := pb.NewAuthServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(c.timeout))
	defer cancel()
	ctx = metadata.NewOutgoingContext(ctx, c.md)

	res, err := client.Authentication(ctx, &pb.AuthenticationRequest{})
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
	client := pb.NewAuthServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(c.timeout))
	defer cancel()
	ctx = metadata.NewOutgoingContext(ctx, c.md)

	res, err := client.Authorization(ctx, req)
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	if !res.Success {
		return fmt.Errorf(res.Message)
	}
	return nil
}

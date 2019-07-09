package client

import (
	"context"
	gateway "github.com/noguchi-hiroshi/grpc-gateway-test/proto"
	"github.com/noguchi-hiroshi/grpc-gateway-test/user"
)

type gRPCUserClient struct {
	sc gateway.UserServiceClient
}

func NewGRPCUserClient(sc gateway.UserServiceClient) UserClient {
	return &gRPCUserClient{sc: sc}
}

func (c *gRPCUserClient) Find(id int64) (*user.Entity, error) {
	ctx := context.Background()
	res, err := c.sc.Find(ctx, &gateway.FindRequest{
		ID: id,
	})
	if err != nil {
		return nil, err
	}
	return &user.Entity{
		ID: res.ID,
		Email: res.Email,
		Password: res.Password,
	}, nil
}

func (c *gRPCUserClient) Create(email string, password string) (int64, error) {
	ctx := context.Background()
	res, err := c.sc.Create(ctx, &gateway.CreateRequest{
		Email: email,
		Password: password,
	})
	if err != nil {
		return 0, err
	}
	return res.ID, nil
}

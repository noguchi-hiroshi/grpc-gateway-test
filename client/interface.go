package client

import "github.com/noguchi-hiroshi/grpc-gateway-test/user"

type UserClient interface {
	Create(email string, password string) (int64, error)
	Find(id int64) (*user.Entity, error)
}

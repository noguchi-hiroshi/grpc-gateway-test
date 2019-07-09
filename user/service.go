package user

import (
	"context"
	gateway "github.com/noguchi-hiroshi/grpc-gateway-test/proto"
)

type service struct {
	model Model
}

func NewService(model Model) gateway.UserServiceServer {
	return &service{model: model}
}

func (s *service) Create(_ context.Context, req *gateway.CreateRequest) (*gateway.CreateResponse, error) {
	user, err := s.model.Create(req.Email, req.Password)
	if err != nil {
		return nil, err
	}
	return &gateway.CreateResponse{ID: user.ID}, nil
}

func (s *service) Find(_ context.Context, req *gateway.FindRequest) (*gateway.FindResponse, error) {
	user, err := s.model.Find(req.ID)
	if err != nil {
		return nil, err
	}
	return &gateway.FindResponse{
		ID: user.ID,
		Email: user.Email,
		Password: user.Password,
	}, nil
}

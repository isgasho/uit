package service

import (
	"context"

	"github.com/uit/example/proto_go/todo"
)

func (s *Service) List(ctx context.Context, in *todo.ListReq) (*todo.ListRsp, error) {
	return &todo.ListRsp{}, nil
}

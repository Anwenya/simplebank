package gapi

import (
	"fmt"

	db "com.wlq/simplebank/db/sqlc"
	"com.wlq/simplebank/pb"
	"com.wlq/simplebank/token"
	"com.wlq/simplebank/util"
	"com.wlq/simplebank/worker"
)

// gRPC服务
type Server struct {
	pb.UnimplementedSimpleBankServer
	config          util.Config
	store           db.Store
	tokenMaker      token.Maker
	taskDistributor worker.TaskDistributor
}

// 创建gRPC服务
func NewServer(
	config util.Config,
	store db.Store,
	taskDistributor worker.TaskDistributor,
) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker:%w", err)
	}

	server := &Server{
		config:          config,
		store:           store,
		tokenMaker:      tokenMaker,
		taskDistributor: taskDistributor,
	}

	return server, nil
}

package gapi

import (
	"fmt"

	db "github.com/sbbullet/simple-bank/db/sqlc"
	pb "github.com/sbbullet/simple-bank/pb"
	"github.com/sbbullet/simple-bank/token"
	"github.com/sbbullet/simple-bank/util"
)

type Server struct {
	pb.UnimplementedSimpleBankServer
	config     util.Config
	store      db.Store
	tokenMaker token.Maker
}

func NewServer(config util.Config, store db.Store) (*Server, error) {
	maker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create paseto token maker. Error: %w", err)
	}
	server := &Server{
		config:     config,
		store:      store,
		tokenMaker: maker,
	}

	return server, nil
}

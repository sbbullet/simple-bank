package gapi

import (
	"context"
	"log"

	"github.com/lib/pq"
	db "github.com/sbbullet/simple-bank/db/sqlc"
	pb "github.com/sbbullet/simple-bank/pb"
	"github.com/sbbullet/simple-bank/util"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	hashedPassword, err := util.HashPassword(req.GetPassword())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "cannot create hashed password")
	}

	arg := db.CreateUserParams{
		Username:       req.Username,
		HashedPassword: hashedPassword,
		FullName:       req.FullName,
		Email:          req.Email,
	}

	user, err := server.store.CreateUser(ctx, arg)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			log.Println(pqErr.Code.Name())
			switch pqErr.Code.Name() {
			case "unique_violation":
				return nil, status.Errorf(codes.AlreadyExists, "user already exists")
			}
		}

		return nil, status.Errorf(codes.Internal, "cannot create user")
	}

	res := &pb.CreateUserResponse{
		User: convertUser(user),
	}

	return res, nil
}

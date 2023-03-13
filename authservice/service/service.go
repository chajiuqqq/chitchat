package service

import (
	"context"

	"github.com/chajiuqqq/chitchat/common/pb"

	. "github.com/chajiuqqq/chitchat/common/data"
	"github.com/chajiuqqq/chitchat/common/entity"
	"github.com/chajiuqqq/chitchat/common/util"
)

type AuthService struct {
	pb.UnimplementedAuthServiceServer
}

func (au *AuthService) Check(ctx context.Context, req *pb.CheckRequest) (*pb.CheckResponse, error) {
	sess := &entity.Session{}
	result := Db.Where("uuid=? and deleted_at is null", req.Uuid).Find(sess)
	err := result.Error
	if err == nil && result.RowsAffected == 1 {
		return &pb.CheckResponse{
			Exist: true,
			Sess: &pb.AuthSession{
				Uuid:   sess.Uuid,
				Email:  sess.Email,
				UserId: uint32(sess.UserId),
			},
		}, nil
	}
	return &pb.CheckResponse{Exist: false}, err
}

func (au *AuthService) GetUserByEmail(ctx context.Context, req *pb.GetUserByEmailRequest) (*pb.GetUserByEmailResponse, error) {
	email := req.Email
	u := entity.User{}
	err := Db.First(&u, "email=?", email).Error
	if err != nil {
		return &pb.GetUserByEmailResponse{}, err
	}
	return &pb.GetUserByEmailResponse{
		Exist: true,
		User: &pb.AuthUser{
			ID:       uint32(u.ID),
			Uuid:     u.Uuid,
			Name:     u.Name,
			Email:    u.Email,
			Password: u.Password,
		},
	}, err
}
func (au *AuthService) Encrypt(ctx context.Context, req *pb.EncryptRequest) (*pb.EncryptResponse, error) {
	return &pb.EncryptResponse{
		Out: req.Src,
	}, nil
}

func (au *AuthService) NewSession(ctx context.Context, req *pb.NewSessionReq) (*pb.AuthSession, error) {
	Db.Delete(&entity.Session{}, "user_id=?", req.UserId)

	sess := entity.Session{
		Uuid:   util.GenerateUuid(),
		UserId: uint(req.UserId),
		Email:  req.Email,
	}
	err := Db.Create(&sess).Error
	return &pb.AuthSession{
		Uuid:   sess.Uuid,
		UserId: uint32(sess.UserId),
		Email:  sess.Email,
	}, err
}

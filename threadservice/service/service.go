package service

import (
	"context"
	"log"

	. "github.com/chajiuqqq/chitchat/common/data"
	"github.com/chajiuqqq/chitchat/common/entity"
	"github.com/chajiuqqq/chitchat/common/pb"
	"github.com/chajiuqqq/chitchat/common/util"
)

type ThreadService struct {
	pb.UnimplementedThreadServiceServer
}

func (ts *ThreadService) Get(ctx context.Context, req *pb.GetThreadRequest) (*pb.GetThreadResponse, error) {
	tid := req.ThreadId
	thread := &entity.Thread{}
	err := Db.Preload("Posts").Preload("Posts.User").Preload("User").First(&thread, tid).Error
	if err != nil {
		log.Println("fail to find thread,", err)
	}
	return util.ToThreadRPC(thread), err
}
func (ts *ThreadService) GetAll(ctx context.Context, em *pb.Empty) (*pb.GetAllResponse, error) {
	threads := make([]entity.Thread, 0)
	err := Db.Preload("Posts").Preload("Posts.User").Preload("User").Find(&threads).Error
	rpcTheads := make([]*pb.GetThreadResponse, 0)
	for _, th := range threads {
		rpcTheads = append(rpcTheads, util.ToThreadRPC(&th))
	}
	return &pb.GetAllResponse{
		Threads: rpcTheads,
	}, err
}
func (ts *ThreadService) AddPost(ctx context.Context, req *pb.AddPostRequest) (*pb.Empty, error) {
	tid := req.ThreadId
	body := req.Body
	userId := req.UserId
	var thread entity.Thread
	Db.First(&thread, tid)
	err := Db.Model(&thread).Association("Posts").Append(
		&entity.Post{Uuid: util.GenerateUuid(), Body: body, UserId: uint(userId)},
	)
	return &pb.Empty{}, err
}

func (ts *ThreadService) Create(ctx context.Context, req *pb.CreateThreadReq) (*pb.Empty, error) {
	thread := &entity.Thread{
		Uuid:   req.Uuid,
		Topic:  req.Topic,
		UserId: uint(req.UserId),
	}
	err := Db.Create(thread).Error
	return &pb.Empty{}, err
}

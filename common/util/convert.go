package util

import (
	"log"
	"time"

	"github.com/chajiuqqq/chitchat/common/entity"
	"github.com/chajiuqqq/chitchat/common/pb"
	"gorm.io/gorm"
)

func ToThreadEntity(th *pb.GetThreadResponse) *entity.Thread {

	posts := make([]*entity.Post, 0)
	for _, item := range th.Posts {
		var user *entity.User
		if item.User != nil {
			user = &entity.User{
				Uuid:  item.User.Uuid,
				Email: item.User.Email,
				Name:  item.User.Name,
			}
		}
		posts = append(posts, &entity.Post{
			Uuid:     item.Uuid,
			Body:     item.Body,
			UserId:   uint(item.UserId),
			ThreadId: uint(item.ThreadId),
			User:     user,
		})
	}

	var user *entity.User
	if th.User != nil {
		user = &entity.User{
			Uuid:  th.User.Uuid,
			Email: th.User.Email,
			Name:  th.User.Name,
		}
	}

	createdAt, err := time.Parse("2016.01.02 15:04:05", th.CreatedAt)
	if err != nil {
		log.Panic(err)
	}

	return &entity.Thread{
		Model: gorm.Model{
			ID:        uint(th.ID),
			CreatedAt: createdAt,
		},
		Uuid:   th.Uuid,
		Topic:  th.Topic,
		UserId: uint(th.UserId),
		Posts:  posts,
		User:   user,
	}
}

func ToThreadRPC(th *entity.Thread) *pb.GetThreadResponse {

	posts := make([]*pb.TheadPost, 0)
	for _, item := range th.Posts {
		var user *pb.TheadUser
		if item.User != nil {
			user = &pb.TheadUser{
				Uuid:  item.User.Uuid,
				Email: item.User.Email,
				Name:  item.User.Name,
			}
		}
		posts = append(posts, &pb.TheadPost{
			Uuid:      item.Uuid,
			Body:      item.Body,
			UserId:    uint32(item.UserId),
			ThreadId:  uint32(item.ThreadId),
			User:      user,
			CreatedAt: timeFormat(&item.CreatedAt),
		})
	}

	var user *pb.TheadUser
	if th.User != nil {
		user = &pb.TheadUser{
			Uuid:  th.User.Uuid,
			Email: th.User.Email,
			Name:  th.User.Name,
		}
	}

	return &pb.GetThreadResponse{
		ID:        uint32(th.ID),
		Uuid:      th.Uuid,
		Topic:     th.Topic,
		UserId:    uint32(th.UserId),
		Posts:     posts,
		User:      user,
		CreatedAt: timeFormat(&th.CreatedAt),
	}
}

func timeFormat(t *time.Time) string {
	return t.Format("2006.01.02 15:04:05")
}

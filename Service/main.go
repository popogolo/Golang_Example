package service

import (
	"../repository"
)

type PageInfo struct {
	Topic    *repository.Topic
	PostList []*repository.Post
}

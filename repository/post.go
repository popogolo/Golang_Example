package repository

import (
	"sync"
)

type Post struct {
	ID         int64  `json:"id"`
	TopicID    int64  `json:"topic_id"`
	Content    string `json:"content"`
	CreateTime string `json:"create_time"`
}
type PostDao struct {
}

var (
	postDao  *PostDao
	postOnce sync.Once
)

func NewPostDaoInstance() *PostDao {
	postOnce.Do(
		func() {
			postDao = &PostDao{}
		})
	return postDao
}

func (*PostDao) QueryPostsByParentId(topicId int64) []*Post {
	return postIndexMap[topicId]
}

package repository

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

var (
	topicIndexMap map[int64]*Topic
	postIndexMap  map[int64][]*Post
)

func initTopicIndexMap(filePath string) error {
	open, err := os.Open(filePath + "topic")
	if err != nil {
		return fmt.Errorf("ERROR IN READ FILE %v", filePath+"topic")
	}
	scanner := bufio.NewScanner(open)
	topicTmpMap := make(map[int64]*Topic)
	for scanner.Scan() {
		text := scanner.Text()
		var topic Topic
		if err := json.Unmarshal([]byte(text), &topic); err != nil {
			return err
		}
		topicTmpMap[topic.ID] = &topic
	}
	topicIndexMap = topicTmpMap
	return nil
}

func initPostIndexMap(filePath string) error {
	open, err := os.Open(filePath + "topic")
	if err != nil {
		return fmt.Errorf("ERROR IN READ FILE %v", filePath+"topic")
	}
	scanner := bufio.NewScanner(open)
	postTmpMap := make(map[int64][]*Post)
	for scanner.Scan() {
		text := scanner.Text()
		var post Post
		if err := json.Unmarshal([]byte(text), &post); err != nil {
			return err
		}
		postList, ok := postTmpMap[post.TopicID]
		if !ok {
			postTmpMap[post.TopicID] = []*Post{&post}
			continue
		}
		postList = append(postList, &post)
		postTmpMap[post.TopicID] = postList
	}
	postIndexMap = postTmpMap
	return nil
}

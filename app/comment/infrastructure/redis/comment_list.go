package redis

import (
	"context"
	"encoding/json"
	"time"
	"wiliwili/app/comment/domain/model"
)

func (c *CommentDBCache) DelCommentCache(ctx context.Context, id string) error {
	return c.client.Del(ctx, "comment_list:"+id).Err()
}
func (c *CommentDBCache) GetCommentList(ctx context.Context, id string, page, pageSize, commenttype int64) ([]*model.Comment, error) {
	key := "comment_list:" + id + ":" + string(commenttype) + ":" + string(page) + ":" + string(pageSize)
	val, err := c.client.Get(ctx, key).Result()
	if err != nil {
		return nil, err
	}
	var models []*model.Comment
	if err := json.Unmarshal([]byte(val), &models); err != nil {
		return nil, err
	}
	return models, nil
}

func (c *CommentDBCache) SetCommentList(ctx context.Context, id string, commenttype int64, models []*model.Comment) error {
	key := "comment_list:" + id + ":" + string(commenttype) + ":" + string(1) + ":" + string(10)
	return c.client.Set(ctx, key, models, 24*time.Hour).Err()
}

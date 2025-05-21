package redis

import (
	"context"
	"strconv"
	"time"
	"wiliwili/app/user/domain/model"
)

func (c *UserDBCache) GetImage(ctx context.Context, imageID int64) (string, error) {
	key := "image_url:" + strconv.FormatInt(imageID, 10)
	return c.client.Get(ctx, key).Result()
}

func (c *UserDBCache) StoreImage(ctx context.Context, image *model.Image) error {
	key := "image_url:" + strconv.FormatInt(image.ImageID, 10)
	return c.client.Set(ctx, key, image.Url, 24*time.Hour).Err()
}

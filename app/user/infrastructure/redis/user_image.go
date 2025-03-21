package redis

import (
	"context"
	"wiliwili/app/user/domain/model"
)

func (c *UserDBCache) GetImage(ctx context.Context, imageid int64) (*model.Image, error) {
	panic("implement me")
}

func (c *UserDBCache) StoreImage(ctx context.Context, image *model.Image) error {
	panic("implement me")
}

package mysql

import (
	"wiliwili/app/like/domain/repository"

	"gorm.io/gorm"
)

type likeDB struct {
	client *gorm.DB
}

func NewLikeDB(client *gorm.DB) repository.LikeDB {
	return &likeDB{client: client}
}

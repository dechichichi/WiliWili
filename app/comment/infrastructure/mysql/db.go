package mysql

import (
	"wiliwili/app/comment/domain/repository"

	"gorm.io/gorm"
)

type commentDB struct {
	client *gorm.DB
}

func NewCommentDB(client *gorm.DB) repository.CommentDB {
	return &commentDB{
		client: client,
	}
}

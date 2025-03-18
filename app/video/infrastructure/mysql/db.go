package mysql

import (
	"wiliwili/app/video/domain/repository"

	"gorm.io/gorm"
)

type videoDB struct {
	client *gorm.DB
}

func NewVideoDB(client *gorm.DB) repository.VideoDB {
	return &videoDB{client: client}
}

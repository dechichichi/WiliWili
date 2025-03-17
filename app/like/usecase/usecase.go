package usecase

import (
	"wiliwili/app/like/domain/repository"
	"wiliwili/app/like/domain/service"
)

type LikeUseCase interface {

}

type UseCase struct {
	db    repository.LikeDB
	svc   *service.LikeService
	cache repository.LikeCache
}

func NewLikeUseCase(db repository.LikeDB, svc *service.LikeService, cache repository.LikeCache) LikeUseCase {
	return &UseCase{
		db:    db,
		svc:   svc,
		cache: cache,
	}
}

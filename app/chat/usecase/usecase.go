package usecase

import "wiliwili/app/chat/domain"

type ChatUsecase struct {
	Client domain.Client
	Hub    domain.Hub
}

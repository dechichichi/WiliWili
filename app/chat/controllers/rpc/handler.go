package rpc

import "wiliwili/app/chat/usecase"

type ChatHandler struct {
	useCase usecase.ChatUsecase
}

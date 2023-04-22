package gateway

import (
	"context"

	"github.com/gfbritto/chatservice/chatservice/internal/domain/entity"
)

type chatGateway interface {
	CreateChat(context context.Context, chat *entity.Chat) error
	FindChatByID(context context.Context, chatID string) (*entity.Chat, error)
	SaveChat(context context.Context, chat *entity.Chat) error
}

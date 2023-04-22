package chatcompletionstream

import (
	"github.com/gfbritto/chatservice/chatservice/internal/domain/gateway"
	openai "github.com/sashabaranov/go-openai"
)

type ChatCompletionUseCase struct {
	chatGateway  gateway.
	OpenAiClient *openai.Client
}

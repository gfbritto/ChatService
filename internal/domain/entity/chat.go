package entity

import (
	"errors"

	"github.com/google/uuid"
)

type ChatConfig struct {
	Model            *Model
	Temperature      float32
	TopP             float32
	N                int
	Stop             []string
	MaxTokens        int
	PresencePenatty  float32
	FrequencyPenatty float32
}

type Chat struct {
	ID                   string
	UserID               string
	InitialSystemMessage *Message
	Messages             []*Message
	ErasedMessages       []*Message
	Status               string
	TokenUsage           int
	Config               *ChatConfig
}

const (
	ChatStatusActive = "active"
	ChatStatusEnded  = "ended"
)

func NewChat(userID string, initialSystemMessage *Message, chatconfig *ChatConfig) (*Chat, error) {
	chat := &Chat{
		ID:                   uuid.New().String(),
		UserID:               userID,
		InitialSystemMessage: initialSystemMessage,
		Status:               ChatStatusActive,
		Config:               chatconfig,
		TokenUsage:           0,
	}
	chat.AddMessage(initialSystemMessage)

	if err := chat.Validate(); err != nil {
		return nil, err
	}
	return chat, nil
}

func (c *Chat) Validate() error {
	if c.UserID == "" {
		return errors.New("userId is empty")
	}
	if c.Status != ChatStatusActive && c.Status != ChatStatusEnded {
		return errors.New("invalid status")
	}
	if c.Config.Temperature < 0 || c.Config.Temperature > 1 {
		return errors.New("invalid temperature")
	}
	return nil
}

func (c *Chat) AddMessage(m *Message) error {
	if c.Status == ChatStatusEnded {
		return errors.New("Chat has ended. No more messages allowed")
	}
	for {
		if c.Config.Model.GetMaxTokens() >= m.GetTokensQuantity()+c.TokenUsage {
			c.ErasedMessages = append(c.ErasedMessages, c.Messages[0])
			c.Messages = c.Messages[1:]
			c.RefreshTokenUsage()
			break
		}
		c.ErasedMessages = append(c.ErasedMessages, c.Messages[0])
		c.Messages = c.Messages[1:]
		c.RefreshTokenUsage()
	}
	return nil
}

func (c *Chat) GetMessages() []*Message {
	return c.Messages
}

func (c *Chat) CountMessages() int {
	return len(c.Messages)
}

func (c *Chat) End() {
	c.Status = ChatStatusEnded
}

func (c *Chat) RefreshTokenUsage() {
	c.TokenUsage = 0
	for m := range c.Messages {
		c.TokenUsage += c.Messages[m].GetTokensQuantity()
	}
}

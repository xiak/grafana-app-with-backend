package biz

import (
	"context"

	l "github.com/grafana/grafana-plugin-sdk-go/backend/log"
)

type ChatRepo interface {
	GetChatPromptsFromText(ctx context.Context, text string) ([]*Propmpt, error)
}

type ChatUsercase struct {
	repo ChatRepo
	log  l.Logger
}

type Propmpt struct {
	Text string
}

func NewChatUsercase(repo ChatRepo, logger l.Logger) *ChatUsercase {
	return &ChatUsercase{
		repo: repo,
		log:  logger,
	}
}

func (uc *ChatUsercase) GetChatPromptsFromText(ctx context.Context, text string) ([]*Propmpt, error) {
	return uc.repo.GetChatPromptsFromText(ctx, text)
}

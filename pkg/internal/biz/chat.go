package biz

import (
	"context"

	l "github.com/grafana/grafana-plugin-sdk-go/backend/log"
)

type ChatRepo interface {
	GetChatPromptsFromText(ctx context.Context, text string) ([]*Propmpt, error)
}

type ChatUsecase struct {
	repo ChatRepo
	log  l.Logger
}

type Propmpt struct {
	Text string
}

func NewChatUsecase(repo ChatRepo, logger l.Logger) *ChatUsecase {
	return &ChatUsecase{
		repo: repo,
		log:  logger,
	}
}

func (uc *ChatUsecase) GetChatPromptsFromText(ctx context.Context, text string) ([]*Propmpt, error) {
	return uc.repo.GetChatPromptsFromText(ctx, text)
}

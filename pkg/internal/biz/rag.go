package biz

import (
	"context"

	l "github.com/grafana/grafana-plugin-sdk-go/backend/log"
)

type RagRepo interface {
	// get rag keyword list
	GetRagKeywords(ctx context.Context, keyword string) ([]string, error)
}

type RagUsecase struct {
	repo RagRepo
	log  l.Logger
}

func NewSystemSecurityUsecase(repo RagRepo, logger l.Logger) *RagUsecase {
	return &RagUsecase{
		repo: repo,
		log:  logger,
	}
}

func (uc *RagUsecase) GetRagKeywords(ctx context.Context, keyword string) ([]string, error) {
	return uc.repo.GetRagKeywords(ctx, keyword)
}

package data

import (
	"context"

	"github.com/xiak/grafana-app-with-backend/pkg/internal/biz"

	l "github.com/grafana/grafana-plugin-sdk-go/backend/log"
)

var _ biz.RagRepo = (*ragRepo)(nil)

type ragRepo struct {
	data *Data
	log  l.Logger
}

func NewRagRepo(data *Data, logger l.Logger) biz.RagRepo {
	return &ragRepo{
		data: data,
		log:  logger,
	}
}

func (repo *ragRepo) GetRagKeywords(ctx context.Context, keyword string) ([]string, error) {
	if len(repo.data.rag.Prompts.Keywords) <= 0 {
		return nil, nil
	}
	return repo.data.rag.Prompts.Keywords, nil
}

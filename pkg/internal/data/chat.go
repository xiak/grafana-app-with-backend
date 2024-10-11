package data

import (
	"context"

	l "github.com/grafana/grafana-plugin-sdk-go/backend/log"
	"github.com/xiak/grafana-app-with-backend/pkg/internal/biz"
	"github.com/xiak/grafana-app-with-backend/pkg/internal/data/keywords"
)

var _ biz.ChatRepo = (*chatRepo)(nil)

// KeywordListOption KeywordList options
type PromptListOption func(o *pOtion)

// options put here
type pOtion struct {
	// keywords retrieve number
	retrieve int64
}

// Retrieve keywords retrieve numbers
func retrieve(num int64) PromptListOption {
	return func(k *pOtion) {
		k.retrieve = num
	}
}

func newPrompts(opts ...PromptListOption) keywords.Keywords {
	op := pOtion{
		retrieve: 5,
	}
	for _, o := range opts {
		o(&op)
	}
	return keywords.NewKeywordList(
		keywords.Retrieve(op.retrieve),
	)
}

type chatRepo struct {
	data    *Data
	prompts keywords.Keywords
	log     l.Logger
}

func NewChatRepo(data *Data, logger l.Logger) biz.ChatRepo {
	repo := chatRepo{
		data: data,
		log:  logger,
		prompts: newPrompts(
			retrieve(data.chat.Prompts.Retrieve),
		),
	}
	repo.prompts.AddRaw(repo.data.chat.Prompts.Keywords)
	return &repo
}

func (r *chatRepo) GetChatPromptsFromText(ctx context.Context, text string) ([]*biz.Propmpt, error) {
	ps := make([]*biz.Propmpt, r.prompts.Length())
	for _, p := range r.prompts.Find(text) {
		ps = append(ps, &biz.Propmpt{Text: p})
	}
	return ps, nil
}

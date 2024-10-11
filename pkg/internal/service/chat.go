package service

import (
	"context"

	v1 "github.com/xiak/grafana-app-with-backend/pkg/api/copilot/v1"
)

func (s *CopilotService) PromptSuggestion(ctx context.Context, req *v1.PromptSuggestionRequest) (*v1.PromptSuggestionReply, error) {
	prompts, err := s.cuc.GetChatPromptsFromText(ctx, req.Text)
	if err != nil {
		return nil, err
	}
	num := len(prompts)
	p := make([]*v1.Suggestion, 0)
	for _, prompt := range prompts {
		p = append(p, &v1.Suggestion{
			Prompt: prompt.Text,
		})
	}
	return &v1.PromptSuggestionReply{
		Num:     int64(num),
		Prompts: p,
	}, nil
}

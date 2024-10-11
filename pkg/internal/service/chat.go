package service

import (
	"context"

	v1 "github.com/xiak/grafana-app-with-backend/pkg/api/copilot/v1"
)

func (s *CopilotService) PromptSuggestion(ctx context.Context, req *v1.PromptSuggestionRequest) (*v1.PromptSuggestionReply, error) {
	return &v1.PromptSuggestionReply{}, nil
}

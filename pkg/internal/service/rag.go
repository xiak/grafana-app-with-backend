package service

import (
	"context"

	v1 "github.com/xiak/grafana-app-with-backend/pkg/api/copilot/v1"
)

func (s *CopilotService) GetRagKeywords(ctx context.Context, req *v1.GetRagKeywordsRequest) (*v1.GetRagKeywordsReply, error) {
	keywords, err := s.ruc.GetRagKeywords(ctx, req.Text)
	if err != nil {
		return nil, err
	}
	return &v1.GetRagKeywordsReply{
		Keywords: keywords,
	}, nil
}

package service

import (
	"github.com/xiak/grafana-app-with-backend/pkg/internal/biz"

	"github.com/google/wire"
	l "github.com/grafana/grafana-plugin-sdk-go/backend/log"
)

var ProviderSet = wire.NewSet(NewCopilotService)

type CopilotService struct {
	ouc *biz.SystemSecurityUsercase
	log l.Logger
}

// NewObservabilityService is provider for CopilotService, and need biz.CopilotService
func NewCopilotService(ouc *biz.SystemSecurityUsercase, logger l.Logger) *CopilotService {
	return &CopilotService{
		ouc: ouc,
		log: logger,
	}
}

package service

import (
	"github.com/xiak/grafana-app-with-backend/pkg/internal/biz"

	"github.com/google/wire"
	l "github.com/grafana/grafana-plugin-sdk-go/backend/log"
)

var ProviderSet = wire.NewSet(NewCopilotService)

type CopilotService struct {
	cuc  *biz.ChatUsecase
	ouc  *biz.SystemSecurityUsercase
	ruc  *biz.RagUsecase
	hauc *biz.HostActivityUsecase
	hsuc *biz.HostStateUsecase
	log  l.Logger
}

// NewObservabilityService is provider for CopilotService, and need biz.CopilotService
func NewCopilotService(
	ouc *biz.SystemSecurityUsercase,
	cuc *biz.ChatUsecase,
	ruc *biz.RagUsecase,
	hauc *biz.HostActivityUsecase,
	hsuc *biz.HostStateUsecase,
	logger l.Logger) *CopilotService {
	return &CopilotService{
		cuc:  cuc,
		ouc:  ouc,
		ruc:  ruc,
		hauc: hauc,
		hsuc: hsuc,
		log:  logger,
	}
}

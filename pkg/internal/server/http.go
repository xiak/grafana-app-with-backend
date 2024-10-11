package server

import (
	l "github.com/grafana/grafana-plugin-sdk-go/backend/log"
	v1 "github.com/xiak/grafana-app-with-backend/pkg/api/copilot/v1"
	"github.com/xiak/grafana-app-with-backend/pkg/common/transport/http"
	"github.com/xiak/grafana-app-with-backend/pkg/internal/service"
)

// NewHTTPServer new an HTTP server.
func NewHTTPServer(logger l.Logger, obs *service.CopilotService) *http.Server {
	var opts = []http.ServerOption{
		http.Middleware(),
	}
	srv := http.NewServer(opts...)
	v1.RegisterCopilotHTTPServer(srv, obs)
	return srv
}

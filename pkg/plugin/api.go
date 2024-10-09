package plugin

import (
	"net/http"

	"github.com/xiak/grafana-app-with-backend/pkg/internal/service"
)

// registerRoutes takes a *http.ServeMux and registers HTTP handlers, including custom endpoints.
func RegisterRoutes(obs *service.ObservabilityService) *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("POST /vvault/obs/security", obs.GetUserAction)
	return mux
}

//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package plugin

import (
	"net/http"

	conf "github.com/xiak/grafana-app-with-backend/pkg/config"
	"github.com/xiak/grafana-app-with-backend/pkg/internal/biz"
	"github.com/xiak/grafana-app-with-backend/pkg/internal/data"
	"github.com/xiak/grafana-app-with-backend/pkg/internal/service"

	"github.com/google/wire"
	"github.com/grafana/grafana-plugin-sdk-go/backend/log"
)

// wireApp init application.
func wireApp(*conf.Data, log.Logger) (*http.ServeMux, func(), error) {
	panic(wire.Build(service.ProviderSet, biz.ProviderSet, data.ProviderSet, RegisterRoutes))
}

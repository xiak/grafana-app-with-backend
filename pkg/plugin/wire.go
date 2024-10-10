//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package plugin

import (
	"github.com/gorilla/mux"
	"github.com/xiak/grafana-app-with-backend/pkg/internal/biz"
	conf "github.com/xiak/grafana-app-with-backend/pkg/internal/config"
	"github.com/xiak/grafana-app-with-backend/pkg/internal/data"
	"github.com/xiak/grafana-app-with-backend/pkg/internal/server"
	"github.com/xiak/grafana-app-with-backend/pkg/internal/service"

	"github.com/google/wire"
	l "github.com/grafana/grafana-plugin-sdk-go/backend/log"
)

// wireApp init application.
func wireApp(*conf.Data, l.Logger) (*mux.Router, func(), error) {
	panic(wire.Build(server.ProviderSet, service.ProviderSet, biz.ProviderSet, data.ProviderSet, GetRouter))
}

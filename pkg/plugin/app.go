package plugin

import (
	"context"
	"os"

	"github.com/gorilla/mux"
	"github.com/grafana/grafana-plugin-sdk-go/backend"
	"github.com/grafana/grafana-plugin-sdk-go/backend/instancemgmt"
	l "github.com/grafana/grafana-plugin-sdk-go/backend/log"
	"github.com/grafana/grafana-plugin-sdk-go/backend/resource/httpadapter"
	"github.com/xiak/grafana-app-with-backend/pkg/common/config"
	"github.com/xiak/grafana-app-with-backend/pkg/common/config/file"
	httpex "github.com/xiak/grafana-app-with-backend/pkg/common/transport/http"
	conf "github.com/xiak/grafana-app-with-backend/pkg/internal/config"
)

// Make sure App implements required interfaces. This is important to do
// since otherwise we will only get a not implemented error response from plugin in
// runtime. Plugin should not implement all these interfaces - only those which are
// required for a particular task.
var (
	_ backend.CallResourceHandler   = (*App)(nil)
	_ instancemgmt.InstanceDisposer = (*App)(nil)
	_ backend.CheckHealthHandler    = (*App)(nil)
)

func GetRouter(srv *httpex.Server) *mux.Router {
	return srv.Router()
}

// App is an example app backend plugin which can respond to data queries.
type App struct {
	backend.CallResourceHandler
}

// NewApp creates a new example *App instance.
func NewApp(_ context.Context, _ backend.AppInstanceSettings) (instancemgmt.Instance, error) {
	var app App
	logger := l.DefaultLogger
	configFile := "/etc/grafana/provisioning/config.yaml"

	c := config.New(
		config.WithSource(
			// The path is the same as docker-compose.yaml if you use docker compose
			// volumes:
			// - ./dist:/var/lib/grafana/plugins/app-with-backend
			// - ./provisioning:/etc/grafana/provisioning
			file.NewSource(configFile),
		),
	)
	defer c.Close()
	if err := c.Load(); err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}

	logger.Info("Config file loded successful:", configFile)

	var bs conf.Bootstrap
	if err := c.Scan(&bs); err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}

	router, cleanup, err := wireApp(bs.Data, bs.Chat, bs.Rag, logger)
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}
	defer cleanup()

	app.CallResourceHandler = httpadapter.New(router)

	return &app, nil
}

// Dispose here tells plugin SDK that plugin wants to clean up resources when a new instance
// created.
func (a *App) Dispose() {
	// cleanup
}

// CheckHealth handles health checks sent from Grafana to the plugin.
func (a *App) CheckHealth(_ context.Context, _ *backend.CheckHealthRequest) (*backend.CheckHealthResult, error) {
	return &backend.CheckHealthResult{
		Status:  backend.HealthStatusOk,
		Message: "ok",
	}, nil
}

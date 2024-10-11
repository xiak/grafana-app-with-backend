package plugin

import (
	"context"

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

	// Use a httpadapter (provided by the SDK) for resource calls. This allows us
	// to use a *http.ServeMux for resource calls, so we can map multiple routes
	// to CallResource without having to implement extra logic.
	// workDir, _ := os.Getwd()
	c := config.New(
		config.WithSource(
			// file.NewSource(filepath.Join(workDir, "app-with-backend", "provisioning", "config.yaml")),
			// The path is the same as docker-compose.yaml if you use docker compose
			// volumes:
			// - ./dist:/var/lib/grafana/plugins/app-with-backend
			// - ./provisioning:/etc/grafana/provisioning
			file.NewSource("/etc/grafana/provisioning"),
		),
	)
	defer c.Close()
	if err := c.Load(); err != nil {
		panic(err)
	}

	var dc conf.Data
	if err := c.Scan(&dc); err != nil {
		panic(err)
	}

	var cc conf.Chat
	if err := c.Scan(&cc); err != nil {
		panic(err)
	}
	router, cleanup, err := wireApp(&dc, &cc, l.DefaultLogger)
	if err != nil {
		panic(err)
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

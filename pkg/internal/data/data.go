package data

import (
	"github.com/google/wire"
	l "github.com/grafana/grafana-plugin-sdk-go/backend/log"
	prom "github.com/prometheus/client_golang/api"
	promAPIV1 "github.com/prometheus/client_golang/api/prometheus/v1"

	conf "github.com/xiak/grafana-app-with-backend/pkg/config"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewPromAPI, NewSystemSecurityRepo)

// Data .
type Data struct {
	// prometheus client
	promAPI promAPIV1.API
	log     l.Logger
}

// NewData .
func NewData(pAPI promAPIV1.API, logger l.Logger) (*Data, func(), error) {
	d := &Data{
		promAPI: pAPI,
		log:     logger,
	}
	return d, func() {}, nil
}

// Prometheus API client
// Useage:
//
//	api := NewPromAPI(conf, logger)
//	query := `node_cpu_seconds_total`
//	result, warnings, err := api.Query(context.Background(), query, time.Now())
func NewPromAPI(conf *conf.Data, logger l.Logger) promAPIV1.API {
	promClient, err := prom.NewClient(
		prom.Config{
			Address: conf.Prometheus.Url,
		},
	)

	if err != nil {
		logger.Error("init prometheus client failed! %s", err)
		panic(err)
	}

	return promAPIV1.NewAPI(promClient)

}

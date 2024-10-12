package data

import (
	"github.com/google/wire"
	l "github.com/grafana/grafana-plugin-sdk-go/backend/log"

	cfg "github.com/xiak/grafana-app-with-backend/pkg/internal/config"
)

// ProviderSet is data providers.
// var ProviderSet = wire.NewSet(NewData, NewPromAPI, NewSystemSecurityRepo)
var ProviderSet = wire.NewSet(
	NewData,
	NewSystemSecurityRepo,
	NewChatRepo,
	NewHostActivityRepo,
	NewHostStateRepo,
	NewRagRepo,
	NewNetworkErrorRepo,
	NewCpuStateRepo,
)

// Data .
type Data struct {
	data *cfg.Data
	chat *cfg.Chat
	rag  *cfg.Rag
	// prometheus client
	// promAPI promAPIV1.API
	log l.Logger
}

// NewData .
// func NewData(pAPI promAPIV1.API, logger l.Logger) (*Data, func(), error) {
func NewData(data *cfg.Data, chat *cfg.Chat, rag *cfg.Rag, logger l.Logger) (*Data, func(), error) {
	d := &Data{
		data: data,
		chat: chat,
		rag:  rag,
		// promAPI: pAPI,
		log: logger,
	}
	return d, func() {}, nil
}

// Prometheus API client
// Useage:
//
//	api := NewPromAPI(conf, logger)
//	query := `node_cpu_seconds_total`
//	result, warnings, err := api.Query(context.Background(), query, time.Now())
// func NewPromAPI(conf *cfg.Data, logger l.Logger) promAPIV1.API {
// 	promClient, err := prom.NewClient(
// 		prom.Config{
// 			Address: conf.Prometheus.Url,
// 		},
// 	)

// 	if err != nil {
// 		logger.Error("init prometheus client failed! %s", err)
// 		panic(err)
// 	}

// 	return promAPIV1.NewAPI(promClient)

// }

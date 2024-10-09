package service

import (
	"context"
	"encoding/json"

	"fmt"
	"net/http"

	"github.com/xiak/grafana-app-with-backend/pkg/internal/biz"

	"github.com/google/wire"
	l "github.com/grafana/grafana-plugin-sdk-go/backend/log"
)

var ProviderSet = wire.NewSet(NewObservabilityService)

type ObservabilityService struct {
	ouc *biz.SystemSecurityUsercase
	log l.Logger
}

// NewObservabilityService is provider for ObservabilityService, and need biz.ObservabilityUsercase
func NewObservabilityService(ouc *biz.SystemSecurityUsercase, logger l.Logger) *ObservabilityService {
	return &ObservabilityService{
		ouc: ouc,
		log: logger,
	}
}

func (s *ObservabilityService) GetUserAction(reply http.ResponseWriter, req *http.Request) {
	var err error
	if req.Method != "POST" {
		http.Error(reply, fmt.Sprintf("error http method. request: POST, now: %s", req.Method), http.StatusMethodNotAllowed)
		return
	}
	type Request struct {
		StartTime int64 `json:"start_time"`
		EndTime   int64 `json:"end_time"`
	}
	var request Request
	if err = json.NewDecoder(req.Body).Decode(&request); err != nil {
		http.Error(reply, err.Error(), http.StatusBadRequest)
		return
	}
	type Response struct {
		Status string            `json:"status"`
		Data   []*biz.UserAction `json:"data"`
	}
	data, err := s.ouc.GetSystemSecurity(context.Background(), request.StartTime, request.EndTime)
	if err != nil {
		http.Error(reply, err.Error(), http.StatusInternalServerError)
	}

	resp := &Response{
		Data: data,
	}

	jsonResponse, err := json.Marshal(resp)
	if err != nil {
		http.Error(reply, err.Error(), http.StatusInternalServerError)
		return
	}
	reply.Header().Set("Content-Type", "application/json")
	reply.WriteHeader(http.StatusOK)
	reply.Write(jsonResponse)
}

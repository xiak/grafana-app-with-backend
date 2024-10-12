// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v1.0.0
// - protoc             v5.28.0
// source: api.proto

package v1

import (
	context "context"
	http "github.com/xiak/grafana-app-with-backend/pkg/common/transport/http"
	binding "github.com/xiak/grafana-app-with-backend/pkg/common/transport/http/binding"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the package it is being compiled against.
var _ = new(context.Context)
var _ = binding.EncodeURL

const _ = http.SupportPackageIsVersion1

const OperationCopilotGetCpuState = "/ads.service.copilot.v1.Copilot/GetCpuState"
const OperationCopilotGetHostsActivity = "/ads.service.copilot.v1.Copilot/GetHostsActivity"
const OperationCopilotGetHostsState = "/ads.service.copilot.v1.Copilot/GetHostsState"
const OperationCopilotGetNetworkError = "/ads.service.copilot.v1.Copilot/GetNetworkError"
const OperationCopilotGetRagKeywords = "/ads.service.copilot.v1.Copilot/GetRagKeywords"
const OperationCopilotGetSecurityAbnormalUser = "/ads.service.copilot.v1.Copilot/GetSecurityAbnormalUser"
const OperationCopilotPromptSuggestion = "/ads.service.copilot.v1.Copilot/PromptSuggestion"

type CopilotHTTPServer interface {
	GetCpuState(context.Context, *GetCpuStateRequest) (*GetCpuStateReply, error)
	GetHostsActivity(context.Context, *GetHostsActivityRequest) (*GetHostsActivityReply, error)
	GetHostsState(context.Context, *GetHostsStateRequest) (*GetHostsStateReply, error)
	GetNetworkError(context.Context, *GetNetworkErrorRequest) (*GetNetworkErrorReply, error)
	GetRagKeywords(context.Context, *GetRagKeywordsRequest) (*GetRagKeywordsReply, error)
	GetSecurityAbnormalUser(context.Context, *GetSecurityAbnormalUserRequest) (*GetSecurityAbnormalUserReply, error)
	PromptSuggestion(context.Context, *PromptSuggestionRequest) (*PromptSuggestionReply, error)
}

func RegisterCopilotHTTPServer(s *http.Server, srv CopilotHTTPServer) {
	r := s.Route("/")
	r.POST("/copilot/prompt/suggestion", _Copilot_PromptSuggestion0_HTTP_Handler(srv))
	r.POST("/copilot/security/user", _Copilot_GetSecurityAbnormalUser0_HTTP_Handler(srv))
	r.POST("/copilot/obs/hosts/state", _Copilot_GetHostsState0_HTTP_Handler(srv))
	r.POST("/copilot/obs/hosts/activity", _Copilot_GetHostsActivity0_HTTP_Handler(srv))
	r.POST("/copilot/rag/search/keyword", _Copilot_GetRagKeywords0_HTTP_Handler(srv))
	r.POST("/copilot/obs/network/error", _Copilot_GetNetworkError0_HTTP_Handler(srv))
	r.POST("/copilot/obs/cpu/state", _Copilot_GetCpuState0_HTTP_Handler(srv))
}

func _Copilot_PromptSuggestion0_HTTP_Handler(srv CopilotHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in PromptSuggestionRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationCopilotPromptSuggestion)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.PromptSuggestion(ctx, req.(*PromptSuggestionRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*PromptSuggestionReply)
		return ctx.Result(200, reply)
	}
}

func _Copilot_GetSecurityAbnormalUser0_HTTP_Handler(srv CopilotHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetSecurityAbnormalUserRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationCopilotGetSecurityAbnormalUser)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetSecurityAbnormalUser(ctx, req.(*GetSecurityAbnormalUserRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetSecurityAbnormalUserReply)
		return ctx.Result(200, reply)
	}
}

func _Copilot_GetHostsState0_HTTP_Handler(srv CopilotHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetHostsStateRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationCopilotGetHostsState)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetHostsState(ctx, req.(*GetHostsStateRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetHostsStateReply)
		return ctx.Result(200, reply)
	}
}

func _Copilot_GetHostsActivity0_HTTP_Handler(srv CopilotHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetHostsActivityRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationCopilotGetHostsActivity)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetHostsActivity(ctx, req.(*GetHostsActivityRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetHostsActivityReply)
		return ctx.Result(200, reply)
	}
}

func _Copilot_GetRagKeywords0_HTTP_Handler(srv CopilotHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetRagKeywordsRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationCopilotGetRagKeywords)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetRagKeywords(ctx, req.(*GetRagKeywordsRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetRagKeywordsReply)
		return ctx.Result(200, reply)
	}
}

func _Copilot_GetNetworkError0_HTTP_Handler(srv CopilotHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetNetworkErrorRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationCopilotGetNetworkError)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetNetworkError(ctx, req.(*GetNetworkErrorRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetNetworkErrorReply)
		return ctx.Result(200, reply)
	}
}

func _Copilot_GetCpuState0_HTTP_Handler(srv CopilotHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetCpuStateRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationCopilotGetCpuState)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetCpuState(ctx, req.(*GetCpuStateRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetCpuStateReply)
		return ctx.Result(200, reply)
	}
}

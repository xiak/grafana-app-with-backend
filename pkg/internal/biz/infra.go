package biz

import (
	"context"

	l "github.com/grafana/grafana-plugin-sdk-go/backend/log"
)

type HostActivity struct {
	Id                 string
	Time               string
	HostNumber         int64
	HostActivityNumber int64
	HostIdleNumber     int64
	HostActivityIps    string
}

type HostState struct {
	Id        string
	Time      string
	AppName   string
	Host      string
	Operation string
	AvgOpTime int64
}

type HostActivityRepo interface {
	GetHostsActivity(ctx context.Context, start int64, end int64) ([]*HostActivity, error)
}

type HostStateRepo interface {
	GetHostsState(ctx context.Context, start int64, end int64) ([]*HostState, error)
}

type HostActivityUsecase struct {
	repo HostActivityRepo
	log  l.Logger
}

type HostStateUsecase struct {
	repo HostStateRepo
	log  l.Logger
}

func NewHostActivityUsecase(repo HostActivityRepo, logger l.Logger) *HostActivityUsecase {
	return &HostActivityUsecase{
		repo: repo,
		log:  logger,
	}
}

func (uc *HostActivityUsecase) GetHostsActivity(ctx context.Context, start int64, end int64) ([]*HostActivity, error) {
	return uc.repo.GetHostsActivity(ctx, start, end)
}

func NewHostStateUsecase(repo HostStateRepo, logger l.Logger) *HostStateUsecase {
	return &HostStateUsecase{
		repo: repo,
		log:  logger,
	}
}

func (uc *HostStateUsecase) GetHostsState(ctx context.Context, start int64, end int64) ([]*HostState, error) {
	return uc.repo.GetHostsState(ctx, start, end)
}

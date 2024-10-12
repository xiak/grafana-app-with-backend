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

type NetworkError struct {
	Time     string
	Port     string
	RxDrops  int64
	TxDrops  int64
	RxErrors int64
	TxErrors int64
}

type CpuState struct {
	Util  *CpuUtilization
	Usage []*CpuUsage
}

type CpuUsage struct {
	AppName string
	Usage   string
}

type CpuUtilization struct {
	Server               string
	AverageUtilization   *CpuAverageUtilization
	PeakUtilization      *CpuPeakUtilization
	BreakdownUtilization *CpuBreakdownUtilization
}

type CpuAverageUtilization struct {
	Value string
}

type CpuPeakUtilization struct {
	Value string
	Time  string
}

type CpuBreakdownUtilization struct {
	UserTime   string
	SystemTime string
}

type HostActivityUsecase struct {
	repo HostActivityRepo
	log  l.Logger
}

type HostStateUsecase struct {
	repo HostStateRepo
	log  l.Logger
}

type NetworkErrorUsecase struct {
	repo NetworkErrorRepo
	log  l.Logger
}

type CpuStateUsecase struct {
	repo CpuStateRepo
	log  l.Logger
}

type HostActivityRepo interface {
	GetHostsActivity(ctx context.Context, start int64, end int64) ([]*HostActivity, error)
}

type HostStateRepo interface {
	GetHostsState(ctx context.Context, start int64, end int64) ([]*HostState, error)
}

type NetworkErrorRepo interface {
	GetNetworkError(ctx context.Context, start int64, end int64) ([]*NetworkError, error)
}

type CpuStateRepo interface {
	GetCpuState(ctx context.Context, start int64, end int64) (*CpuState, error)
}

func NewHostActivityUsecase(repo HostActivityRepo, logger l.Logger) *HostActivityUsecase {
	return &HostActivityUsecase{
		repo: repo,
		log:  logger,
	}
}

func NewHostStateUsecase(repo HostStateRepo, logger l.Logger) *HostStateUsecase {
	return &HostStateUsecase{
		repo: repo,
		log:  logger,
	}
}

func NewNetworkUsecase(repo NetworkErrorRepo, logger l.Logger) *NetworkErrorUsecase {
	return &NetworkErrorUsecase{
		repo: repo,
		log:  logger,
	}
}

func NewCpuStateUsecase(repo CpuStateRepo, logger l.Logger) *CpuStateUsecase {
	return &CpuStateUsecase{
		repo: repo,
		log:  logger,
	}
}

func (uc *HostActivityUsecase) GetHostsActivity(ctx context.Context, start int64, end int64) ([]*HostActivity, error) {
	return uc.repo.GetHostsActivity(ctx, start, end)
}

func (uc *HostStateUsecase) GetHostsState(ctx context.Context, start int64, end int64) ([]*HostState, error) {
	return uc.repo.GetHostsState(ctx, start, end)
}

func (uc *NetworkErrorUsecase) GetNetworkError(ctx context.Context, start int64, end int64) ([]*NetworkError, error) {
	return uc.repo.GetNetworkError(ctx, start, end)
}

func (uc *CpuStateUsecase) GetCpuState(ctx context.Context, start int64, end int64) (*CpuState, error) {
	return uc.repo.GetCpuState(ctx, start, end)
}

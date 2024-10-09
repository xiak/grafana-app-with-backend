package biz

import (
	"context"

	l "github.com/grafana/grafana-plugin-sdk-go/backend/log"
)

type SystemSecurityRepo interface {
	GetUserActionFromTime2Time(ctx context.Context, start int64, end int64) ([]*UserAction, error)
}

type SystemSecurityUsercase struct {
	repo SystemSecurityRepo
	log  l.Logger
}

// Directory Access Log
type Journal struct {
	// which directory
	Directory string
	// Directory Access Status
	Status string
	// Directory Access Message
	Message string
}

// Detecting user behavior over a time period
type UserAction struct {
	// Login username
	User string
	// Standard, Admin, Guest, others...
	Privilege string
	// Login status
	Status string
	// Login message
	LoginMessage string
	// Directory Access Logs
	Journal []*Journal
	// Login time
	LoginTime int64
	// Logout time
	LogoutTime int64
	// Online time
	Online int64
}

func NewSystemSecurityUsercase(repo SystemSecurityRepo, logger l.Logger) *SystemSecurityUsercase {
	return &SystemSecurityUsercase{
		repo: repo,
		log:  logger,
	}
}

func (uc *SystemSecurityUsercase) GetSystemSecurity(ctx context.Context, start int64, end int64) ([]*UserAction, error) {
	return uc.repo.GetUserActionFromTime2Time(ctx, start, end)
}

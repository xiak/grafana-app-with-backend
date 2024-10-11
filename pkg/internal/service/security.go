package service

import (
	"context"

	v1 "github.com/xiak/grafana-app-with-backend/pkg/api/copilot/v1"
)

func (s *CopilotService) GetSecurityAbnormalUser(ctx context.Context, req *v1.GetSecurityAbnormalUserRequest) (*v1.GetSecurityAbnormalUserReply, error) {
	data, err := s.ouc.GetSystemSecurity(ctx, req.StartTime, req.EndTime)
	if err != nil {
		return nil, err
	}
	au := make([]*v1.AbnormalUser, 0)
	j := make([]*v1.Journal, 0)
	for _, d := range data {
		if d == nil {
			continue
		}
		for _, dj := range d.Journal {
			j = append(j, &v1.Journal{
				Directory: dj.Directory,
				Status:    dj.Status,
				Message:   dj.Message,
			})
		}
		au = append(au, &v1.AbnormalUser{
			User:         d.User,
			Privilege:    d.Privilege,
			Status:       d.Status,
			LoginMessage: d.LoginMessage,
			Journal:      j,
			LoginTime:    d.LoginTime,
			LogoutTime:   d.LogoutTime,
			Online:       d.Online,
		})
	}
	return &v1.GetSecurityAbnormalUserReply{
		Domain:       "system_safety",
		AbnormalUser: au,
	}, nil
}

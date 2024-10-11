package service

import (
	"context"

	v1 "github.com/xiak/grafana-app-with-backend/pkg/api/copilot/v1"
)

func (s *CopilotService) GetHostsActivity(ctx context.Context, req *v1.GetHostsActivityRequest) (*v1.GetHostsActivityReply, error) {
	hs, err := s.hauc.GetHostsActivity(ctx, req.StartTime, req.EndTime)
	if err != nil {
		return nil, err
	}
	hostsActivity := make([]*v1.HostActivity, 0)
	for _, h := range hs {
		hostsActivity = append(hostsActivity, &v1.HostActivity{
			RecordId:        h.Id,
			Timestamp:       h.Time,
			HostNum:         h.HostNumber,
			HostActivityNum: h.HostActivityNumber,
			HostIdleNum:     h.HostIdleNumber,
			HostActivityIps: h.HostActivityIps,
		})
	}
	return &v1.GetHostsActivityReply{
		Domain:        "host_resource",
		HostsActivity: hostsActivity,
	}, nil
}

func (s *CopilotService) GetHostsState(ctx context.Context, req *v1.GetHostsStateRequest) (*v1.GetHostsStateReply, error) {
	hs, err := s.hsuc.GetHostsState(ctx, req.StartTime, req.EndTime)
	s.log.Error("===>", hs)
	if err != nil {
		return nil, err
	}
	hostsState := make([]*v1.HostState, 0)
	for _, h := range hs {
		s.log.Error("===>", h)
		hostsState = append(hostsState, &v1.HostState{
			RecordId:    h.Id,
			Timestamp:   h.Time,
			Application: h.AppName,
			Host:        h.Host,
			Operation:   h.Operation,
			AvgOpTime:   h.AvgOpTime,
		})
	}
	return &v1.GetHostsStateReply{
		Domain:     "system_status",
		HostsState: hostsState,
	}, nil
}

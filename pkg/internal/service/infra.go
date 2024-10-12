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
	if err != nil {
		return nil, err
	}
	hostsState := make([]*v1.HostState, 0)
	for _, h := range hs {

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

func (s *CopilotService) GetNetworkError(ctx context.Context, req *v1.GetNetworkErrorRequest) (*v1.GetNetworkErrorReply, error) {
	nes, err := s.neuc.GetNetworkError(ctx, req.StartTime, req.EndTime)
	if err != nil {
		return nil, err
	}
	networkError := make([]*v1.NetworkError, 0)
	for _, n := range nes {
		networkError = append(networkError, &v1.NetworkError{
			Timestamp: n.Time,
			Port:      n.Port,
			RxDrops:   n.RxDrops,
			TxDrops:   n.TxDrops,
			RxErrors:  n.RxErrors,
			TxErrors:  n.TxErrors,
		})
	}
	return &v1.GetNetworkErrorReply{
		Domain:       "network_error",
		NetworkError: networkError,
	}, nil
}

func (s *CopilotService) GetCpuState(ctx context.Context, req *v1.GetCpuStateRequest) (*v1.GetCpuStateReply, error) {
	data, err := s.cpuuc.GetCpuState(ctx, req.StartTime, req.EndTime)
	if err != nil {
		return nil, err
	}
	cpuUsage := make([]*v1.CpuUsage, 0)

	for _, d := range data.Usage {
		cpuUsage = append(cpuUsage, &v1.CpuUsage{
			AppName: d.AppName,
			Usage:   d.Usage,
		})
	}
	return &v1.GetCpuStateReply{
		Domain: "cpu_status",
		CpuState: &v1.CpuState{
			CpuUtilization: &v1.CpuUtilization{
				Server: data.Util.Server,
				AverageUtilization: &v1.AverageUtilization{
					Value: data.Util.AverageUtilization.Value,
				},
				PeakUtilization: &v1.PeakUtilization{
					Value: data.Util.PeakUtilization.Value,
					Time:  data.Util.PeakUtilization.Time,
				},
				BreakdownUtilization: &v1.BreakdownUtilization{
					UserTime:   data.Util.BreakdownUtilization.UserTime,
					SystemTime: data.Util.BreakdownUtilization.SystemTime,
				},
			},
			CpuUsage: cpuUsage,
		},
	}, nil
}

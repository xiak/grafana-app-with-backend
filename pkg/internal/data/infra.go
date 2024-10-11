package data

import (
	"context"
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"github.com/xiak/grafana-app-with-backend/pkg/internal/biz"

	l "github.com/grafana/grafana-plugin-sdk-go/backend/log"
)

var _ biz.HostActivityRepo = (*hostActivityRepo)(nil)
var _ biz.HostStateRepo = (*hostStateRepo)(nil)

type hostActivityRepo struct {
	data *Data
	log  l.Logger
}

type hostStateRepo struct {
	data *Data
	log  l.Logger
}

func NewHostActivityRepo(data *Data, logger l.Logger) biz.HostActivityRepo {
	return &hostActivityRepo{
		data: data,
		log:  logger,
	}
}

func (repo *hostActivityRepo) GetHostsActivity(ctx context.Context, start int64, end int64) ([]*biz.HostActivity, error) {
	timeFormat := "2006-01-02_15:04:05"
	startTime := time.Unix(start, 0)
	endTime := time.Unix(end, 0)
	if endTime.Before(startTime) {
		return nil, fmt.Errorf("the starttime [%s] is greater than endtime [%s]", startTime.Format(timeFormat), endTime.Format(timeFormat))
	}
	// Fake data for test
	count := 20

	ha := make([]*biz.HostActivity, 0)
	var tmpTime int64
	for i := 0; i < count; i++ {
		tmpTime = start
		start = start + 1
		startTime := time.Unix(start, 0)
		if startTime.Before(endTime) {
			start = tmpTime
		}
		var d int = 5
		ip1 := rand.Intn(100)
		ip2 := ip1 + d
		ips := fmt.Sprintf("%d.%d.%d.[%d-%d]", rand.Intn(193), rand.Intn(193), rand.Intn(193), ip1, ip2)
		ha = append(ha, &biz.HostActivity{
			Id:                 strconv.Itoa(i),
			Time:               time.Unix(start, 0).Format(timeFormat),
			HostNumber:         50,
			HostActivityNumber: int64(d),
			HostIdleNumber:     45,
			HostActivityIps:    ips,
		})
	}
	return ha, nil
}

func NewHostStateRepo(data *Data, logger l.Logger) biz.HostStateRepo {
	return &hostStateRepo{
		data: data,
		log:  logger,
	}
}

func (repo *hostStateRepo) GetHostsState(ctx context.Context, start int64, end int64) ([]*biz.HostState, error) {
	timeFormat := "2006-01-02_15:04:05"
	startTime := time.Unix(start, 0)
	endTime := time.Unix(end, 0)
	if endTime.Before(startTime) {
		return nil, fmt.Errorf("the starttime [%s] is greater than endtime [%s]", startTime.Format(timeFormat), endTime.Format(timeFormat))
	}
	// Fake data for test
	count := 20
	apps := []string{"chrome", "ps", "docker", "top", "ntpd", "netstat", "kube-apiserver", "kube-controller-manager", "kube-scheduler", "kubelet"}
	hostname := []string{"host-1", "host-2", "host-3", "host-4", "host-5"}
	ops := []string{"读", "写"}
	hs := make([]*biz.HostState, 0)
	var tmpTime int64
	for i := 0; i < count; i++ {
		tmpTime = start
		start = start + 1
		startTime := time.Unix(start, 0)
		if startTime.Before(endTime) {
			start = tmpTime
		}
		hs = append(hs, &biz.HostState{
			Id:        strconv.Itoa(i),
			Time:      time.Unix(start, 0).Format(timeFormat),
			AppName:   apps[rand.Intn(len(apps)-1)],
			Host:      hostname[rand.Intn(len(hostname)-1)],
			Operation: ops[rand.Intn(len(ops)-1)],
			AvgOpTime: int64(rand.Intn(100000)),
		})
	}
	repo.log.Error("===========================>", hs)
	return hs, nil
}

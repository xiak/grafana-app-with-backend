package data

import (
	"context"

	"github.com/xiak/grafana-app-with-backend/pkg/internal/biz"

	l "github.com/grafana/grafana-plugin-sdk-go/backend/log"
)

var _ biz.SystemSecurityRepo = (*systemSecurityRepo)(nil)

type systemSecurityRepo struct {
	data *Data
	log  l.Logger
}

func NewSystemSecurityRepo(data *Data, logger l.Logger) biz.SystemSecurityRepo {
	return &systemSecurityRepo{
		data: data,
		log:  logger,
	}
}

func (ssr *systemSecurityRepo) GetUserActionFromTime2Time(ctx context.Context, start int64, end int64) ([]*biz.UserAction, error) {
	// query := `rate(system_security_user_action{mode="idle"}[5m])`
	// result, warnings, err := ssr.data.promAPI.Query(ctx, query, time.Now())
	// if err != nil {
	// 	ssr.log.Error("query failed: ", err)
	// 	return nil, err
	// }
	// if len(warnings) > 0 {
	// 	ssr.log.Warn("query failed: ", warnings)
	// }

	// ssr.log.Info("ResultTypeis: ", result.Type())
	// vector, ok := result.(model.Vector)
	// if !ok {
	// 	ssr.log.Info("result is not a vector")
	// 	return nil, errors.New("result is not a vector")
	// }
	// for _, sample := range vector {
	// 	log.Infof("Metrics: %s, Value: %s\n", sample.String(), sample.Value)
	// }

	// Fake data for test
	uaList := make([]*biz.UserAction, 10)
	j := make([]*biz.Journal, 0)
	for i := 0; i < 10; i++ {
		j1 := &biz.Journal{
			Directory: "/etc/hosts",
			Status:    "正常",
			Message:   "",
		}
		j2 := &biz.Journal{
			Directory: "/home/xiak",
			Status:    "正常",
			Message:   "",
		}
		j = append(j, j1, j2)
		uaList = append(uaList, &biz.UserAction{
			User:         "user10",
			Privilege:    "Standard",
			Status:       "正常",
			LoginMessage: "",
			Journal:      j,
			LoginTime:    0,
			LogoutTime:   3000,
			Online:       3000,
		})
	}
	return uaList, nil
}

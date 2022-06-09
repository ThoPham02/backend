package topic

import (
	"back-end/cmd/database/repo"
	"back-end/cmd/svc"
	"back-end/cmd/types"
	"context"
	"database/sql"
	"github.com/go-kratos/kratos/v2/log"
)

type TopicGetAllLogic struct {
	ctx       context.Context
	svcCtx    *svc.ServiceContext
	logHelper *log.Helper
}

func NewTopicGetAllLogic(ctx context.Context, svcCtx *svc.ServiceContext, logHelper *log.Helper) TopicGetAllLogic {
	return TopicGetAllLogic{
		ctx:       ctx,
		svcCtx:    svcCtx,
		logHelper: logHelper,
	}
}

func (l *TopicGetAllLogic) TopicGetAll() (*types.TopicGetAllResponse, error) {
	l.logHelper.Infof("Start process get topic")
	condition := &repo.GetCondition{}
	topics, err := l.svcCtx.TopicRepo.GetTopic(l.ctx, condition)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		l.logHelper.Errorf("Failed while get topic, error: %s", err.Error())
		return nil, err
	}

	var dataResponse []types.Topic
	for _, value := range topics {
		dataResponse = append(dataResponse, types.Topic{
			TopicID:        value.TopicID,
			TopicName:      value.TopicName,
			GroupStudentID: value.GroupStudentID,
			LecturerID:     value.LecturerID,
			StartDay:       value.StartDay,
			EndDay:         value.EndDay,
			Allowance:      value.Allowance,
			TopicStatus:    value.TopicStatus,
		})
	}

	return &types.TopicGetAllResponse{
		TopicResponse: dataResponse,
	}, nil
}

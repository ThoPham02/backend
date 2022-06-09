package topic

import (
	"back-end/cmd/database/model"
	"back-end/cmd/svc"
	"back-end/cmd/types"
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

type TopicUpdateLogic struct {
	ctx       context.Context
	svcCtx    *svc.ServiceContext
	logHelper *log.Helper
}

func NewTopicUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext, logHelper *log.Helper) TopicUpdateLogic {
	return TopicUpdateLogic{
		ctx:       ctx,
		svcCtx:    svcCtx,
		logHelper: logHelper,
	}
}

func (l *TopicUpdateLogic) TopicUpdate(input *types.UpdateTopicInput) error {
	l.logHelper.Infof("Start process get topic")

	topic := &model.Topic{
		TopicID:        input.Path.Id,
		TopicName:      input.Request.TopicName,
		GroupStudentID: input.Request.GroupStudentID,
		LecturerID:     input.Request.LecturerID,
		StartDay:       input.Request.StartDay,
		EndDay:         input.Request.EndDay,
		Allowance:      input.Request.Allowance,
		TopicStatus:    input.Request.TopicStatus,
	}

	err := l.svcCtx.TopicRepo.CreateTopic(l.ctx, topic)
	if err != nil {
		l.logHelper.Errorf("Failed while insert topic, error: %s", err.Error())
		return err
	}
	return nil

}

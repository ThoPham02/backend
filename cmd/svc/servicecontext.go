package svc

import (
	"back-end/cmd/config"
	"back-end/cmd/database/repo"
)

type ServiceContext struct {
	Config    config.Config
	TopicRepo repo.TopicRepo
}

func NewServiceContext(c config.Config, topicRepo repo.TopicRepo) *ServiceContext {

	return &ServiceContext{
		Config:    c,
		TopicRepo: topicRepo,
	}
}

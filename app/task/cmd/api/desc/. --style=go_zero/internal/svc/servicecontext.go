package svc

import (
	"calendarAttempt3/app/task/cmd/api/desc/. --style=go_zero/internal/config"
)

type ServiceContext struct {
	Config config.Config
	
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		
	}
}

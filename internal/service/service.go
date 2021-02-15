package service

import (
	"time"

	"github.com/epociask/go-rest-api-template/internal/models"
	"github.com/google/wire"
)

type Service interface {
	CheckHealth() *models.HealthCheck
}

var Module = wire.NewSet(
	New,
)

type ExampleService struct {
}

func New() *ExampleService {
	return &ExampleService{}
}

func (svc *ExampleService) CheckHealth() *models.HealthCheck {
	return &models.HealthCheck{Timestamp: time.Now(), Healthy: true}
}

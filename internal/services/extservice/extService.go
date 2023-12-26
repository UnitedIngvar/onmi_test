package extservice

import (
	"context"
	"time"

	"github.com/UnitedIngvar/onmi_test/internal/services/client"
)

type MyMockService struct {
	countLimits uint64
	timeLimits  time.Duration
}

func NewMyMockService() *MyMockService {
	return &MyMockService{
		countLimits: 100,
		timeLimits:  time.Second}
}

func (s *MyMockService) GetLimits() (n uint64, p time.Duration) {
	return s.countLimits, s.timeLimits
}

func (s *MyMockService) Process(ctx context.Context, batch client.Batch) error {
	return nil
}

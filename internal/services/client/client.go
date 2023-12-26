package client

import (
	"context"
	"errors"
	"time"

	"github.com/UnitedIngvar/onmi_test/internal/utils"
)

// ErrBlocked reports if service is blocked.
var ErrBlocked = errors.New("blocked")

// Service defines external service that can process batches of items.
//
//go:generate mockery --name Service
type Service interface {
	GetLimits() (n uint64, p time.Duration)
	Process(ctx context.Context, batch Batch) error
}

// Batch is a batch of items.
type Batch []Item

// Item is some abstract item.
type Item struct{}

type Client struct {
	service Service
}

func NewClient(service Service) *Client {
	return &Client{
		service: service,
	}
}

func (s *Client) SendRequest(ctx context.Context, itemCount uint64) error {
	timer := time.NewTimer(0)

	for {
		countLimit, timeLimit := s.service.GetLimits()
		timer.Reset(timeLimit)
		batchLen := utils.Min(itemCount, countLimit)
		batch := make(Batch, batchLen)

		if err := s.service.Process(ctx, batch); err != nil {
			return err
		}

		itemCount -= batchLen
		if itemCount <= 0 {
			break
		}

		select {
		case <-ctx.Done():
			return context.Canceled
		case <-timer.C:
		}
	}

	return nil
}

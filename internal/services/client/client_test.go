package client

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func Test_SendRequest_ReturnError_WhenServiceReturnsError(t *testing.T) {
	// Arrange
	serviceMock := NewMockService(t)
	client := NewClient(serviceMock)
	ctx := context.TODO()

	itemCount := uint64(10)
	itemLimit := uint64(100)
	timeLimit := time.Millisecond * 100
	batch := make(Batch, itemCount)

	serviceMock.EXPECT().GetLimits().Return(itemLimit, timeLimit).Once()
	serviceMock.EXPECT().Process(ctx, batch).Return(assert.AnError).Once()

	// Act
	err := client.SendRequest(ctx, itemCount)

	// Assert
	assert.EqualError(t, err, assert.AnError.Error())
}

func Test_SendRequest_Success_WhenItemCountLessThanLimit(t *testing.T) {
	// Arrange
	serviceMock := NewMockService(t)
	client := NewClient(serviceMock)
	ctx := context.TODO()

	itemCount := uint64(10)
	itemLimit := uint64(100)
	timeLimit := time.Millisecond * 100
	batch := make(Batch, itemCount)

	serviceMock.EXPECT().GetLimits().Return(itemLimit, timeLimit).Once()
	serviceMock.EXPECT().Process(ctx, batch).Return(nil).Once()

	// Act
	err := client.SendRequest(ctx, itemCount)

	// Assert
	assert.Nil(t, err)
}

func Test_SendRequest_Success_WhenItemCountEqualToLimit(t *testing.T) {
	// Arrange
	serviceMock := NewMockService(t)
	client := NewClient(serviceMock)
	ctx := context.TODO()

	itemCount := uint64(10)
	itemLimit := uint64(itemCount)
	timeLimit := time.Millisecond * 100
	batch := make(Batch, itemCount)

	serviceMock.EXPECT().GetLimits().Return(itemLimit, timeLimit).Once()
	serviceMock.EXPECT().Process(ctx, batch).Return(nil).Once()

	// Act
	err := client.SendRequest(ctx, itemCount)

	// Assert
	assert.Nil(t, err)
}

func Test_SendRequest_Success_WhenItemCountMoreThanLimit(t *testing.T) {
	// Arrange
	serviceMock := NewMockService(t)
	client := NewClient(serviceMock)
	ctx := context.TODO()

	itemCount := uint64(150)
	itemLimit := uint64(100)
	timeLimit := time.Millisecond * 100
	firstBatch := make(Batch, 100)
	secondBatch := make(Batch, 50)

	serviceMock.EXPECT().GetLimits().Return(itemLimit, timeLimit).Twice()
	serviceMock.EXPECT().Process(ctx, firstBatch).Return(nil).Once()
	serviceMock.EXPECT().Process(ctx, secondBatch).Return(nil).Once()

	// Act
	err := client.SendRequest(ctx, itemCount)

	// Assert
	assert.Nil(t, err)
}

func Test_SendRequest_Success_WhenItemCountMoreThanLimitAndItemLimtsDiffer(t *testing.T) {
	// Arrange
	serviceMock := NewMockService(t)
	client := NewClient(serviceMock)
	ctx := context.TODO()

	itemCount := uint64(150)
	firstItemLimit := uint64(100)
	secondItemLimit := uint64(40)
	thirdItemLimit := uint64(50)
	timeLimit := time.Millisecond * 100
	firstBatch := make(Batch, 100)
	secondBatch := make(Batch, 40)
	thirdBatch := make(Batch, 10)

	serviceMock.EXPECT().GetLimits().Return(firstItemLimit, timeLimit).Once()
	serviceMock.EXPECT().GetLimits().Return(secondItemLimit, timeLimit).Once()
	serviceMock.EXPECT().GetLimits().Return(thirdItemLimit, timeLimit).Once()
	serviceMock.EXPECT().Process(ctx, firstBatch).Return(nil).Once()
	serviceMock.EXPECT().Process(ctx, secondBatch).Return(nil).Once()
	serviceMock.EXPECT().Process(ctx, thirdBatch).Return(nil).Once()

	// Act
	err := client.SendRequest(ctx, itemCount)

	// Assert
	assert.Nil(t, err)
}

func Test_SendRequest_Success_WhenItemCountMoreThanLimitAndAllLimtsDiffer(t *testing.T) {
	// Arrange
	serviceMock := NewMockService(t)
	client := NewClient(serviceMock)
	ctx := context.TODO()

	itemCount := uint64(150)
	firstItemLimit := uint64(100)
	secondItemLimit := uint64(40)
	thirdItemLimit := uint64(50)
	firstTimeLimit := time.Millisecond * 200
	secondTimeLimit := time.Millisecond * 100
	thirdTimeLimit := time.Millisecond * 10
	firstBatch := make(Batch, 100)
	secondBatch := make(Batch, 40)
	thirdBatch := make(Batch, 10)

	serviceMock.EXPECT().GetLimits().Return(firstItemLimit, firstTimeLimit).Once()
	serviceMock.EXPECT().GetLimits().Return(secondItemLimit, secondTimeLimit).Once()
	serviceMock.EXPECT().GetLimits().Return(thirdItemLimit, thirdTimeLimit).Once()
	serviceMock.EXPECT().Process(ctx, firstBatch).Return(nil).Once()
	serviceMock.EXPECT().Process(ctx, secondBatch).Return(nil).Once()
	serviceMock.EXPECT().Process(ctx, thirdBatch).Return(nil).Once()

	// Act
	err := client.SendRequest(ctx, itemCount)

	// Assert
	assert.Nil(t, err)
}

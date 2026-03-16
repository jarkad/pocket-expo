package services

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

// Counter is a Gorm model used by CounterService.
//
// This is kinda funky, but the counter is stored in a separate table
// with only one row.
type Counter struct {
	ID    int
	Count int
}

// CounterService is a service that provides a simple counter.
type CounterService struct {
	DB *gorm.DB
}

// NewCounter creates a new CounterService.
//
// During initialization, database migrations will be run.
//
// TODO: factor out database concerns.
func NewCounter(db *gorm.DB) (*CounterService, error) {
	err := db.AutoMigrate(&Counter{}) //nolint:exhaustruct
	if err != nil {
		return nil, fmt.Errorf("while migrating counter: %w", err)
	}

	err = db.Transaction(func(tx *gorm.DB) error {
		cnt, err := gorm.G[Counter](tx).Count(context.Background(), "*")
		if err != nil {
			return fmt.Errorf("while counting counters: %w", err)
		}

		if cnt == 0 {
			err := gorm.G[Counter](tx).Create(context.Background(), &Counter{
				ID:    1,
				Count: 0,
			})
			if err != nil {
				return fmt.Errorf("while creating counter: %w", err)
			}
		}

		return nil
	})
	if err != nil {
		return nil, fmt.Errorf("while initializing counter: %w", err)
	}

	return &CounterService{
		DB: db,
	}, nil
}

// Get retrieves current value of the counter.
func (cs *CounterService) Get(ctx context.Context) (int, error) {
	counter, err := gorm.G[Counter](cs.DB).
		Where("id = 1").
		Take(ctx)
	if err != nil {
		return 0, fmt.Errorf("while querying counter: %w", err)
	}

	return counter.Count, nil
}

// Inc increments the counter.
func (cs *CounterService) Inc(ctx context.Context) error {
	_, err := gorm.G[Counter](cs.DB).
		Where("id = 1").
		Update(ctx, "count", gorm.Expr("count + 1"))
	if err != nil {
		return fmt.Errorf("while incrementing counter: %w", err)
	}

	return nil
}

// Dec decrements the counter.
func (cs *CounterService) Dec(ctx context.Context) error {
	_, err := gorm.G[Counter](cs.DB).
		Where("id = 1").
		Update(ctx, "count", gorm.Expr("count - 1"))
	if err != nil {
		return fmt.Errorf("while decrementing counter: %w", err)
	}

	return nil
}

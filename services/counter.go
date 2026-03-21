// SPDX-FileCopyrightText: 2026 Jarkad <jarkad@jarkad.net.eu.org>
//
// SPDX-License-Identifier: EUPL-1.2

package services

import (
	"context"
	"fmt"
)

// CounterRepository is a repository that stores a single integer.
type CounterRepository interface {
	Get(ctx context.Context) (int, error)
	Set(ctx context.Context, value int) error
	Update(ctx context.Context, updateFunc func(value int) (int, error)) error
}

// CounterService is a service that provides a simple counter.
type CounterService struct {
	store CounterRepository
}

// NewCounter creates a new CounterService.
//
// During initialization, database migrations will be run.
//
// TODO: factor out database concerns.
func NewCounter(store CounterRepository) (*CounterService, error) {
	return &CounterService{store}, nil
}

// Get retrieves current value of the counter.
func (cs *CounterService) Get(ctx context.Context) (int, error) {
	count, err := cs.store.Get(ctx)
	if err != nil {
		return 0, fmt.Errorf("while querying counter: %w", err)
	}

	return count, nil
}

// Inc increments the counter.
func (cs *CounterService) Inc(ctx context.Context) error {
	err := cs.store.Update(ctx, func(value int) (int, error) {
		return value + 1, nil
	})
	if err != nil {
		return fmt.Errorf("while incrementing counter: %w", err)
	}

	return nil
}

// Dec decrements the counter.
func (cs *CounterService) Dec(ctx context.Context) error {
	err := cs.store.Update(ctx, func(value int) (int, error) {
		return value - 1, nil
	})
	if err != nil {
		return fmt.Errorf("while decrementing counter: %w", err)
	}

	return nil
}

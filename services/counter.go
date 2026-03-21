// SPDX-FileCopyrightText: 2026 Jarkad <jarkad@jarkad.net.eu.org>
//
// SPDX-License-Identifier: EUPL-1.2

package services

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"git.jarkad.net.eu.org/jarkad/pocket-expo/orm/counter"
	"github.com/vinovest/sqlx"
)

// CounterService is a service that provides a simple counter.
type CounterService struct {
	db counter.DBTX
}

// NewCounter creates a new CounterService.
//
// During initialization, database migrations will be run.
//
// TODO: factor out database concerns.
func NewCounter(ctx context.Context, db *sqlx.DB) (*CounterService, error) {
	err := sqlx.TransactContext(ctx, db, func(ctx context.Context, tx sqlx.Queryable) error {
		q := counter.New(tx)

		_, err := q.Get(ctx)
		if errors.Is(err, sql.ErrNoRows) {
			err = q.Reset(ctx)
			if err != nil {
				return err //nolint:wrapcheck // the wrapper is outside this lambda
			}
		} else if err != nil {
			return err //nolint:wrapcheck // the wrapper is outside this lambda
		}

		return nil
	})
	if err != nil {
		return nil, fmt.Errorf("while initializing the counter: %w", err)
	}

	return &CounterService{
		db: db,
	}, nil
}

// Get retrieves current value of the counter.
func (cs *CounterService) Get(ctx context.Context) (int, error) {
	c := counter.New(cs.db)

	count, err := c.Get(ctx)
	if err != nil {
		return 0, fmt.Errorf("while querying counter: %w", err)
	}

	return int(count), nil
}

// Inc increments the counter.
func (cs *CounterService) Inc(ctx context.Context) error {
	c := counter.New(cs.db)

	_, err := c.Increment(ctx)
	if err != nil {
		return fmt.Errorf("while incrementing counter: %w", err)
	}

	return nil
}

// Dec decrements the counter.
func (cs *CounterService) Dec(ctx context.Context) error {
	c := counter.New(cs.db)

	_, err := c.Decrement(ctx)
	if err != nil {
		return fmt.Errorf("while decrementing counter: %w", err)
	}

	return nil
}

package services

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"git.jarkad.net.eu.org/jarkad/pocket-expo/orm/counter"
)

// CounterService is a service that provides a simple counter.
type CounterService struct {
	db counter.DBTX
}

func withTx(ctx context.Context, db *sql.DB, body func(tx *sql.Tx) error) error {
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return err //nolint:wrapcheck // this is a thin wrapper around an external package, for internal use only
	}

	err = body(tx)
	if err != nil {
		err = errors.Join(err, tx.Rollback())
	} else {
		err = errors.Join(err, tx.Commit())
	}

	return err
}

// NewCounter creates a new CounterService.
//
// During initialization, database migrations will be run.
//
// TODO: factor out database concerns.
func NewCounter(ctx context.Context, db *sql.DB) (*CounterService, error) {
	err := withTx(ctx, db, func(tx *sql.Tx) error {
		c := counter.New(tx)

		_, err := c.Get(ctx)
		if errors.Is(err, sql.ErrNoRows) {
			err = c.Reset(ctx)
			if err != nil {
				return err //nolint:wrapcheck // this is a thin wrapper around an external package, for internal use only
			}
		} else if err != nil {
			return err //nolint:wrapcheck // this is a thin wrapper around an external package, for internal use only
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

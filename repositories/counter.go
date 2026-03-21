package repositories

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"git.jarkad.net.eu.org/jarkad/pocket-expo/queries/number"
)

// NumberRepository is a repository that stores a single integer.
//
// Get() retrieves the current value.
// Set() updates the value.
type NumberRepository struct {
	db      *sql.DB
	queries *number.Queries
}

// NewNumberRepository creates a new Counter.
//
// This, among other things, prepares the SQL statements it will need and initializes the counter row.
func NewNumberRepository(ctx context.Context, db *sql.DB) (repo *NumberRepository, err error) {
	defer func() {
		if err != nil {
			err = fmt.Errorf("while initializing NumberRepository: %w", err)
		}
	}()

	queries, err := number.Prepare(ctx, db)
	if err != nil {
		return nil, err //nolint:wrapcheck // defer above handles wrapping
	}

	repo = &NumberRepository{
		db:      db,
		queries: queries,
	}

	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err //nolint:wrapcheck // defer above handles wrapping
	}

	defer func() {
		e := tx.Rollback()
		if !errors.Is(e, sql.ErrTxDone) {
			err = errors.Join(err, e)
		}
	}()

	err = repo.transact(ctx, func(txRepo *NumberRepository) error {
		count, err := txRepo.queries.Get(ctx)
		if err != nil && !errors.Is(err, sql.ErrNoRows) {
			return fmt.Errorf("while initializing counter: %w", err)
		}

		err = txRepo.queries.DeleteAll(ctx)
		if err != nil {
			return fmt.Errorf("while pruning counter: %w", err)
		}

		err = txRepo.queries.Create(ctx, count)
		if err != nil {
			return fmt.Errorf("while re-inserting counter: %w", err)
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	return repo, nil
}

// Get retrieves current counter value.
func (c *NumberRepository) Get(ctx context.Context) (int, error) {
	count, err := c.queries.Get(ctx)
	if err != nil {
		return 0, fmt.Errorf("while getting counter: %w", err)
	}

	return int(count), nil
}

// Set updates the counter.
func (c *NumberRepository) Set(ctx context.Context, value int) error {
	err := c.queries.Set(ctx, int64(value))
	if err != nil {
		return fmt.Errorf("could not update counter: %w", err)
	}

	return nil
}

// Update runs txFunc in a transaction.
func (c *NumberRepository) Update(ctx context.Context, updateFunc func(value int) (int, error)) error {
	return c.transact(ctx, func(txRepo *NumberRepository) error {
		value, err := txRepo.Get(ctx)
		if err != nil {
			return err
		}

		value, err = updateFunc(value)
		if err != nil {
			return err
		}

		err = txRepo.Set(ctx, value)
		if err != nil {
			return err
		}

		return nil
	})
}

func (c *NumberRepository) transact(ctx context.Context, txFunc func(tx *NumberRepository) error) (err error) {
	defer func() {
		if err != nil {
			err = fmt.Errorf("while running a transaction in NumberRepository: %w", err)
		}
	}()

	tx, err := c.db.BeginTx(ctx, nil)
	if err != nil {
		return err //nolint:wrapcheck // defer above handles wrapping
	}

	defer func() {
		e := tx.Rollback()
		if !errors.Is(e, sql.ErrTxDone) {
			err = errors.Join(err, e)
		}
	}()

	txRepo := &NumberRepository{
		db:      c.db,
		queries: c.queries.WithTx(tx),
	}

	err = txFunc(txRepo)
	if err != nil {
		return err
	}

	err = tx.Commit()

	return err //nolint:wrapcheck // defer above handles wrapping
}

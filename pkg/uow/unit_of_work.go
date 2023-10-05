package uow

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
)

type RepositoryFactory func(tx *sql.Tx) interface{}

type Uow struct {
	DB           *sql.DB
	Tx           *sql.Tx
	Repositories map[string]RepositoryFactory
}

type IUowInterface interface {
	Register(name string, fc RepositoryFactory)
	GetRepository(ctx context.Context, name string) (interface{}, error)
	Do(ctx context.Context, fn func(uow *Uow) error) error
	CommitOrRollback() error
	Rollback() error
	UnRegister(name string)
}

func NewUow(ctx context.Context, db *sql.DB) *Uow {
	return &Uow{
		DB:           db,
		Repositories: make(map[string]RepositoryFactory),
	}
}

func (uow *Uow) Register(name string, fc RepositoryFactory) {
	uow.Repositories[name] = fc
}

func (uow *Uow) UnRegister(name string) {
	delete(uow.Repositories, name)
}

func (uow *Uow) GetRepository(ctx context.Context, name string) (interface{}, error) {
	if uow.Tx == nil {
		tx, err := uow.DB.BeginTx(ctx, nil)
		if err != nil {
			return nil, err
		}
		uow.Tx = tx
	}

	repo := uow.Repositories[name](uow.Tx)
	return repo, nil
}

func (uow *Uow) Do(ctx context.Context, fn func(Uow *Uow) error) error {
	if uow.Tx != nil {
		return fmt.Errorf("transaction already started")
	}
	tx, err := uow.DB.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	uow.Tx = tx
	err = fn(uow)
	if err != nil {
		errRb := uow.Rollback()
		if errRb != nil {
			return errors.New(fmt.Sprintf("original error: %s, rollback error: %s", err.Error(), errRb.Error()))
		}
		return err
	}
	return uow.CommitOrRollback()
}

func (uow *Uow) Rollback() error {
	if uow.Tx == nil {
		return errors.New("no transaction to rollback")
	}
	err := uow.Tx.Rollback()
	if err != nil {
		return err
	}
	uow.Tx = nil
	return nil
}

func (uow *Uow) CommitOrRollback() error {
	err := uow.Tx.Commit()
	if err != nil {
		errRb := uow.Rollback()
		if errRb != nil {
			return errors.New(fmt.Sprintf("original error: %s, rollback error: %s", err.Error(), errRb.Error()))
		}
		return err
	}
	uow.Tx = nil
	return nil
}

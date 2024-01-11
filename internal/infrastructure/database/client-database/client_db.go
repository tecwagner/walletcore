package clientDatabase

import (
	"context"
	"database/sql"
	"errors"

	"github.com/tecwagner/walletcore-service/internal/entity"
	"github.com/tecwagner/walletcore-service/internal/telemetry"
	"go.opentelemetry.io/otel/codes"
)

type ClientDB struct {
	DB        *sql.DB
	telemetry telemetry.Telemetry
}

func NewClientDB(db *sql.DB, telemetry telemetry.Telemetry) *ClientDB {
	return &ClientDB{
		DB:        db,
		telemetry: telemetry,
	}
}

func (cli *ClientDB) Get(id string) (*entity.Client, error) {
	ctx := context.Background()
	ctx, span := cli.telemetry.Start(ctx, "mysql-FindByClient")
	defer span.End()

	client := &entity.Client{}
	stmt, err := cli.DB.Prepare("SELECT id, name, email, created_at FROM clients WHERE id = ?")

	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
		return nil, err
	}
	defer stmt.Close()

	row := stmt.QueryRow(id)
	if err := row.Scan(&client.ID, &client.Name, &client.Email, &client.CreatedAt); err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
		return nil, err
	}

	return client, nil
}

func (cli *ClientDB) FindByClient(email string) (*entity.Client, error) {

	ctx := context.Background()
	ctx, span := cli.telemetry.Start(ctx, "mysql-FindByClient")
	defer span.End()

	client := &entity.Client{}
	stmt, err := cli.DB.Prepare("SELECT id, email, password FROM clients WHERE email = ? ")
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
		return nil, err
	}
	defer stmt.Close()

	var passwordHash string
	row := stmt.QueryRow(email)
	if err := row.Scan(&client.ID, &client.Email, &passwordHash); err != nil {

		if err == sql.ErrNoRows {
			span.RecordError(err)
			span.SetStatus(codes.Error, err.Error())
			return nil, errors.New("customer not found")
		}
		return nil, err
	}
	client.Password = passwordHash
	return client, nil
}

func (cli *ClientDB) Save(client *entity.Client) error {

	ctx := context.Background()
	ctx, span := cli.telemetry.Start(ctx, "mysql-FindByClient")
	defer span.End()

	stmt, err := cli.DB.Prepare("INSERT INTO clients ( id, name, email, password, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?)")
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(client.ID, client.Name, client.Email, client.Password, client.CreatedAt, client.UpdatedAt)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
		return err
	}
	return nil
}

func (cli *ClientDB) IsEmailExists(email string) bool {

	ctx := context.Background()
	ctx, span := cli.telemetry.Start(ctx, "mysql-FindByClient")
	defer span.End()

	client := &entity.Client{}
	stmt, err := cli.DB.Prepare("SELECT id FROM clients WHERE email = ?")
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
		return false
	}
	defer stmt.Close()

	row := stmt.QueryRow(email)
	if err := row.Scan(&client.ID); err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
		return false
	}

	return true
}

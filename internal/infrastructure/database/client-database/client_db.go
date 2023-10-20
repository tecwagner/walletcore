package clientDatabase

import (
	"database/sql"

	"github.com/tecwagner/walletcore-service/internal/entity"
)

type ClientDB struct {
	DB *sql.DB
}

func NewClientDB(db *sql.DB) *ClientDB {
	return &ClientDB{
		DB: db,
	}
}

func (cli *ClientDB) Get(id string) (*entity.Client, error) {
	client := &entity.Client{}
	stmt, err := cli.DB.Prepare("SELECT id, name, email, created_at FROM clients WHERE id = ?")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	row := stmt.QueryRow(id)
	if err := row.Scan(&client.ID, &client.Name, &client.Email, &client.CreatedAt); err != nil {
		return nil, err
	}

	return client, nil
}

func (cli *ClientDB) Save(client *entity.Client) error {

	stmt, err := cli.DB.Prepare("INSERT INTO clients ( id, name, email, created_at, updated_at) VALUES (?, ?, ?, ?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(client.ID, client.Name, client.Email, client.CreatedAt, client.UpdatedAt)
	if err != nil {
		return err
	}
	return nil
}

func (cli *ClientDB) IsEmailExists(email string) bool {
	client := &entity.Client{}
	stmt, err := cli.DB.Prepare("SELECT id FROM clients WHERE email = ?")
	if err != nil {
		return false
	}
	defer stmt.Close()

	row := stmt.QueryRow(email)
	if err := row.Scan(&client.ID); err != nil {
		return false
	}

	return true
}

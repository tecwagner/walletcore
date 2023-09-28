package accountDatabase

import (
	"database/sql"

	"github.com/tecwagner/walletcore-service/internal/entity"
)

type AccountDB struct {
	DB *sql.DB
}

func NewAccountDB(db *sql.DB) *AccountDB {
	return &AccountDB{
		DB: db,
	}
}


func (acc *AccountDB) FindByID(id string) (*entity.Account, error) {
	var account entity.Account
	var client entity.Client
	account.Client = &client

	stmt, err := acc.DB.Prepare("SELECT acc.id, acc.client_id, acc.balance, acc.created_at, cli.id, cli.name, cli.email, cli.created_at FROM accounts acc INNER JOIN clients cli ON acc.client_id = cli.id WHERE acc.id = ?")

	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	row := stmt.QueryRow(id)
	err = row.Scan(
		&account.ID,
		&account.Client.ID,
		&account.Balance,
		&account.CreatedAt,
		&client.ID,
		&client.Name,
		&client.Email,
		&client.CreatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &account, nil
}

func (acc *AccountDB) Save(account *entity.Account) error {
	stmt, err := acc.DB.Prepare("INSERT INTO accounts ( id, client_id,  balance, created_at) VALUES (?, ?, ?, ?)")
	if err != nil { 
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(account.ID, account.Client.ID, account.Balance, account.CreatedAt)
	if err != nil {
		return err
	}

	return nil
}
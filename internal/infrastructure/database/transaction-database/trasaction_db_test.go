package transactionDatabase

import (
	"database/sql"
	"testing"

	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/suite"
	"github.com/tecwagner/walletcore-service/internal/entity"
)

type TransactionDBTestSuite struct {
	suite.Suite
	db            *sql.DB
	payer         *entity.Client
	accountFrom   *entity.Account
	payee         *entity.Client
	accountTo     *entity.Account
	transactionDB *TransactionDB
}

func (suite *TransactionDBTestSuite) SetupSuite() {
	db, err := sql.Open("sqlite3", ":memory:")
	suite.Nil(err)
	suite.db = db
	db.Exec("CREATE TABLE clients (id varchar(255) PRIMARY KEY NOT NULL, name varchar(255) NOT NULL, email varchar(255) NOT NULL, password varchar(50), created_at date, updated_at date)")
	db.Exec("CREATE TABLE accounts (id varchar(255) PRIMARY KEY NOT NULL, client_id varchar(255) NOT NULL, balance int, created_at date, updated_at date)")
	db.Exec("CREATE TABLE transactions (id varchar(255) PRIMARY KEY NOT NULL, account_id_from varchar(255) NOT NULL, account_id_to varchar(255) NOT NULL, amount int, created_at date)")
	client, err := entity.NewClient("John", "j@j.com", "123")
	suite.Nil(err)
	suite.payer = client
	payee, err := entity.NewClient("John2", "jj@j.com", "123")
	suite.Nil(err)
	suite.payee = payee
	//creating accounts
	accountFrom := entity.NewAccount(suite.payer)
	accountFrom.Balance = 1000
	suite.accountFrom = accountFrom
	accountTo := entity.NewAccount(suite.payee)
	accountTo.Balance = 1000
	suite.accountTo = accountTo
	suite.transactionDB = NewTransactionDB(db)
}

func (s *TransactionDBTestSuite) TearDownSuite() {
	defer s.db.Close()
	s.db.Exec("DROP TABLE clients")
	s.db.Exec("DROP TABLE accounts")
	s.db.Exec("DROP TABLE transactions")
}

func TestTransactionDBTestSuite(t *testing.T) {
	suite.Run(t, new(TransactionDBTestSuite))
}

func (s *TransactionDBTestSuite) TestCreate() {
	transaction, err := entity.NewTransaction(s.accountFrom, s.accountTo, 100)
	s.Nil(err)
	err = s.transactionDB.Create(transaction)
	s.Nil(err)
}

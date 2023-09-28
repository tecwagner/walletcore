package accountDatabase

import (
	"database/sql"
	"testing"

	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/suite"
	"github.com/tecwagner/walletcore-service/internal/entity"
)

type AccountDBTestSuite struct {
	suite.Suite
	db *sql.DB
	accountDB *AccountDB
	client *entity.Client
}

func (setup *AccountDBTestSuite) SetupSuite() {
	db, err := sql.Open("sqlite3", ":memory:")
	setup.Nil(err)
	setup.db = db
	db.Exec("CREATE TABLE clients (id varchar(255), name varchar(255), email varchar(255), created_at date)")
	db.Exec("CREATE TABLE accounts (id varchar(255), client_id varchar(255), balance int, created_at date)")
	setup.accountDB = NewAccountDB(db)
	setup.client, _ = entity.NewClient("joh", "joh@example.com")
}

func (setup *AccountDBTestSuite) TearDownSuite() {
	defer setup.db.Close()
	setup.db.Exec("DROP TABLE clients")
	setup.db.Exec("DROP TABLE accounts")
}

func TestClientDBTestSuite(t *testing.T) {
	suite.Run(t, new(AccountDBTestSuite))
}

func (setup *AccountDBTestSuite) TestSave() {
	account := entity.NewAccount(setup.client)
	err := setup.accountDB.Save(account)
	setup.Nil(err)
}

func (setup *AccountDBTestSuite) TestFindByID() {
	setup.db.Exec("INSERT INTO clients (id, name, email, created_at) VALUES (?, ?, ?, ?)", setup.client.ID, setup.client.Name, setup.client.Email, setup.client.CreatedAt,
	)	

	account := entity.NewAccount(setup.client)
	err := setup.accountDB.Save(account)
	setup.Nil(err)
	accountDB, err := setup.accountDB.FindByID(account.ID)
	setup.Nil(err)
	setup.Equal(account.ID, accountDB.ID)
	setup.Equal(account.ClientID, accountDB.ClientID)
	setup.Equal(account.Balance, accountDB.Balance)
	setup.Equal(account.Client.ID, accountDB.Client.ID)
	setup.Equal(account.Client.Name, accountDB.Client.Name)
	setup.Equal(account.Client.Email, accountDB.Client.Email)
} 

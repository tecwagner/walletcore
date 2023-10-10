package clientDatabase

import (
	"database/sql"
	"testing"

	uuid "github.com/google/uuid"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/suite"
	"github.com/tecwagner/walletcore-service/internal/entity"
)

type ClientDBTestSuite struct {
	suite.Suite
	db *sql.DB
	clientDB *ClientDB
}

func (setup *ClientDBTestSuite) SetupSuite() {
	db, err := sql.Open("sqlite3", ":memory:")
	setup.Nil(err)
	setup.db = db
	db.Exec("CREATE TABLE clients (id varchar(255) PRIMARY KEY, name varchar(255), email varchar(255), created_at date, updated_at date)")
	setup.clientDB = NewClientDB(db)	
}

func (setup *ClientDBTestSuite) TearDownSuite() {
	defer setup.db.Close()
	setup.db.Exec("DROP TABLE clients")
}

func TestClientDBTestSuite(t *testing.T) {
	suite.Run(t, new(ClientDBTestSuite))
}

func (suite *ClientDBTestSuite) TestSave() {
	client := &entity.Client{ID: uuid.NewString() , Name: "joh", Email: "joh@example.com"}
	err := suite.clientDB.Save(client)
	suite.Nil(err)
}

func (suite *ClientDBTestSuite) TestGet() {
	client, _ := entity.NewClient("joh", "joh@example.com")
	suite.clientDB.Save(client)

	clientDB, err := suite.clientDB.Get(client.ID)
	suite.Nil(err)
	suite.Equal(client.ID, clientDB.ID)
	suite.Equal(client.Name, clientDB.Name)
	suite.Equal(client.Email, clientDB.Email)
}

func (suite *ClientDBTestSuite) TestSaveWithDuplicateEmail() {
	
	client := &entity.Client{ID: uuid.NewString(), Name: "joh", Email: "joh@example.com"}
	err := suite.clientDB.Save(client)
	suite.Nil(err)

	duplicateClient := &entity.Client{ID: uuid.NewString(), Name: "jane", Email: "joh@example.com"}
	result := suite.clientDB.IsEmailExists(duplicateClient.Email)

	suite.True(result)
}


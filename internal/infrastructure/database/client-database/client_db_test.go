package clientDatabase

import (
	"database/sql"
	"testing"

	uuid "github.com/google/uuid"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/suite"
	"github.com/tecwagner/walletcore-service/internal/entity"
	"github.com/tecwagner/walletcore-service/internal/telemetry"
)

type ClientDBTestSuite struct {
	suite.Suite
	db       *sql.DB
	clientDB *ClientDB
}

func (setup *ClientDBTestSuite) SetupSuite() {
	db, err := sql.Open("sqlite3", ":memory:")
	setup.Nil(err)
	setup.db = db
	db.Exec("CREATE TABLE clients (id varchar(255) PRIMARY KEY, name varchar(255), email varchar(255), password varchar(150), created_at date, updated_at date)")
	setup.clientDB = NewClientDB(db, &telemetry.OTel{})
}

func (setup *ClientDBTestSuite) TearDownSuite() {
	defer setup.db.Close()
	setup.db.Exec("DROP TABLE clients")
}

func TestClientDBTestSuite(t *testing.T) {
	suite.Run(t, new(ClientDBTestSuite))
}

func (suite *ClientDBTestSuite) TestSave() {
	client := &entity.Client{ID: uuid.NewString(), Name: "joh", Email: "joh@example.com", Password: "123"}
	err := suite.clientDB.Save(client)
	suite.Nil(err)
}

func (suite *ClientDBTestSuite) TestGet() {
	client, _ := entity.NewClient("joh", "joh@example.com", "123")
	suite.clientDB.Save(client)

	clientDB, err := suite.clientDB.Get(client.ID)
	suite.Nil(err)
	suite.Equal(client.ID, clientDB.ID)
	suite.Equal(client.Name, clientDB.Name)
	suite.Equal(client.Email, clientDB.Email)
}

func (suite *ClientDBTestSuite) TestSaveWithDuplicateEmail() {

	client := &entity.Client{ID: uuid.NewString(), Name: "joh", Email: "joh@example.com", Password: "123"}
	err := suite.clientDB.Save(client)
	suite.Nil(err)

	duplicateClient := &entity.Client{ID: uuid.NewString(), Name: "jane", Email: "joh@example.com", Password: "123"}
	result := suite.clientDB.IsEmailExists(duplicateClient.Email)

	suite.True(result)
}
func (suite *ClientDBTestSuite) TestVerifyUserExist() {

	client := &entity.Client{ID: uuid.NewString(), Name: "joh", Email: "joh@example.com", Password: "123"}
	err := suite.clientDB.Save(client)
	suite.Nil(err)

	result, err := suite.clientDB.FindByClient(client.Email)

	suite.NoError(err)
	suite.NotNil(result)
}

func (suite *ClientDBTestSuite) TestVerifyUserNotExist() {

	client := &entity.Client{ID: uuid.NewString(), Name: "joh", Email: "joh@example.com", Password: "123"}
	err := suite.clientDB.Save(client)
	suite.Nil(err)

	clientNotExit := &entity.Client{ID: uuid.NewString(), Name: "joh teste", Email: "johteste@example.com", Password: "1235"}

	result, err := suite.clientDB.FindByClient(clientNotExit.Email)

	suite.NotNil(err)
	suite.Nil(result)

	suite.Equal("customer not found", err.Error())
}

package main

import (
	"context"
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/tecwagner/walletcore-service/internal/event"
	accountDatabase "github.com/tecwagner/walletcore-service/internal/infrastructure/database/account-database"
	clientDatabase "github.com/tecwagner/walletcore-service/internal/infrastructure/database/client-database"
	transactionDatabase "github.com/tecwagner/walletcore-service/internal/infrastructure/database/transaction-database"
	createAccount "github.com/tecwagner/walletcore-service/internal/useCase/create_account"
	createClient "github.com/tecwagner/walletcore-service/internal/useCase/create_client"
	createTransaction "github.com/tecwagner/walletcore-service/internal/useCase/create_transaction"
	"github.com/tecwagner/walletcore-service/internal/web"
	"github.com/tecwagner/walletcore-service/internal/web/webserver"
	"github.com/tecwagner/walletcore-service/pkg/events"
	"github.com/tecwagner/walletcore-service/pkg/uow"
)

func main() {
	//Criar config para para variavel de ambiente
	db := setupDatabase()
	defer db.Close()

	// Iniciando as dependencias para disparos de email
	eventDispatcher := setupEventDispatcher()
	createTransactionEvent := event.NewTransactionCreated()
	balanceUpdatedEvent := event.NewBalanceUpdated()
	// eventDispatcher.Register("TransactionCreated", handler)

	clientDB := clientDatabase.NewClientDB(db)
	accountDB := accountDatabase.NewAccountDB(db)

	ctx := context.Background()
	uow := uow.NewUow(ctx, db)

	uow.Register("AccountDB", func(tx *sql.Tx) interface{} {
		return accountDatabase.NewAccountDB(db)
	})

	uow.Register("TransactionDB", func(tx *sql.Tx) interface{} {
		return transactionDatabase.NewTransactionDB(db)
	})

	createClientUseCase := createClient.NewCreateClientUseCase(clientDB)
	createAccountUseCase := createAccount.NewCreateAccountUseCase(accountDB, clientDB)
	createTransactionUseCase := createTransaction.NewCreateTransactionUseCase(uow, eventDispatcher, createTransactionEvent, balanceUpdatedEvent)

	// Porta da aplicação
	webserver := webserver.NewWebServer(":8080")

	// Mapeando as rotas
	clientHandler := web.NewWebClientHandler(*createClientUseCase)
	accountHandler := web.NewWebAccountHandler(*createAccountUseCase)
	transactionHandler := web.NewWebTransactionHandler(*createTransactionUseCase)

	webserver.AddHandler("/clients", clientHandler.CreateClient)
	webserver.AddHandler("/accounts", accountHandler.CreateAccount)
	webserver.AddHandler("/transactions", transactionHandler.CreateTransaction)

	fmt.Println("Server is running", webserver.WebServerPort)

	webserver.Start()
}

func setupDatabase() *sql.DB {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", "root", "root", "mysql", "3306", "wallet"))
	if err != nil {
		panic(err)
	}
	return db
}

func setupEventDispatcher() *events.EventDispatcher {
	return events.NewEventDispatcher()
}
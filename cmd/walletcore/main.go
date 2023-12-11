package main

import (
	"context"
	"database/sql"
	"fmt"

	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
	_ "github.com/go-sql-driver/mysql"
	"github.com/tecwagner/walletcore-service/internal/event"
	"github.com/tecwagner/walletcore-service/internal/event/handler"
	accountDatabase "github.com/tecwagner/walletcore-service/internal/infrastructure/database/account-database"
	clientDatabase "github.com/tecwagner/walletcore-service/internal/infrastructure/database/client-database"
	transactionDatabase "github.com/tecwagner/walletcore-service/internal/infrastructure/database/transaction-database"
	authenticationUser "github.com/tecwagner/walletcore-service/internal/useCase/authentication_user"
	createAccount "github.com/tecwagner/walletcore-service/internal/useCase/create_account"
	createClient "github.com/tecwagner/walletcore-service/internal/useCase/create_client"
	createTransaction "github.com/tecwagner/walletcore-service/internal/useCase/create_transaction"
	"github.com/tecwagner/walletcore-service/internal/web"
	"github.com/tecwagner/walletcore-service/internal/web/webserver"
	"github.com/tecwagner/walletcore-service/pkg/events"
	"github.com/tecwagner/walletcore-service/pkg/kafka"
	"github.com/tecwagner/walletcore-service/pkg/uow"
)

func main() {

	//Criar config para para variavel de ambiente
	db := setupDatabase()
	defer db.Close()

	// Mapeando o configMap Kafka
	configMap := ckafka.ConfigMap{
		"bootstrap.servers": "kafka:29092",
		"group.id":          "wallet",
	}
	kafkaProducer := kafka.NewProducer(&configMap)

	// Iniciando as dependencias de registro de event dispatched
	eventDispatcher := setupEventDispatcher()
	createTransactionEvent := event.NewTransactionCreated()
	balanceUpdatedEvent := event.NewBalanceUpdated()

	// Registrando o evento handler Transaction Criado no Kafka Producer
	eventDispatcher.Register("TransactionCreated", handler.NewTransactionCreatedKafkaHandler(kafkaProducer))
	eventDispatcher.Register("BalanceUpdated", handler.NewUpdatedBalanceKafkaHandler(kafkaProducer))

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

	authenticationUseCase := authenticationUser.NewAuthenticationUseCase(clientDB)
	createClientUseCase := createClient.NewCreateClientUseCase(clientDB)
	createAccountUseCase := createAccount.NewCreateAccountUseCase(accountDB, clientDB)
	createTransactionUseCase := createTransaction.NewCreateTransactionUseCase(uow, eventDispatcher, createTransactionEvent, balanceUpdatedEvent)

	// Porta da aplicação
	webserver := webserver.NewWebServer(":8081")

	// Mapeando as rotas
	authenticationHandler := web.NewWebAuthenticationHandler(*authenticationUseCase)
	clientHandler := web.NewWebClientHandler(*createClientUseCase)
	accountHandler := web.NewWebAccountHandler(*createAccountUseCase)
	transactionHandler := web.NewWebTransactionHandler(*createTransactionUseCase)

	webserver.AddHandlerPublic("/api/v1/login", authenticationHandler.AuthUser, true)
	webserver.AddHandlerPublic("/api/v1/clients", clientHandler.CreateClient, true)
	webserver.AddHandlerPublic("/api/v1/accounts", accountHandler.CreateAccount, false)
	webserver.AddHandlerPublic("/api/v1/transactions", transactionHandler.CreateTransaction, false)

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

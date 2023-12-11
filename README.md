## WalletCore - Sistema de Gerenciamento de Carteira

# Tecnologias

    - Linguagem: Go
    - Docker e Docker-Compose
    - Github para versionamento
    - Biblioteca Testify/Mock para testes de unidade e suites de testes
    - SQLite3 para testes de integração
    - Banco de Dados MySQL
    - Apache Kafka (Producer e Consumer)
    - JWT

# Iniciando o container docker

    -  docker compose up -d

# Executando o projeto com docker e acessando o bash da aplicação

    - docker compose exec goapp bash

# Executando o projeto dentro do container

    - root@56965583d20d:/app/cmd/walletcore# go run main.go

# Utilizando o docker para acessar o banco de dados

    - docker compose exec mysql bash
        - bash-4.2# mysql -uroot -p wallet
        - Enter password: root

# Criando tabelas na base de dados

    - clients
        - CREATE TABLE clients (id varchar(255) PRIMARY KEY NOT NULL, name varchar(255) NOT NULL, email varchar(255) NOT NULL, password varchar(150), created_at date, updated_at date);

    - accounts
        - CREATE TABLE accounts (id varchar(255) PRIMARY KEY NOT NULL, client_id varchar(255) NOT NULL, balance int, created_at date, updated_at date);

    - transaction
        - CREATE TABLE transactions (id varchar(255) PRIMARY KEY NOT NULL, account_id_from varchar(255) NOT NULL, account_id_to varchar(255) NOT NULL, amount int, created_at date);

# Implementando o Kafka ao Projeto para Produzir e Consumir Mensagens

    - pkg/kafka
        - Kafka Producer
        - Kafka Consumer

# Disparando Eventos para o Kafka

    - O Docker Control-Center é para registrar os eventos
    - Acesse o Control-Center http://localhost:9021/clusters
        - Crie os tópicos
            - como transactions e balances, para que outros micro serviços possam consumir esses dados e registrar transações realizadas.

# Para realizar as transações utilize os andpoints

    - api/client.http

# Teste de Strees com Siege

    - strees test sucesso
        - siege --header="Content-Type: application/json" -c 100 -r 100 -d 30s "http://localhost:8081/api/v1/login POST {\"email\": \"wagner.oliveira@gmail.com\", \"password\": \"1234\"}"

    - strees test error
        - siege --header="Content-Type: application/json" -c 100 -r 100 -d 30s "http://localhost:8081/api/v1/login POST {\"email\": \"wagner.oliveira@gmail.com\", \"password\": \"12\"}"

    - siege -c 10 -r 1000 "http://localhost:8081/api/v1/login

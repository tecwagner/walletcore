## WalletCore - Sistema de Gerenciamento de Carteira

# Tecnologias

    - Linguagem: Go
    - Docker e Docker-Compose
    - Github para versionamento
    - Biblioteca Testify/Mock para testes de unidade e suites de testes
    - SQLite3 para testes de integração
    - Banco de Dados MySQL
    - Apache Kafka (Producer e Consumer)

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

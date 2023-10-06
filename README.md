## walletcore

# Tecnologias

    - Go Language
    - Docker
    - Docker-Compose
    - Github
    - Moks Utiliando biblioteca testify/mock"
    - Suite de Test Utiliando biblioteca testify/suite"
    - sqlite3 para teste de integração
    - Mysql Banco de Dados
    - Kafka (Producer e Consumer)

# Criando o projeto atraves do comando:

    - go mod init github.com/tecwagner/walletcore-service

# Camada de Entidade com as Regras de Negocio

    - internal
        - entity
            - account
            - client
            - transaction

# Camada de Gateway

    - Que será uma inteface da implementação do Repository

    - internal
        - gateway
            - account_gateway
                - account-interface.go
            - client_gateway
                - client-interface.go
            - transaction_gateway
                - transaction-interface.go

# Camada de UseCase

    - Casos de Uso da aplicação

    - internal
        - useCase
            - create_client
                - dto
                - interface-gateway
                - test de UseCase
                - criando novos clients com metodo Execute
             - create_account
                - dto
                - interface-gateway
                - test de UseCase
                - criando novos account com metodo Execute
             - create_transaction
                - dto
                - interface-gateway
                - test de UseCase
                - criando novos transaction com metodo Execute

# Camada externa de comunicação

    - internal
        - infrastructure
            - database
                - account-database
                - client-database
                - transaction-database

# Package de comunicação por Eventos

    - pkg
        - events

            - IEventInterface: Metodos de eventos que irá executar.
            - IEventHandlerInterface: Tem apenas um metodo obrigatorio. Que é responsavel por executar os eventos.
            - IEventDispatcherInterface: Gerenciador dos eventos com os metodos a serem implementados: Registra, Dispacha, Remove e Limpa

        - Dispatcher Test unitario
            - event_dispatcher
            - event_dispatcher_test

# Implementando o modelo eventos para criação de Transação implementando a interface de events

    - Sempre que for trabalhar com eventos, crie os seus eventos de caso de uso utilizando da interface criada para enventos

    - internal
        - event
            - transaction_create

    - useCase
        - create_transaction
            - create_transaction_dto: Foi adicionado os metodos que será implementado.
            - create_transaction: Instanciado os novos metodos e implementado o dispatcher de criação

# Implementando o modulo de container docker criando um banco de dados

    - Criando um banco de dados MySQL
        - Dockerfile
        - docker-compose.yaml

# Implementando o modulo de webserver que será o coração da aplicação

    - cmd
        - wallercore
            - main.go

                - Instanciado a conexão com banco de dados

                - Instanciado o web e webserver conroller
                    - Que mapeia as rotas da aplicação

                - Instanciado o serviço que comunica kafka
                - Instanciado os event dispatched para execução dos eventos e regitro no kafka

# Criando o Modulo de Controller Handler

    - Ele será o controller para executar os usecase e o handler

    - internal
        - web
            - client_handler
            - account_handler
            - transaction_handler

# Implementando webserver utilizando a biblioteca chi

    - Tem como objetivo criar as rotas da aplicação
    - Documentação: https://github.com/go-chi/chi

    - internal
        - web
            - webserver
                webserver.go

# Criando os primeiros registros

    - Acessar o banco de dados que está no docker.
        - docker-compose exec mysql bash
            - mysql -uroot -p wallet
            - connction: root
        - Cria Tabelas:
            - Utilizando Migrations ou scripts
                - internal
                    - gateway
                        - CREATE TABLE clients (id varchar(255) PRIMARY KEY, name varchar(255), email varchar(255), created_at date, updated_at date);
                        - CREATE TABLE accounts (id varchar(255) PRIMARY KEY, client_id varchar(255), balance int, created_at date, updated_at date);
                        - CREATE TABLE transactions (id varchar(255) PRIMARY KEY, account_id_from varchar(255), account_id_to varchar(255), amount int, created_at date);

# Implementando Unit Of Work

    - Por exemplo, em um contexto de banco de dados, uma "unit of work" geralmente se refere a uma transação de banco de dados. Uma transação é uma sequência de operações de banco de dados que devem ser executadas de forma atômica, ou seja, todas as operações devem ser concluídas com sucesso ou nenhuma delas deve ser aplicada. Isso ajuda a manter a consistência dos dados no banco de dados.

    - pkg
        - uow
            - Foi implementado os metodos de Unit Of Work.

# Implementando o Kafka ao Projeto para produzir e consumir as mensagens

        - pkg
            - kafka
                - Criando o Kafka producer
                - Criando o Kafka consumer

# Criando TransactionHandler

    - Implementado o handle para fins de registro que serão enviados para o kafka e registrado na base de datos.

    - internal
        - event
            - handler
                - transaction_created_kafka
                - balance_uptaded_kafka
                    - Criando os metodo Handle para o Producer das mensagens para o kafka via envento handler

# Registrando handler

    - No main.go
        - Registra o kafka mapeando a configuração do cliente do kafka

        - Registrando o envio dos enventos ao kafka utilizados da interface implementada de event handler e dispatched

# Disparando o evento para o kafka

    - Criando um container no doker do control-center para registrar os eventos.
        - Criando um Topic: transctions.
        - Criando um Topic: balances.
            - Para que os demais micro serviços consome esses dados.

            - Para registro e controle das transações realizadas.

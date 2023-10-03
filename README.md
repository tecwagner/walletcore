## walletcore

# Tecnologias

    - Go Language
    - Docker
    - Docker-Compose
    - Github
    - Moks Utiliando biblioteca testify/mock"
    - Suite de Test Utiliando biblioteca testify/suite"
    - sqlite3 para teste de integração

# Criando o projeto atarves do comando:

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

            - Interfaces

                - event_interface

                    - IEventInterface: Metodos de eventos que irá executar.

                    - IEventHandlerInterface: Tem apenas um metodo obrigatorio. Que é responsavel por executar os eventos.

                    - IEventDispacherInterface: Gerenciado dos eventos Registra, Dispacha, Remove e Limpa

            - Dispacher

                - event_dispatcher
                - event_dispatcher_test

# Implementando eventos create para as Trasaction

    - Sempre que for trabalhar com eventos, crie os seus eventos de caso de uso

    - internal
        - event
            - transaction_create

    - useCase
        - create_transaction
            - create_transaction_dto: Foi adicionado os metodos que será implementado.
            - create_transaction: Instanciado os novos metodos e implementado o dispatcher de criação

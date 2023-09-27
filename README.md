## walletcore

# Technologies

    - Go Language
    - Docker
    - Docker-Compose
    - Github
    - Moks Utiliando biblioteca testify/mock"

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

## walletcore

# Technologies

    - Go Language
    - Docker
    - Docker-Compose
    - Github

# Criando o projeto atarves do comando:

    - go mod init github.com/tecwagner/walletcore-service

# Criado as Entidade com as Regras de Negocio

    - internal
        - entity
            - account
            - client
            - transaction

# Criando os Gateway

    - Que será uma inteface da implementação do Repository

    - internal
        - gateway
            - client.go

# Criando os UseCase

    - Casos de Uso da aplicação

    - internal
        - useCase
            - create_client
                - dto
                - interface-gateway
                - test de createClient UseCase
                - criando novos clients com metodo Execute

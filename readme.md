# Full Cycle - APIs e Mensageria com Golang

Projeto desenvolvido durante o vídeo "APIs e Mensageria com Golang" do canal "Full Cycle" em 01/12/2023 ([Assista ao Vídeo](https://www.youtube.com/live/oyk7PFRufQY?si=0yZ3HTd0gGK3Wkd-))

## Sobre o Projeto

Este projeto em Golang apresenta uma aplicação que demonstra os seguintes conceitos:
- Cadastro e listagem de produtos;
- Utilização do pacote `net/http` e `go-chi/chi` para criar endpoints HTTP;
- Integração com MySQL através de `database/sql`;
- Processamento de mensagens com Kafka;
- Demonstração do uso de Clean Architecture.

A aplicação oferece os seguintes endpoints:
- `POST /product` para cadastrar produtos no banco de dados;
- `GET /products` para listar os produtos armazenados no banco de dados.

Simultaneamente, há o processamento de mensagens do Kafka, onde cada mensagem proveniente do tópico `product` é processada e salva no banco de dados.

## Executando o Projeto

### Requisitos:
- Docker e docker-compose instalados.

### Passo a Passo:
1. Execute `docker-compose up -d` para iniciar os containers;
2. Crie a tabela `products` no banco de dados `products` usando o comando:

    ```sql
    create table products (id varchar(255), name varchar(255), price float);
    ```

3. Acesse o container do Kafka: `docker-compose exec kafka bash`;
4. Crie o tópico: `kafka-topics --bootstrap-server=localhost:9092 --topic=product --create`;
5. Acesse o container da aplicação Go: `docker-compose exec goapp bash`;
6. Execute a aplicação: `go run cmd/app/main.go`;
7. Aguarde a inicialização. : )

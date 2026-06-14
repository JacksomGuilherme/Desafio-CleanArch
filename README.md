# Clean Architecture - Listagem de Orders (REST, gRPC e GraphQL)

## Descrição

Este projeto foi desenvolvido como parte do desafio **Clean Architecture: Listagem de Orders**, utilizando os princípios da Clean Architecture para demonstrar o desacoplamento entre regras de negócio e interfaces de entrada.

A aplicação possui dois casos de uso:

* **CreateOrderUseCase**: responsável pela criação de pedidos.
* **ListOrdersUseCase**: responsável pela listagem de pedidos.

Esses casos de uso são expostos simultaneamente através de três interfaces de comunicação diferentes:

* REST
* gRPC
* GraphQL

Além disso, a aplicação realiza automaticamente:

* Inicialização do banco de dados via Docker;
* Aplicação das migrations na inicialização;
* Inicialização dos serviços REST, gRPC e GraphQL;
* Espera pelo banco de dados antes de iniciar a aplicação, evitando problemas de race condition.

---

## Tecnologias Utilizadas

* Go (Golang)
* Clean Architecture
* MySQL
* RabbitMQ
* REST
* gRPC
* GraphQL
* Docker
* Docker Compose
* golang-migrate
* Google Wire
* gqlgen

---

## Execução

### Pré-requisitos

* Docker
* Docker Compose

### Comando único de execução

```bash
docker compose up
```

Nenhum outro comando é necessário.

Ao executar o comando acima:

1. O MySQL será iniciado;
2. O RabbitMQ será iniciado;
3. A aplicação aguardará o banco de dados ficar disponível;
4. As migrations serão executadas automaticamente;
5. Os serviços REST, gRPC e GraphQL serão iniciados.

---

## Serviços Disponíveis

### REST API

Porta:

```text
8000
```

#### Criar Order

```http
POST /order
```

Payload:

```json
{
  "id": "123",
  "price": 100.0,
  "tax": 10.0
}
```

#### Listar Orders

```http
GET /order
```

Exemplo:

```bash
curl http://localhost:8000/order
```

---

### GraphQL

Porta:

```text
8080
```

Playground:

```text
http://localhost:8080
```

#### Mutation de Criação

```graphql
mutation createOrder{
  createOrder(input: {id: "GQL-1", Price: 100.50, Tax: 10.00}){
    id,
    Price,
    Tax
  }
}
```

#### Query de Listagem

```graphql
query ListOrders {
  listOrders {
    id,
    Price,
    Tax,
    FinalPrice
  }
}
```

---

### gRPC

Porta:

```text
50051
```

Services disponíveis:

* CreateOrder
* ListOrders

---

## Estrutura da Solução

### Use Cases

* CreateOrderUseCase
* ListOrdersUseCase

### Interfaces de Entrada

#### REST

* POST /order
* GET /order

#### GraphQL

* Mutation createOrder
* Query listOrders

#### gRPC

* CreateOrder
* ListOrders

Todas as interfaces utilizam os mesmos casos de uso da camada de aplicação, demonstrando o desacoplamento proposto pela Clean Architecture.

---

## Banco de Dados

O banco de dados é provisionado automaticamente via Docker.

As migrations estão localizadas em:

```text
internal/infra/database/migrations
```

As migrations são executadas automaticamente durante a inicialização da aplicação.

---

## Arquivo api.http

O projeto contém um arquivo `api.http` na raiz do projeto com exemplos prontos para:

* Criar Orders via REST
* Listar Orders via REST

---

## Portas Utilizadas

| Serviço             | Porta |
| ------------------- | ----- |
| REST                | 8000  |
| GraphQL             | 8080  |
| gRPC                | 50051 |

---

## Autor

Jacksom Guilherme Novack dos Santos

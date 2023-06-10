# Simple Commerce Monorepo

![Coverage](coverage.svg)

## goals

### user

- can create an account or use OAuth, edit account details, login, and logout
- account details are: name, email, password, addresses (pick from an interactive map)
- can view order's status (pending, shipped, delivered, cancelled) and orders history
- can view products, categories, and reviews
- can search and sort products
- can add products to their cart and to wishlist
- can view their cart and check out
- can write review on a product if they've purchased it

### order

- handle concurrent orders
- integrate with a mock third-party payment API
- integrate with a mock third-party shipping API

### inventory

- handle concurrent stock change

### blazingly fast SSR pages and mobile app

- user frontend
- admin frontend

### reliability

- full integration test pipeline
- log aggregation and full text search
- production alert

## technical

### models

- [schemas](/database/readme.md)

### overall

- user service
- category service
- product service
- review service
- warehouse service
- inventory service
- order service
- payment service
- shipping service
- aggregator service
- gateway service
- common service
- database service
- frontend service
- server service

each service will have:

- handle layer
- service layer
- middleware layer
- database layer

### tech stack

- backend in Go, support generic funtional programming and fuzzy testing
- communication via HTTP
- frontend SSR in react
- main database in Postgres
- log database in Mongo
- auth in Paseto
- use Docker Compose
- GitHub Actions CI
- custom log aggregator
- test coverage goal of 100%

### backend

- Gin for handlers
- Viper as config loader
- ZeroLog as logger
- SQLC for generating database layer
- GoMock for generating test mocks

### frontend

- using unstyled components with Base UI and Tailwind

### database

- handcrafted schemas
- GolangMigrate for database migration

## development

- create your own `config.yaml` based on `config.yaml.example`

```bash
docker compose up -d
```

```bash
docker compose logs -f
```

```bash
docker compose down
```

- full project test and coverage

```bash
make test
```

```bash
make cover
```

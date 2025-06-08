## 🛠 Simple and Adaptable CRUD Project Structure
A clean and extensible project layout for building CRUD operations that can be easily integrated into any application.

### 📁 Structure Overview:

`cmd/` – Entry point for the application (currently set up for web, but easily extendable for CLI or other interfaces).

`infrastructure/` – Contains PostgreSQL-specific implementations and integrations.

`migrations/` – Database schema migrations (PostgreSQL); includes Makefile targets for creating and running migrations.

`model/` – Struct definitions that map directly to database tables.

`service/` – Business logic layer; defines interfaces and provides concrete implementations for core services like task management.

## INSTALL 

`go get github.com/lib/pq`


## MIGRATION CLI

https://github.com/golang-migrate/migrate

## Run Docker

`docker-compose up -d`

## Run make file commands

```
make create-migration
make migrate
```

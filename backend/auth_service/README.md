# auth-service

## Implementation

Used clean architecture.
Logging with zap.
Scraping of tracers, metrics and logs via otel.
Database is PostgreSQL.
gRPC server contains interceptors: logging, error, recovery.

## Adapters:

##### Controllers:
- grpc

- http (gin)

##### Repositories:
- PostgreSQL

## Dependencies:

##### Used packages:
- google.golang.org/grpc v1.72.0 (for grpc controller)

- github.com/gin-gonic/gin v1.10.0 (for http controller)

- go.uber.org/zap v1.27.0 (for logging)

- go.uber.org/zap/zaptest (for mock logging in unit tests)

- github.com/joho/godotenv v1.5.1 (for .env loading)

- github.com/caarlos0/env/v7 v7.1.0 (for config .env parsing)

- postgrecon in pkg (for connection to psql)

- github.com/sethvargo/go-password (for password generation)
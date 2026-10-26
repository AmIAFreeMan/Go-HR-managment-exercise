An enterprise-grade, secure **HR Management System** built using **Go**, **gRPC**, **PostgreSQL**, and **Auth0 JWT Authentication**,
following **Clean Architecture** principles.

---

## Tech Stack & Tools
* **Language:** Go (Golang)
* **API Framework:** gRPC & Protocol Buffers (`proto3`)
* **Database:** PostgreSQL (`pgx/v5` connection pooling)
* **Authentication:** Auth0 (RS256 JWT Middleware validation via JWKS)
* **Security & Vulnerability:** `govulncheck`, `golangci-lint`, parameterized queries (SQL injection protection)

---

## Project Structure

hr-system/
├── cmd/
│   └── server/
│       └── main.go           # Application entrypoint & dependency wiring
├── internal/
│   ├── domain/               # Pure core business entities & structs
│   │   └── employee.go
│   ├── repository/           # Database layer (PostgreSQL CRUD with pgxpool)
│   │   ├── db.go
│   │   └── employee_repo.go
│   └── delivery/
│       └── grpc/             # gRPC transport handlers & server interceptors
│           └── employee_service.go
├── pkg/
│   └── config/               # Configuration management & environment loaders
│       └── config.go
├── proto/
│   ├── employee.proto        # Protocol Buffers service contract
│   └── hr/
│       └── v1/               # Compiled Go gRPC stubs (.pb.go)
├── go.mod
└── README.md

# 🚀 go-fiber-boilerplate

A modern, production-ready Go backend boilerplate using Fiber, featuring clean architecture, modular service initialization, and standardized API responses.

---

## ✨ Features

-   **Fiber Web Framework**: Fast, expressive, and modern HTTP routing
-   **Clean Architecture**: Layered separation of concerns
-   **Modular Service Initialization**: All services initialized via `InitServices`
-   **Standardized API Responses**: Unified error and success helpers
-   **Configuration Management**: Environment-based config
-   **Structured Logging**: JSON logging with Zap
-   **Redis Integration**: Idiomatic Redis service
-   **JWT Authentication**: Secure, stateless middleware
-   **Health Checks**: Built-in endpoint
-   **Error Handling**: Unified error/success responses
-   **Docker Support**: Multi-stage builds
-   **Database Ready**: PostgreSQL integration
-   **Testing**: Unit & integration examples
-   **CI/CD**: GitHub Actions workflow
-   **API Documentation**: OpenAPI/Swagger ready

---

## 📁 Project Structure

```
├── cmd/                # Main applications
│   └── server/
│       └── main.go     # Entry point
├── internal/           # Private app code
│   ├── api/
│   │   ├── handlers/   # HTTP handlers
│   │   ├── middleware/ # HTTP middleware
│   │   └── routes/     # Route definitions
│   ├── config/         # Config management
│   ├── database/       # DB layer
│   ├── models/         # Data models
│   ├── services/       # Business logic
│   └── utils/          # Utilities
├── migrations/         # DB migrations
├── docs/               # Documentation
├── scripts/            # Build/deploy scripts
├── tests/              # Integration tests
├── .env.example        # Env template
├── docker-compose.yml
├── Dockerfile
├── Makefile
├── go.mod
└── README.md
```

---

## 🧰 Tech Stack

-   **Go 1.21+**
-   **Fiber**: Fast, expressive web framework
-   **Zap**: Structured logging
-   **Redis**: Caching & atomic ops
-   **go-playground/validator**: Input validation
-   **PostgreSQL**: Primary DB (optional)
-   **Docker**: Containerization
-   **Make**: Build automation

---

## ⚡ Quick Start

### Prerequisites

-   Go 1.21 or higher
-   Docker and Docker Compose (optional)
-   PostgreSQL (optional)

### Installation

1. **Clone the repository**

    ```bash
    git clone https://github.com/md-asharaf/go-fiber-boilerplate.git
    cd go-fiber-boilerplate
    ```

2. **Copy environment variables**

    ```bash
    go mod download
    ```

3. **Run the application**
    ```bash
    make run
    # or
    go run ./cmd/main.go
    ```

### Using Docker

1. **Build and run with Docker Compose**

    ```bash
    docker-compose up --build
    ```

2. **Run with Docker only**
    ```bash
    docker build -t go-backend .
    docker run -p 8000:8000 go-backend
    ```

---

## 📝 Environment Variables

| Variable        | Description                          | Default                                                         |
| --------------- | ------------------------------------ | --------------------------------------------------------------- |
| SERVER_HOST     | Server host                          | localhost                                                       |
| SERVER_PORT     | Server port                          | 8000                                                            |
| ENV             | Application environment              | development                                                     |
| LOG_LEVEL       | Log level (debug, info, warn, error) | info                                                            |
| DATABASE_URL    | Database URI                         | postgres://username:password@host:port/database?sslmode=require |
| REDIS_HOST      | Redis host                           | localhost                                                       |
| REDIS_PORT      | Redis port                           | 6379                                                            |
| REDIS_PASSWORD  | Redis password                       |                                                                 |
| JWT_SECRET      | JWT secret                           | your-super-secret-jwt-key-change-this-in-production             |
| SMTP_HOST       | SMTP server host                     | smtp.gmail.com                                                  |
| SMTP_PORT       | SMTP server port                     | 587                                                             |
| SMTP_USERNAME   | SMTP username                        | your-email@gmail.com                                            |
| SMTP_PASSWORD   | SMTP password                        | your-app-password                                               |
| SMTP_FROM_EMAIL | From email address                   | your-email@gmail.com                                            |

---

## 🔗 API Endpoints

### Health Check

```http
GET /api/v1/health
```

### Example Endpoints

```http
GET /api/v1/items
POST /api/v1/items
GET /api/v1/items/{id}
PUT /api/v1/items/{id}
DELETE /api/v1/items/{id}
```

---

## 🧪 Testing

```bash
# Run all tests
make test

# Run tests with coverage
make test-coverage

# Run integration tests
make test-integration
```

---

## 🛠️ Development

### Available Make Commands

```bash
make build          # Build the application
make run             # Run the application
make test            # Run tests
make test-coverage   # Run tests with coverage
make lint            # Run linter
make fmt             # Format code
make clean           # Clean build artifacts
make docker-build    # Build Docker image
make docker-run      # Run Docker container
```

### Code Generation

```bash
# Generate mocks for testing
make generate-mocks

# Generate API documentation
make generate-docs
```

---

## 📦 Service Initialization & Dependency Injection

All core services (DB, JWT, Auth, User, Encryption, Redis) are initialized in one place using the `InitServices` function in `internal/services/app.go`. This function returns an `AppServices` container, which is passed throughout the application for dependency injection.

**Example:**

```go
import "github.com/yourusername/go-backend-boilerplate/internal/services"

// In main.go
appServices := services.InitServices(cfg, db)
```

---

## 🧩 Modularity & Fiber-First Design

-   All service types and initialization logic are centralized in `internal/services/app.go`.
-   Handlers and middleware use Fiber context and receive only the dependencies they need, improving testability and maintainability.
-   No global variables; all dependencies are injected via the `AppServices` container.
-   All routing, middleware, and handlers use Fiber patterns only.

---

## 🆕 How to Add a New Service

1. Add the new service to the `AppServices` struct in `internal/services/app.go`.
2. Update the `InitServices` function to initialize and include the new service.
3. Inject the new service into Fiber handlers or middleware as needed.

---

## 🐳 Docker & Containerization

### Multi-stage Dockerfile

The project includes a multi-stage Dockerfile that:

-   Uses Go modules for dependency management
-   Creates a minimal final image with just the binary
-   Runs as non-root user for security

### Docker Compose

Includes services for:

-   Go application
-   PostgreSQL database
-   Redis (optional)

---

## 📊 Monitoring & Observability

-   **Health Checks**: `/api/v1/health` endpoint
-   **Structured Logging**: JSON format with correlation IDs
-   **Metrics**: Ready for Prometheus integration
-   **Tracing**: Ready for OpenTelemetry integration

---

## 🔒 Security

-   **CORS**: Configurable Fiber CORS middleware
-   **Rate Limiting**: Built-in rate limiting
-   **Authentication**: JWT middleware (Fiber compatible)
-   **Input Validation**: Centralized request validation utilities
-   **Security Headers**: Standard security headers

---

## 📚 Documentation

-   **API Docs**: OpenAPI/Swagger specification
-   **Code Comments**: Comprehensive code documentation
-   **Architecture**: Clean architecture documentation
-   **Fiber Usage**: All examples and docs use Fiber patterns

---

## 🚀 Deployment

### Manual Deployment

1. Build the binary:

    ```bash
    make build
    ```

2. Deploy the binary to your server

### Docker Deployment

1. Build the image:

    ```bash
    docker build -t your-app .
    ```

2. Run the container:
    ```bash
    docker run -p 8000:8000 your-app
    ```

---

## 🤝 Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

---

## 📄 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

---

## 🙏 Acknowledgments

-   [Fiber](https://github.com/gofiber/fiber) for the web framework
-   [Zap](https://github.com/uber-go/zap) for structured logging
-   [go-redis](https://github.com/redis/go-redis) for Redis client
-   Go community for best practices and patterns

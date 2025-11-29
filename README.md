# Go AI Skeleton

This is a skeleton project for building Go applications, designed to be easily extensible for AI-based functionalities. It provides a solid foundation with a modular structure, configuration management, and a robust HTTP server using `chi`.

## Features

- **Modular Architecture**: Clean separation of concerns with dedicated directories for handlers, models, server logic, and utilities.
- **Fast Routing**: Utilizes [go-chi/chi](https://github.com/go-chi/chi) for lightweight and idiomatic Go HTTP routing.
- **Configuration Management**: YAML-based configuration support located in `configs/`.
- **Middleware**: Pre-configured with essential middleware including RequestID, Logger, Recoverer, and Timeout.
- **Command Pattern**: Extensible command manager in `main.go` for easily adding CLI commands.

## Prerequisites

- Go 1.24.9 or higher

## Getting Started

### Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/sonitx/golang-agent-service.git
   cd go-ai-skeleton
   ```

2. Install dependencies:
   ```bash
   go mod download
   ```

### Configuration

The application is configured via `configs/application.yml`. You can modify server settings like the port and mode here.

```yaml
server:
  mode: dev
  port: 8080
```

### Running the Server

To start the backend server, simply run:

```bash
go run main.go
```

The server will start on the port specified in the configuration (default is `8080`).

You should see output indicating the server is running:
```
🚀 Server running on :8080
```

## API Endpoints

The following basic endpoints are available out of the box:

- `GET /`: Home endpoint.
- `GET /ping`: Health check endpoint.

## Project Structure

- `agents/`: Directory for AI agent implementations.
- `api/`: API definitions.
- `base/`: Base classes or interfaces.
- `configs/`: Configuration files.
- `handlers/`: HTTP request handlers.
- `models/`: Data models.
- `server/`: Server initialization, routing, and middleware.
- `services/`: Business logic services.
- `utils/`: Utility functions and helpers.
- `main.go`: Application entry point.

## License

[MIT](LICENSE)

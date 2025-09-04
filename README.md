# MovSeek Server

The MovSeek Server is the backend service powering the MovSeek application. Built with Golang, following a clean project structure (cmd/, internal/, pkg/) for maintainability and scalability. The server exposes APIs for movie search, filtering, pagination, and user-related operations, with HTTPS support out of the box.

## ⚙️ Prerequisites

- Go 1.20+
- Make (optional, for Makefile targets)
- A running database instance (PostgreSQL or MongoDB, depending on config)
- Docker (optional, for containerized deployment)

## 📂 Project Structure

```
server/
├── cmd/api/        # Application entry point (main.go)
├── config/         # Application configuration (env)
├── docs/           # API documentation and specifications
├── internal/       # Internal packages (business logic, services, DB, handlers)
├── pkg/            # Shared helper libraries/utilities
├── .env.example    # Example environment variables
├── Makefile        # Build and run commands
├── certificate.pem # SSL certificate
├── private.key     # SSL private key
├── go.mod          # Go module dependencies
└── go.sum          # Dependency checksums
```

## 🛠️ Installation

Clone the repository and navigate to the server folder:
```bash
git clone https://github.com/khiemta03/movseek.git
cd movseek/server
```

Install dependencies:
```bash
go mod tidy
```

## 🔑 Environment Variables

Create a `.env` file in the server/ directory based on `.env.example`.

Example:
```env
MONGO_URI=
MONGO_DB_NAME=MovSeek
PORT=8080
```

## ▶️ Running the Server

```bash
go run ./cmd/api
```

Or using Make:
```bash
make run
```

## 🔒 HTTPS Support

The server supports HTTPS by default using `certificate.pem` and `private.key`. Ensure these files exist and are valid before running in production.

## 🤝 Contributing

1. Fork the project
2. Create a feature branch: `git checkout -b feature/your-feature`
3. Commit your changes: `git commit -m "Add feature"`
4. Push to the branch: `git push origin feature/your-feature`
5. Open a Pull Request

## 📄 License

This project is licensed under the MIT License.

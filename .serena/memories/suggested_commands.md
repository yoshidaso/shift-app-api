# Suggested Commands for Development

## Development Commands

### Running the Application
```bash
# Run with hot reload (recommended for development)
air

# Build and run manually
go build -o ./tmp/main .
./tmp/main

# Run directly
go run main.go
```

### Docker Commands
```bash
# Start the full stack (API + MySQL)
docker-compose up

# Start in detached mode
docker-compose up -d

# Stop services
docker-compose down

# Rebuild services
docker-compose up --build
```

### Go Commands
```bash
# Format code
go fmt ./...

# Vet code for potential issues
go vet ./...

# Run tests
go test ./...

# Build the application
go build .

# Clean build cache
go clean

# Tidy dependencies
go mod tidy

# Download dependencies
go mod download
```

### Database Commands
```bash
# Prisma commands (if using Prisma migrations)
npx prisma migrate dev
npx prisma generate
npx prisma studio
```

## Available Tools
- **Air**: Hot reload tool (configured in .air.toml)
- **Docker**: Containerization
- **Prisma**: Database schema management
- **GORM**: ORM for database operations

## Port Configuration
- **API**: localhost:8080
- **MySQL**: localhost:3306 (when running with Docker)
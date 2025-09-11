# Shift App API - Project Overview

## Purpose
This is a REST API for a shift/attendance management application. It manages users and their work shifts, tracking shift times, work content, and issues.

## Tech Stack
- **Language**: Go 1.22.1
- **Web Framework**: Gin (github.com/gin-gonic/gin v1.10.0)
- **Database**: MySQL 8.0
- **ORM**: GORM (gorm.io/gorm v1.30.1)
- **Schema Management**: Prisma (for database schema definitions)
- **Hot Reload**: Air (for development)
- **Environment**: Docker and Docker Compose
- **Dependency Management**: Go Modules

## Architecture
The project follows a clean architecture pattern with clear separation of concerns:

```
├── controllers/     # HTTP handlers (Gin controllers)
├── services/        # Business logic layer
├── repositories/    # Data access layer (GORM)
├── models/          # Database models (GORM structs)
├── dto/             # Data Transfer Objects
├── infra/           # Infrastructure (DB setup, initialization)
├── prisma/          # Database schema and migrations
└── main.go          # Application entry point
```

## Database Schema
- **Users**: id, name, email, createdAt
- **Shifts**: id, userId, startAt, endAt, workContent, issues

The application provides full CRUD operations for both users and shifts.
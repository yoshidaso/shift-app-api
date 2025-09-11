# Task Completion Checklist

When completing development tasks in this project, follow these steps:

## Code Quality Checks
1. **Format Code**: Run `go fmt ./...` to ensure consistent formatting
2. **Vet Code**: Run `go vet ./...` to catch potential issues
3. **Build Check**: Run `go build .` to ensure the code compiles
4. **Test**: Run `go test ./...` if tests exist

## Development Workflow
1. Make your changes
2. Test locally with `air` for hot reload or `go run main.go`
3. Verify the API endpoints work as expected (port 8080)
4. Check database connectivity if making data-related changes

## Pre-commit Checklist
- [ ] Code is properly formatted (`go fmt ./...`)
- [ ] No vet warnings (`go vet ./...`)
- [ ] Application builds successfully (`go build .`)
- [ ] API endpoints respond correctly
- [ ] Database migrations applied if schema changed
- [ ] Follow the established architecture pattern (controllers → services → repositories)
- [ ] Proper error handling implemented
- [ ] JSON response format is consistent

## Architecture Compliance
- Maintain separation of concerns (controllers, services, repositories, models)
- Use dependency injection pattern
- Follow interface-based design
- Implement proper error handling
- Use consistent naming conventions
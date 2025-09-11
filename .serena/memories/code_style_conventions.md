# Code Style and Conventions

## Go Conventions
The project follows standard Go conventions:

### Naming
- **Interfaces**: Prefixed with `I` (e.g., `IShiftController`, `IShiftService`)
- **Structs**: PascalCase (e.g., `ShiftController`, `Shifts`)
- **Methods**: PascalCase for exported, camelCase for unexported
- **Variables**: camelCase
- **Constants**: UPPER_CASE

### File Organization
- One main struct per file
- Files named with snake_case (e.g., `shift_controller.go`, `user_service.go`)
- Clear separation between interfaces and implementations

### Error Handling
- Standard Go error handling with explicit error returns
- HTTP status codes used appropriately (500 for internal errors, 200 for success)
- Consistent error response format: `gin.H{"error": "message"}`

### JSON Tags
- GORM tags for database mapping: `gorm:"column:name;not null"`
- JSON tags for API responses: `json:"fieldName"`
- Database column names use camelCase

### Constructor Pattern
- Factory functions named `New{StructName}` (e.g., `NewShiftController`)
- Dependency injection through constructor parameters

### HTTP Response Format
- Success responses: `gin.H{"data": result}`
- Error responses: `gin.H{"error": "message"}`
# System Utilities for Darwin (macOS)

## File System Commands
```bash
# List files
ls -la

# Find files
find . -name "*.go" -type f

# Search in files (prefer ripgrep if available)
rg "pattern" --type go
grep -r "pattern" .

# File operations
cp source destination
mv source destination
rm filename
mkdir dirname
```

## Process Management
```bash
# Find processes
ps aux | grep "process_name"
pgrep process_name

# Kill processes
kill PID
pkill process_name
killall process_name
```

## Network Commands
```bash
# Check port usage
lsof -i :8080
netstat -an | grep 8080

# Test connectivity
curl http://localhost:8080/shifts
```

## Git Commands
```bash
git status
git add .
git commit -m "message"
git push
git pull
git branch
git checkout branch_name
```

## Development Tools
- **Homebrew**: Package manager for macOS
- **Air**: Hot reload for Go (available at /Users/sota/go/bin/air)
- **Go**: Available at /usr/local/go/bin/go
- **Docker**: Container runtime
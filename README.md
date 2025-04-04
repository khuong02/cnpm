# CNPM Service Guile

---

## Go Version
- Go 1.24.2

---

## Tools
- docker
---
## Start Service
1. Start Database
```bash
docker-compose up -d
```

2. Install Dependencies
```bash
go mod tidy
```

3.Start Service
```bash
go run main.go
```
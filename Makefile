document:
	@echo "Creating swagger documentation files"
	swag init --parseDependency -g  ./cmd/api/main.go

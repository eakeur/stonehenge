document:
	@echo "Creating swagger documentation files"
	swag init --parseDependency -g  ./cmd/api/main.go

run-api:
	docker-compose up
	go run cmd/api/main.go
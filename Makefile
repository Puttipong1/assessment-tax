export ADMIN_USERNAME=adminTax
export ADMIN_PASSWORD=admin!
export PORT=9999
export DATABASE_URL=localhost
export DATABASE_USER=postgres
export DATABASE_PASSWORD=postgres
export DATABASE_PORT=5432
export DATABASE_NAME=ktaxes
run:
	docker compose up -d
	go run main.go
test:
	go test -v ./...
test-cover: 
	go test -coverprofile coverage.html .
	go tool cover -html=coverage.html
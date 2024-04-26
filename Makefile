export ADMIN_USERNAME=adminTax
export ADMIN_PASSWORD=admin!
run:
	docker compose up -d
	go run main.go
test:
	go test -v ./...
test-cover: 
	go test -coverprofile coverage.html .
	go tool cover -html=coverage.html
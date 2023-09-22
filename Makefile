run-rest:
	go run cmd/rest/main.go 

build:
	go build -o rest cmd/rest/main.go

swag-generate:
	swag init --parseDependency cmd/rest/main.go

swag-fmt:
	swag fmt
run:
	go run cmd/main.go

docker-build:
	docker build -f build/Dockerfile -t go-microservice-example .

compose-build:
	docker-compose -f deployments/docker-compose.yml up --build

compose-up:
	docker-compose -f deployments/docker-compose.yml up

compose-down:
	docker-compose -f deployments/docker-compose.yml down

postgres-up:
	docker-compose -f deployments/docker-compose.yml up postgres

proto:
	protoc --proto_path=api --go_out=plugins=grpc:api audit.proto
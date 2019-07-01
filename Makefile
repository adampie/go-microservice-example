docker-build:
	docker build -f build/Dockerfile -t go-microservice-example .

compose-up:
	docker-compose -f deployments/docker-compose.yml up

compose-down:
	docker-compose -f deployments/docker-compose.yml down

proto:
	protoc --proto_path=api --go_out=plugins=grpc:pkg/api audit.proto
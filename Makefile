docker-build:
	docker build -f build/Dockerfile -t go-microservice-example .

compose-up:
	docker-compose -f deployments/docker-compose.yml up --build

compose-down:
	docker-compose -f deployments/docker-compose.yml down

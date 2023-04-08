AUTH_SERVICE_BINARY=authServiceApp

up:
	@echo "Starting Docker images..."
	docker-compose up -d
	@echo "Docker images started!"

up_build:
	#lint
	@echo "Building (when required) and starting docker images..."
	docker-compose up --build -d
	@echo "Docker images built and started!"

down:
	@echo "Stopping docker compose..."
	docker-compose down
	@echo "docker-compose down: Done!"

lint:
	@golangci-lint run ./../... -v

# Variables
DOCKER_IMAGE_NAME=pirita_backend
DOCKER_CONTAINER_NAME=pirita_backend

build-test:
	go mod download
	go get -d -v
	go build -o backend
	chmod +x backend

build-prod:
	go mod download
	go get -d -v
	go build -ldflags="-X 'main.states=""'" -o backend

debug:
	go run main.go


# Commands for running in docker
docker-build:
	go mod download
	go get -d -v
	go build -o backend
	chmod +x backend
	docker build -t $(DOCKER_IMAGE_NAME) .

docker-run:
	docker run -d --name $(DOCKER_CONTAINER_NAME) -p 3000:3000 $(DOCKER_IMAGE_NAME)

docker-stop:
	docker stop $(DOCKER_CONTAINER_NAME)

docker-rm:
	docker rm $(DOCKER_CONTAINER_NAME)

# PHONY targets
.PHONY: build debug docker-build docker-run docker-stop docker-rm

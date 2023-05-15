# Variables
DOCKER_IMAGE_NAME=frontend
DOCKER_CONTAINER_NAME=pirita-frontend

build:
	npm install
	npm run build

debug:
	npm install
	npm run dev

test_before_push:
	npm install
	npm run test
	npm run check
	npm run test:unit
	npm run lint
	npm run format

test:
	npm install
	npm run test
	npm run test:unit
	npm run lint

# Commands for running in containers

docker-build:
	docker build -t $(DOCKER_IMAGE_NAME) .

docker-run:
	docker run -d -p 80:80 --name $(DOCKER_CONTAINER_NAME) $(DOCKER_IMAGE_NAME)

docker-stop:
	docker stop $(DOCKER_CONTAINER_NAME)

docker-rm:
	docker rm $(DOCKER_CONTAINER_NAME)

# PHONY targets
.PHONY: build debug test_before_push docker-build docker-run docker-stop docker-rm

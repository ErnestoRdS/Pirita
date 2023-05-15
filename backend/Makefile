build_backend:
	go mod download && go get -d -v && go build -o backend && chmod +x backend

build:
	docker build -t backend .

run:
	docker run --rm -it -p 3000:3000 backend

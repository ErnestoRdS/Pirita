FROM golang:latest

WORKDIR /app

COPY . .

RUN go mod download && go get -d -v && go build -o backend

EXPOSE 3000

CMD ["./backend"]

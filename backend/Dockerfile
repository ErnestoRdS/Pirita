FROM golang:latest

WORKDIR /app

COPY . .

RUN make build

EXPOSE 3000

CMD ["./backend"]

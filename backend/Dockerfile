FROM golang:latest

WORKDIR /app

COPY . .

RUN make build_backend

EXPOSE 3000

CMD ["./backend"]

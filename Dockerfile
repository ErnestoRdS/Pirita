FROM golang:latest AS backend

WORKDIR /app/backend
COPY backend/ .

RUN go mod download && go get -d -v && go build -o backend

FROM node:latest AS frontend

WORKDIR /app/frontend
COPY frontend/ .

RUN npm install && npm run build

FROM alpine:latest

WORKDIR /app

COPY --from=backend /app/backend/backend /app/backend
COPY --from=frontend /app/frontend/build /app/frontend

EXPOSE 3000
EXPOSE 5173

CMD ["./app/backend"]

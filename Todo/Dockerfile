FROM golang:1.19.4

WORKDIR /app/gin-todo-prc

COPY go.mod .
COPY go.sum .

RUN go mod tidy 
RUN go mod download

COPY postgres-db.sql /docker-entrypoint-initdb.db/
COPY . .


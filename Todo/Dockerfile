# FROM golang:1.19.4

# RUN mkdir /todo

# ADD . /todo

# WORKDIR /todo

# COPY go.mod .
# COPY go.sum .

# RUN go mod tidy 
# RUN go mod download

# COPY *.go ./

# RUN go build -o main ./main.go

# EXPOSE 8080

# CMD [ "/todo/main"]

FROM golang:latest

WORKDIR /todo_1
COPY . .
RUN go build -o main .
EXPOSE 8080
CMD ["./"]
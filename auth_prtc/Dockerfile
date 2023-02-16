FROM golang:1.19.4

RUN mkdir /app

ADD . /app

WORKDIR /app

COPY . .

RUN go build -o main .

# COPY go.mod ./
# COPY go.sum ./

# RUN go mod download

# COPY *.go ./

EXPOSE 8080

CMD [ "./main" ]
FROM golang:1.16.4

WORKDIR /go/src/tutorial/mvc

COPY go.mod go.sum ./
COPY .env ./
RUN go mod download
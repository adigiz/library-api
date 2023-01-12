# syntax=docker/dockerfile:1
FROM golang:1.18-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

#COPY *.go ./
COPY . ./

ENV GIN_MODE=release

RUN go build -o /library-api

EXPOSE 8080

CMD [ "/library-api" ]

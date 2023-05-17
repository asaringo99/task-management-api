#### Build ####
FROM golang:1.20-buster AS build

WORKDIR /app

COPY ./go.mod /app
COPY ./go.sum /app

RUN go mod download

COPY *.go /app

# RUN go build -o /docker-gs-ping
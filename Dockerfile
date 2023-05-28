FROM golang:1.20.4-alpine3.18 AS build

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod tidy

COPY . .

RUN go build -o main cmd/rest/main.go

FROM alpine:latest

WORKDIR /

COPY --from=build app/main /main
COPY --from=build app/pkg/simpDB/db simpDB/db

CMD ["./main"]


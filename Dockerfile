FROM 1.20.4-alpine3.17 AS build

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod tidy

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o btcinform

FROM alpine:latest

WORKDIR /app

COPY --from=build /app/btcinform .

EXPOSE 8080

CMD ["./btcinform"]


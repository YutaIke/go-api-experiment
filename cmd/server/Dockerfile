FROM golang:1.20.1-alpine AS builder

WORKDIR /app

RUN apk update && apk add --no-cache git
COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build --trimpath --ldflags "-s -w" -o go-api-experiment

# ------------------------------
FROM debian:bookworm-slim AS deploy

RUN apt update

COPY --from=deploy-builder /app/go-api-experiment .

CMD ["./go-api-experiment"]

# ------------------------------
FROM golang:1.20.1-alpine AS dev
WORKDIR /go/src/app
RUN go install github.com/cosmtrek/air@latest
CMD ["air", "-c", ".air.toml"]
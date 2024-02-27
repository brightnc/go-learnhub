FROM golang:1.22.0-alpine3.18 AS build
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
COPY .env ./.env

RUN CGO_ENABLED=0 GOOS=linux go build -o go-learnhub-api ./cmd

FROM alpine:latest
WORKDIR /app

COPY --from=build /app/go-learnhub-api .
COPY --from=build /app/.env .

EXPOSE 8000

CMD ["./go-learnhub-api"]

FROM golang:1.22.0-alpine as builder

RUN apk update && \
    apk upgrade && \
    apk add curl

WORKDIR /app

COPY . .

RUN go mod tidy

CMD ["go", "run", "/app/cmd/api/main.go"]
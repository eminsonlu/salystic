# Build stage
FROM golang:1.23-alpine AS builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o import ./cmd/import

# Production stage
FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /app

COPY --from=builder /app/import .

ENV DATA_FILE=data.json

VOLUME ["/app/data"]

CMD ["sh", "-c", "./import ${DATA_FILE}"]
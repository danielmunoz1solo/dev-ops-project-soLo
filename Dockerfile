# multi stage setup for decreased file size, golang as builder
FROM golang:1.24.5-alpine AS builder

WORKDIR /app

# layer caching
COPY go.mod go.sum ./
RUN go mod download

COPY . .

# build for alpine linux, no build to C used
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /app/quote-scraper ./cmd/quote-scraper

# runner
FROM alpine:latest

RUN apk add --no-cache wget

WORKDIR /app

# copy only compiled file to keep size minimum
COPY --from=builder /app/quote-scraper .

EXPOSE 8080

HEALTHCHECK --interval=30s --timeout=10s --retries=3 \
  CMD wget --quiet --tries=1 --timeout=5 -O /dev/null http://localhost:8080/daily-quotes || exit 1  

CMD ["./quote-scraper"]
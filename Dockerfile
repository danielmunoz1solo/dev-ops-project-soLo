# local dockerfile as template
# multi stage setup for decreased file size, golang as builder
FROM golang:1.24-alpine AS builder

WORKDIR /app

# layer caching
COPY go.mod go.sum ./
RUN go mod download

COPY . .

# build for alpine linux, no build to C used
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /app/quote-scraper .

# runner
FROM alpine:latest

WORKDIR /app

# copy only compiled file to keep size minimum
COPY --from=builder /app/quote-scraper .

EXPOSE 8080

CMD ["./quote-scraper"]
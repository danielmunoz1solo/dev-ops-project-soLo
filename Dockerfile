FROM golang:1.24-alpine

WORKDIR /app

# Copy all files into the container's /app directory
COPY . .

# Download the dependencies listed in go.mod/go.sum
RUN go mod tidy

# Build the Go application inside the container
RUN go build -o /quote-scraper .

# Expose the port that your server runs on
EXPOSE 8080

# The command to run when the container starts
CMD ["/quote-scraper"]
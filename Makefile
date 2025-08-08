.PHONY: build run test docker-build docker-push clean

IMAGE_NAME = dmun1solo/daily-quote-scraper
VERSION = latest

# default target set to build
all: build

build: 
go build -o quote-scraper ./cmd/quote-scraper

run: build
./quote-scraper

# test: all packages
test: 
go test ./...

build-image:
docker build -t $(IMAGE_NAME):$(VERSION)

push-image:
docker push $(IMAGE_NAME):$(VERSION)

clean:
rm -f quote-scraper


Go Quote Scraper API
=====================
This project scrapes the first 100 quotes from https://quotes.toscrape.com
and exposes them via a local JSON API endpoint at:
```
http://localhost:8080/daily-quotes
```

This project utilizes the Go web-scraping framework [Colly](https://github.com/gocolly/colly)

##### You can run this app in several ways depending on your setup

## Requirements:

- Internet access to reach the quotes endpoint

## Runing my service using Go

### OPTION 1: Clone the repository
```
git clone https://github.com/danielmunoz1solo/dev-ops-project-soLo.git
cd dev-ops-project-soLo
```
- Or just download the ZIP file from GitHub and extract it.

Then start the server by running:
```
go run main.go
```

### Option 2: Pull the go module:

Run the executable:
```
./quote-scraper (macOS/Linux)
./quote-scraper.exe (windows)
```

Both options will start the web server at:
```
http://localhost:8080/daily-quotes
```
**Press Ctrl+C to stop the server**

## How to Run Without Installing Go (TODO: setup releases page)

### OPTION 1: Download the prebuilt executable

Visit the GitHub Releases page:

    https://github.com/danielmunoz1solo/dev-ops-project-soLo/releases

Download the appropriate file for your system:

- macOS/Linux:   [quote-scraper]()
- Windows:       [quote-scraper.exe]()

Then run it from your terminal:
```
./quote-scraper      (Linux/macOS)
quote-scraper.exe    (Windows)
```
- Once it's running, visit: http://localhost:8080/daily-quotes


### OPTION 2: Run using Docker

If you have Docker installed, you can run the app without building from source. Pull the pre-built image from Docker Hub and run it. 

- Or follow the instructions to install:
    - [Docker Engine](https://docs.docker.com/engine/install/) - lightweight command-line interface tool for running Docker containers. Recommended for Linux distributions or developers who prefer terminal-based workflows.
    - [Docker Desktop](https://docs.docker.com/get-started/get-docker/) - application that includes Docker Engine, a GUI dashboar and integrated Kubernetes. Recommended for most local development environments.

Then pull the image using:
```bash
docker pull dmun1/quote-scraper:latest
```

- This command runs the container and maps port 8080 on your machine to the container's port 8080.
```
docker run -p 8080:8080 dmun1/quote-scraper:latest
```

And finally, visit the endpoint in your browser at: http://localhost:8080/daily-quotes

## Project Structure

    main.go         - Entry point
    scraper/        - Web scraper using Colly
    quotes/         - Data model for quotes
    server/         - HTTP server logic
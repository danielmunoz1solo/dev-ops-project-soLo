Go Quote Scraper API
=====================
## Work in Progress

This project scrapes the first 100 quotes from https://quotes.toscrape.com
and exposes them via a local JSON API endpoint at:
```
http://localhost:8080/quotes
```

This project utilizes the Go web-scraping framework [Colly](https://github.com/gocolly/colly)

##### You can run this app in several ways depending on your setup

## Requirements:

- Git **(for cloning the repository)**
- Internet access to reach the quotes website
- Go 1.18 or higher installed [(install)](https://go.dev/dl/) **(Only if building from source)**
## How to Run From Source

#### Step 1: Clone the repository
```
git clone https://github.com/danielmunoz1solo/dev-ops-project-soLo.git
cd dev-ops-project-soLo
```
- OR just download the ZIP file from GitHub and extract it.

#### Run the server

##### Option 1: Run main.go directly from the root path of the repository using:
```
go run main.go
```

##### Option 2: Build to binary:
```
go build -o quote-scraper.exe
```

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

## How to Run Without Installing Go (TODO: setup releases page + docker)

##### OPTION 1: Download the prebuilt executable

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
#### OPTION 2: Run using Docker

If you have Docker installed, you can run the app without installing anything else.
```
docker run -p 8080:8080 SETUP WITH MY DOCKER REGISTRY
```
Then visit:
```
http://localhost:8080/daily-quotes
```

## Project Structure


    main.go         - Entry point
    scraper/        - Web scraper using Colly
    quotes/         - Data model for quotes
    server/         - HTTP server logic
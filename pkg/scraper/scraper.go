package scraper

import (
	"fmt"

	"github.com/gocolly/colly"
	"github.com/danielmunoz1solo/dev-ops-project-soLo/pkg/quotes"
)

func ScrapeQuotes() ([]quotes.Quote, error) {
	// Channel to hold the quotes scraped by the goroutines
	quotesChan := make(chan quotes.Quote)
	var dailyQuotes []quotes.Quote
	const maxQuotes = 100

	c := colly.NewCollector(
		// Use async for concurrency
		colly.Async(true),
	)

	// Set a parallelism limit
	c.Limit(&colly.LimitRule{DomainGlob: "*", Parallelism: 4})

	c.OnHTML(".quote", func(e *colly.HTMLElement) {
		var tags []string
		e.ForEach(".tags .tag", func(_ int, el *colly.HTMLElement) {
			tags = append(tags, el.Text)
		})

		quote := quotes.Quote{
			Text:   e.ChildText(".text"),
			Author: e.ChildText(".author"),
			Tags:   tags,
		}
		
		// Send the new quote to our channel
		quotesChan <- quote
	})

	c.OnHTML(".next a", func(e *colly.HTMLElement) {
		nextPage := e.Request.AbsoluteURL(e.Attr("href"))
		e.Request.Visit(nextPage)
	})

	// Start the scraping process in a separate goroutine.
	// This goroutine will be responsible for closing the channel when it's done.
	go func() {
		// When this goroutine finishes, close the channel
		defer close(quotesChan)

		err := c.Visit("https://quotes.toscrape.com")
		if err != nil {
			fmt.Println("Error visiting initial page:", err)
		}
		
		// Wait for all scraping jobs to finish
		c.Wait()
	}()

	// Collect quotes from the channel until it's closed or we have enough.
	for quote := range quotesChan {
		if len(dailyQuotes) < maxQuotes {
			dailyQuotes = append(dailyQuotes, quote)
		} else {
			// Once we have 100 quotes, we can stop collecting.
			break
		}
	}

	return dailyQuotes, nil
}


// ~3 seconds slower with no concurrency
// import (
// 	"fmt"

// 	"github.com/gocolly/colly"
// 	"dev_ops_th/quotes"
// )

// func ScrapeQuotes() ([]quotes.Quote, error) {
// 	c := colly.NewCollector()
// 	var dailyQuotes []quotes.Quote
// 	const maxQuotes = 100

// 	c.OnHTML(".quote", func(e *colly.HTMLElement) {
// 		if len(dailyQuotes) >= maxQuotes {
// 			return
// 		}

// 		var tags []string
// 		e.ForEach(".tags .tag", func(_ int, el *colly.HTMLElement) {
// 			tags = append(tags, el.Text)
// 		})

// 		quote := quotes.Quote{
// 			Text:   e.ChildText(".text"),
// 			Author: e.ChildText(".author"),
// 			Tags:   tags,
// 		}
		
// 		dailyQuotes = append(dailyQuotes, quote)
// 	})


// 	c.OnHTML(".next a", func(e *colly.HTMLElement) {
// 		if len(dailyQuotes) < maxQuotes {
// 			nextPage := e.Request.AbsoluteURL(e.Attr("href"))
// 			e.Request.Visit(nextPage)
// 		}
// 	})

// 	c.OnRequest(func(r *colly.Request) {
// 		fmt.Println("Request received, now visiting:", r.URL)
// 	})

// 	err := c.Visit("https://quotes.toscrape.com")
// 	if err != nil {
// 		return nil, err
// 	}

// 	return dailyQuotes, nil
// }
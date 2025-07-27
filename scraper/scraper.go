package scraper

import (
	"fmt"

	"github.com/gocolly/colly"
	"dev_ops_th/quotes"
)

func ScrapeQuotes() ([]quotes.Quote, error) {
	c := colly.NewCollector()
	var dailyQuotes []quotes.Quote
	const maxQuotes = 100

	c.OnHTML(".quote", func(e *colly.HTMLElement) {
		if len(dailyQuotes) >= maxQuotes {
			return
		}

		var tags []string
		e.ForEach(".tags .tag", func(_ int, el *colly.HTMLElement) {
			tags = append(tags, el.Text)
		})

		quote := quotes.Quote{
			Text:   e.ChildText(".text"),
			Author: e.ChildText(".author"),
			Tags:   tags,
		}
		
		dailyQuotes = append(dailyQuotes, quote)
	})


	c.OnHTML(".next a", func(e *colly.HTMLElement) {
		if len(dailyQuotes) < maxQuotes {
			nextPage := e.Request.AbsoluteURL(e.Attr("href"))
			e.Request.Visit(nextPage)
		}
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Request received, now visiting:", r.URL)
	})

	err := c.Visit("https://quotes.toscrape.com")
	if err != nil {
		return nil, err
	}

	return dailyQuotes, nil
}
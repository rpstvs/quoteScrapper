package main

import (
	"fmt"
	"log"
	"time"

	"github.com/gocolly/colly"
)

func main() {
	var quotes []Resultado
	c := colly.NewCollector(
		colly.AllowedDomains("www.goodreads.com"),
	)

	c.OnRequest(func(r *colly.Request) {
		log.Printf("Visiting: %s \n", r.URL)
	})

	c.OnHTML(".quote", func(h *colly.HTMLElement) {
		result := Resultado{}
		result.Quote = h.ChildText("div.quoteText")
		Parser(result.Quote)
		quotes = append(quotes, result)
	})

	for i := 1; i < 2; i++ {
		time.Sleep(5 * time.Second)
		url := fmt.Sprintf("https://www.goodreads.com/author/quotes/17212.Marcus_Aurelius?page=%d", i)
		c.Visit(url)
	}

	//fmt.Println(quotes)

}

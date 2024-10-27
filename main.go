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
		s := h.ChildText("div.quoteText")
		result = Parser(s)
		quotes = append(quotes, result)
	})

	authors := []string{"4918776.Seneca", "17212.Marcus_Aurelius", "13852.Epictetus"}

	for i := 1; i < 10; i++ {
		for _, author := range authors {
			time.Sleep(5 * time.Second)
			url := fmt.Sprintf("https://www.goodreads.com/author/quotes/%s?page=%d", i, author)
			c.Visit(url)
		}

	}

	//fmt.Println(quotes)

}

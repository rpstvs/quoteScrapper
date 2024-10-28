package scrapper

import (
	"context"

	"github.com/rpstvs/quoteScrapper/database"
)

func (cfg *Configure) WriteToDB(quotes []Resultado) {

	for _, quote := range quotes {
		cfg.db.InsertQuote(context.Background(), database.InsertQuoteParams{
			Quote:  quote.Quote,
			Author: quote.Author,
			Book:   quote.Book,
		})
	}

}

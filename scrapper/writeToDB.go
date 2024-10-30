package scrapper

import (
	"context"

	"github.com/rpstvs/quoteScrapper/database"
)

func (cfg *Configure) WriteToDB(quote Resultado) {

	if len(quote.Author) == 1 {
		quote.Author = quote.Book

	}
	cfg.db.InsertQuote(context.Background(), database.InsertQuoteParams{
		Quote:  quote.Quote,
		Author: quote.Author,
		Book:   quote.Book,
	})

}

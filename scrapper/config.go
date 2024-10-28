package scrapper

import (
	"database/sql"
	"os"

	"github.com/gocolly/colly"
	"github.com/rpstvs/quoteScrapper/database"
)

type Configure struct {
	c  *colly.Collector
	db *database.Queries
}

func InitConfig(domain string) *Configure {
	dbUrl := os.Getenv("DATABASE_URL")
	db, _ := sql.Open("postgres", dbUrl)
	queries := database.New(db)
	return &Configure{
		c:  colly.NewCollector(colly.AllowedDomains(domain)),
		db: queries,
	}
}

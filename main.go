package main

import (
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/rpstvs/quoteScrapper/scrapper"
)

func main() {

	godotenv.Load(".env")
	cfg := scrapper.InitConfig("www.goodreads.com")

	cfg.GetQuotes()
	//fmt.Println(quotes)

}

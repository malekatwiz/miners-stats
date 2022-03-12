package scraper

import (
	"log"

	"github.com/gocolly/colly"
	"github.com/gocolly/colly/debug"
)

type Operation struct {
	Currency string
	PoolUrl  string
	Wallet   string
}

func (o Operation) ReadOperationStats() {
	c := colly.NewCollector(colly.Debugger(&debug.LogDebugger{}))
	c.OnError(func(r *colly.Response, err error) {
		log.Println(err.Error())
	})
	c.OnHTML("", func(h *colly.HTMLElement) {

	})
}

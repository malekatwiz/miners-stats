package minerstats

import (
	"fmt"

	minerstats "github.com/malekatwiz/miners-stats-insights/models"
	"github.com/malekatwiz/miners-stats-insights/scraper"
)

func GetOperationWorkers(currency string) []minerstats.Worker {
	stats := scraper.GetHeroMinersStats("https://flux.herominers.com/api/stats_address?address=t1VNTLUKB1bamRQ8ExgsrMaM3dfoQGct3bd")
	return []minerstats.Worker{
		{
			Name:          "miner-1",
			Hashrate:      fmt.Sprint(stats.Stats.Hashrate),
			Currency:      "flux",
			Fiat:          "usd",
			ExchangeValue: 1.35,
		},
	}
}

package minerstats

import (
	"fmt"
	"os"

	minerstats "github.com/malekatwiz/miners-stats-insights/models"
	"github.com/malekatwiz/miners-stats-insights/scraper"
)

func GetOperationWorkers(wallet string) []minerstats.Worker {
	stats := scraper.GetHeroMinersStats(fmt.Sprintf("https://flux.herominers.com/api/stats_address?address=%s", wallet))
	return []minerstats.Worker{
		{
			Name:     "miner-1",
			Hashrate: fmt.Sprint(stats.Stats.Hashrate),
		},
	}
}

func GetOperation(currency string) minerstats.Operation {
	os.Getenv("")
	return minerstats.Operation{}
}

package minerstats

import (
	minerstats "github.com/malekatwiz/miners-stats-insights/models"
)

func GetOperationWorkers(currency string) []minerstats.Worker {
	return []minerstats.Worker{
		{
			Name:          "miner-1",
			Hashrate:      "300 sol/s",
			Currency:      "flux",
			Fiat:          "usd",
			ExchangeValue: 1.3,
		},
	}
}

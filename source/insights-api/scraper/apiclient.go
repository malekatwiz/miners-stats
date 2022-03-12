package scraper

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
)

func GetHeroMinersStats(url string) HeroMinersStats {
	var stats HeroMinersStats
	response, e := http.Get(url)
	if e == nil && response.StatusCode == http.StatusOK {
		defer response.Body.Close()
		responseContent, e := io.ReadAll(response.Body)
		if e == nil {
			json.NewDecoder(bytes.NewBuffer(responseContent)).Decode(&stats)
		}
	} else {
		log.Println(e.Error())
	}
	return stats
}

type HeroMinersStats struct {
	Stats HMStats `json:"stats"`
}

type HMStats struct {
	ValidShares  int64   `json:"shares_good"`
	Blocks       int     `json:"blocksFound"`
	Hashrate     float32 `json:"hashrate"`
	Hashrate1Hr  float32 `json:"hashrate_1h"`
	Hashrate6Hr  float32 `json:"hashrate_6h"`
	Hashrate24Hr float32 `json:"hashrate_24h"`
}

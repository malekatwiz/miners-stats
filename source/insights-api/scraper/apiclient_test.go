package scraper

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_GetHRstats_ReturnsAllStats(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Write([]byte(`{
			"stats": {
				"donation_level": "0",
				"shares_good": 74947,
				"hashes": "384656380",
				"lastShare": "1646597962",
				"balance": "84906234",
				"unpaid_pa_balance": "2323217198",
				"solo_shares_good": 9314,
				"solo_hashes": "53678086",
				"shares_invalid": 39809,
				"shares_stale": 19,
				"paid": "8046355163",
				"blocksFound": "2",
				"hashrate": 268.095,
				"roundScore": 92787,
				"roundHashes": 92787,
				"poolRoundScore": 87929348,
				"poolRoundHashes": 88025187,
				"networkHeight": 1071564,
				"hashrate_1h": 254.2,
				"hashrate_6h": 267.6551724137931,
				"hashrate_24h": 250.9579831932773,
				"solo_shares_invalid": 0,
				"solo_shares_stale": 0,
				"soloRoundHashes": 53678086,
				"payments_24h": 465803268,
				"payments_7d": 3490928387
			}
		}`))
	}))
	defer testServer.Close()

	stats := GetHeroMinersStats(testServer.URL)

	expectedBlocks := 2
	expectedHashrate := 268.095

	if stats.Stats.Blocks != expectedBlocks {
		t.Errorf("expected '%s', found '%s'", expectedBlocks)
	}
	if stats.Stats.Hashrate != float32(expectedHashrate) {
		t.Errorf("expected '%s', found '%s'", expectedHashrate, stats.Stats.Hashrate)
	}
}

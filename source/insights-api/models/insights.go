package minerstats

type Worker struct {
	Name          string  `json:"name"`
	Hashrate      string  `json:"hashrate"`
	Currency      string  `json:"currency"`
	Fiat          string  `json:"fiat"`
	ExchangeValue float32 `json"exchangeValue"`
}

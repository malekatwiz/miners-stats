package minerstats

type Worker struct {
	Name     string `json:"name"`
	Hashrate string `json:"hashrate"`
}

type Operation struct {
	Currency string   `json:"currency"`
	Workers  []Worker `json:"workers"`
}

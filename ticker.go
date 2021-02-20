package client

import (
	"encoding/json"
)

type CoinTicker struct {
	Name string
	EUR  string `json:"EUR"`
	USD  string `json:"USD"`
	CHF  string `json:"CHF"`
	GBP  string `json:"GBP"`
	TRY  string `json:"TRY"`
}

func (c *Client) TickerGet() (coins []CoinTicker, err error) {
	req, err := c.newGetRequest("/ticker", nil)
	if err != nil {
		return
	}

	var f interface{}
	_, _, err = c.do(req, &f)
	if err != nil {
		return
	}

	itemsMap := f.(map[string]interface{})
	for key, value := range itemsMap {
		coinByte, _ := json.Marshal(value)
		coin := CoinTicker{}
		_ = json.Unmarshal(coinByte, &coin)
		coin.Name = key
		coins = append(coins, coin)
	}

	return
}

package client

import "time"

type Trade struct {
	ID         string `json:"id"`
	Type       string `json:"type"`
	Attributes struct {
		Status           string `json:"status"`
		Type             string `json:"type"`
		CryptocoinID     string `json:"cryptocoin_id"`
		FiatID           string `json:"fiat_id"`
		AmountFiat       string `json:"amount_fiat"`
		AmountCryptocoin string `json:"amount_cryptocoin"`
		FiatToEurRate    string `json:"fiat_to_eur_rate"`
		WalletID         string `json:"wallet_id"`
		FiatWalletID     string `json:"fiat_wallet_id"`
		PaymentOptionID  string `json:"payment_option_id"`
		Time             struct {
			DateIso8601 time.Time `json:"date_iso8601"`
			Unix        string    `json:"unix"`
		} `json:"time"`
		Price  string `json:"price"`
		IsSwap bool   `json:"is_swap"`
	} `json:"attributes"`
}

type TradesGetResponse struct {
	Data  []Trade       `json:"data"`
	Meta  ResponseMeta  `json:"meta"`
	Links ResponseLinks `json:"links"`
}

func (c *Client) TradesGet(tradeType string, page ParamsPage) (retResp TradesGetResponse, err error) {
	queryParams := page.ToQueryParams()
	if tradeType != "" {
		queryParams["type"] = tradeType
	}

	req, err := c.newGetRequest("/trades", queryParams)
	if err != nil {
		return
	}

	_, _, err = c.do(req, &retResp)
	return
}

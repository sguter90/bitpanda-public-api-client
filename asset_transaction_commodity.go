package client

import "time"

type AssetTransactionsCommodity struct {
	Type       string `json:"type"`
	Attributes struct {
		Amount    string `json:"amount"`
		Recipient string `json:"recipient"`
		Time      struct {
			DateIso8601 time.Time `json:"date_iso8601"`
			Unix        string    `json:"unix"`
		} `json:"time"`
		InOrOut      string `json:"in_or_out"`
		Type         string `json:"type"`
		Status       string `json:"status"`
		AmountEur    string `json:"amount_eur"`
		WalletID     string `json:"wallet_id"`
		Confirmed    bool   `json:"confirmed"`
		CryptocoinID string `json:"cryptocoin_id"`
		Trade        struct {
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
				Price     string `json:"price"`
				IsSwap    bool   `json:"is_swap"`
				IsSavings bool   `json:"is_savings"`
			} `json:"attributes"`
			ID string `json:"id"`
		} `json:"trade"`
		LastChanged struct {
			DateIso8601 time.Time `json:"date_iso8601"`
			Unix        string    `json:"unix"`
		} `json:"last_changed"`
		Fee               string   `json:"fee"`
		CurrentFiatID     string   `json:"current_fiat_id"`
		CurrentFiatAmount string   `json:"current_fiat_amount"`
		TxID              string   `json:"tx_id"`
		IsSavings         bool     `json:"is_savings"`
		IsMetalStorageFee bool     `json:"is_metal_storage_fee"`
		Tags              []string `json:"tags"`
	} `json:"attributes"`
	ID string `json:"id"`
}

type AssetTransactionsCommodityGetResponse struct {
	Data  []AssetTransactionsCommodity `json:"data"`
	Meta  ResponseMeta                 `json:"meta"`
	Links ResponseLinks                `json:"links"`
}

func (c *Client) AssetTransactionsCommodityGet(page ParamsPage) (retResp AssetTransactionsCommodityGetResponse, err error) {
	queryParams := page.ToQueryParams()

	req, err := c.newGetRequest("/assets/transactions/commodity", queryParams)
	if err != nil {
		return
	}

	_, _, err = c.do(req, &retResp)
	return
}

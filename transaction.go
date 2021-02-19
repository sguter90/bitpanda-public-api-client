package client

import "time"

type Transaction struct {
	ID         string `json:"id"`
	Type       string `json:"type"`
	Attributes struct {
		Amount    string `json:"amount"`
		Recipient string `json:"recipient"`
		Time      struct {
			DateIso8601 time.Time `json:"date_iso8601"`
			Unix        string    `json:"unix"`
		} `json:"time"`
		Confirmations              int    `json:"confirmations"`
		InOrOut                    string `json:"in_or_out"`
		Type                       string `json:"type"`
		Status                     string `json:"status"`
		AmountEur                  string `json:"amount_eur"`
		PurposeText                string `json:"purpose_text"`
		RelatedWalletTransactionID string `json:"related_wallet_transaction_id"`
		RelatedWalletID            string `json:"related_wallet_id"`
		WalletID                   string `json:"wallet_id"`
		Confirmed                  bool   `json:"confirmed"`
		CryptocoinID               string `json:"cryptocoin_id"`
		LastChanged                struct {
			DateIso8601 time.Time `json:"date_iso8601"`
			Unix        string    `json:"unix"`
		} `json:"last_changed"`
		Fee               string `json:"fee"`
		CurrentFiatID     string `json:"current_fiat_id"`
		CurrentFiatAmount string `json:"current_fiat_amount"`
		TxID              string `json:"tx_id"`
	} `json:"attributes"`
}

type WalletsTransactionsGet struct {
	Data  []Transaction `json:"data"`
	Meta  ResponseMeta  `json:"meta"`
	Links ResponseLinks `json:"links"`
}

func (c *Client) WalletTransactionGet(transactionType string, status string, page ParamsPage) (retResp WalletsTransactionsGet, err error) {
	queryParams := page.ToQueryParams()
	if transactionType != "" {
		queryParams["type"] = transactionType
	}
	if status != "" {
		queryParams["status"] = status
	}

	req, err := c.newGetRequest("/wallets/transactions", queryParams)
	if err != nil {
		return
	}

	_, _, err = c.do(req, &retResp)
	return
}

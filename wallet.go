package client

type Wallet struct {
	ID         string `json:"id"`
	Type       string `json:"type"`
	Attributes struct {
		CryptocoinID             string `json:"cryptocoin_id"`
		CryptocoinSymbol         string `json:"cryptocoin_symbol"`
		Balance                  string `json:"balance"`
		IsDefault                bool   `json:"is_default"`
		Name                     string `json:"name"`
		PendingTransactionsCount int    `json:"pending_transactions_count"`
		Deleted                  bool   `json:"deleted"`
	} `json:"attributes"`
}

type WalletsGetResponse struct {
	Data []Wallet `json:"data"`
}

func (c *Client) WalletsGet() (retResp WalletsGetResponse, err error) {
	req, err := c.newGetRequest("/wallets", nil)
	if err != nil {
		return
	}

	_, _, err = c.do(req, &retResp)
	return
}

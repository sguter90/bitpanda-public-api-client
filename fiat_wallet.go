package client

type FiatWallet struct {
	Type       string `json:"type"`
	Attributes struct {
		FiatID                   string `json:"fiat_id"`
		FiatSymbol               string `json:"fiat_symbol"`
		Balance                  string `json:"balance"`
		Name                     string `json:"name"`
		PendingTransactionsCount int    `json:"pending_transactions_count"`
	} `json:"attributes"`
	ID string `json:"id"`
}

type FiatWalletsGetResponse struct {
	Data []FiatWallet `json:"data"`
}

func (c *Client) FiatWalletsGet() (retResp FiatWalletsGetResponse, err error) {
	req, err := c.newGetRequest("/fiatwallets", nil)
	if err != nil {
		return
	}

	_, _, err = c.do(req, &retResp)
	return
}

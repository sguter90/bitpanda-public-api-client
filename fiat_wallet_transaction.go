package client

import "time"

type FiatWalletTransaction struct {
	Type       string `json:"type"`
	Attributes struct {
		FiatWalletID string `json:"fiat_wallet_id"`
		UserID       string `json:"user_id"`
		FiatID       string `json:"fiat_id"`
		Amount       string `json:"amount"`
		Fee          string `json:"fee"`
		ToEurRate    string `json:"to_eur_rate"`
		Time         struct {
			DateIso8601 time.Time `json:"date_iso8601"`
			Unix        string    `json:"unix"`
		} `json:"time"`
		InOrOut             string `json:"in_or_out"`
		Type                string `json:"type"`
		Status              string `json:"status"`
		ConfirmationBy      string `json:"confirmation_by"`
		Confirmed           bool   `json:"confirmed"`
		PaymentOptionID     string `json:"payment_option_id"`
		Requires2FaApproval bool   `json:"requires_2fa_approval"`
		LastChanged         struct {
			DateIso8601 time.Time `json:"date_iso8601"`
			Unix        string    `json:"unix"`
		} `json:"last_changed"`
	} `json:"attributes"`
	ID string `json:"id"`
}

type FiatWalletsTransactionsGetResponse struct {
	Data  []FiatWalletTransaction `json:"data"`
	Meta  ResponseMeta            `json:"meta"`
	Links ResponseLinks           `json:"links"`
}

func (c *Client) FiatWalletsTransactionsGet(transactionType string, status string, page ParamsPage) (retResp FiatWalletsTransactionsGetResponse, err error) {
	queryParams := page.ToQueryParams()
	if transactionType != "" {
		queryParams["type"] = transactionType
	}
	if status != "" {
		queryParams["status"] = status
	}

	req, err := c.newGetRequest("/fiatwallets/transactions", queryParams)
	if err != nil {
		return
	}

	_, _, err = c.do(req, &retResp)
	return
}

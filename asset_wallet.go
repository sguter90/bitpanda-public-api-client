package client

import "time"

type AssetWallet struct {
	Type       string `json:"type"`
	Attributes struct {
		Cryptocoin struct {
			Type       string `json:"type"`
			Attributes struct {
				Wallets []struct {
					Type       string `json:"type"`
					Attributes struct {
						CryptocoinID     string `json:"cryptocoin_id"`
						CryptocoinSymbol string `json:"cryptocoin_symbol"`
						Balance          string `json:"balance"`
						IsDefault        bool   `json:"is_default"`
						Name             string `json:"name"`
						Deleted          bool   `json:"deleted"`
					} `json:"attributes"`
					ID string `json:"id"`
				} `json:"wallets"`
			} `json:"attributes"`
		} `json:"cryptocoin"`
		Commodity struct {
			Metal struct {
				Type       string `json:"type"`
				Attributes struct {
					Wallets []struct {
						Type       string `json:"type"`
						Attributes struct {
							CryptocoinID     string `json:"cryptocoin_id"`
							CryptocoinSymbol string `json:"cryptocoin_symbol"`
							Balance          string `json:"balance"`
							IsDefault        bool   `json:"is_default"`
							Name             string `json:"name"`
							Deleted          bool   `json:"deleted"`
						} `json:"attributes"`
						ID string `json:"id"`
					} `json:"wallets"`
				} `json:"attributes"`
			} `json:"metal"`
		} `json:"commodity"`
	} `json:"attributes"`
}

type AssetsWalletsGetResponse struct {
	Data           AssetWallet `json:"data"`
	LastUserAction struct {
		DateIso8601 time.Time `json:"date_iso8601"`
		Unix        string    `json:"unix"`
	} `json:"last_user_action"`
}

func (c *Client) AssetWalletsGet() (retResp AssetsWalletsGetResponse, err error) {
	req, err := c.newGetRequest("/asset-wallets", nil)
	if err != nil {
		return
	}

	_, _, err = c.do(req, &retResp)
	return
}

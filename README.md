# 🐼 Bitpanda public API go client

An unofficial Go client for the "Public API" of [Bitpanda](https://www.bitpanda.com).  
Bitpanda is a trading platform for cryptocurrencies and metals.

More details about the API can be found on the [official API documentation](https://developers.bitpanda.com/platform/).

## Installation

```bash
go get github.com/sguter90/bitpanda-public-api-client
```

## Usage

### Example

```go
package main

import (
	"fmt"
	bitpanda "github.com/sguter90/bitpanda-public-api-client"
)

func main() {
	apiKey := "insert-your-api-key-here"
	conf := bitpanda.NewConfig(apiKey)
	c := bitpanda.NewClient(conf)

	resp, err := c.WalletsGet()

	fmt.Println(resp, err)
}
```

### Supported Endpoints
| URL                                  | Func name                     | Description                        |
| ------------------------------------ | ----------------------------- | ---------------------------------- |
| `GET /trades`                        | TradesGet                     | Get all trades                     |
| `GET /wallets`                       | WalletsGet                    | Get all user crypto wallets        |
| `GET /wallets/transactions`          | WalletTransactionGet          | Get all user's crypto transactions |
| `GET /asset-wallets`                 | AssetWalletsGet               | Get user's wallets based on assets |
| `GET /assets/transactions/commodity` | AssetTransactionsCommodityGet | Get user's commodity transactions  |
| `GET /fiatwallets`                   | FiatWalletsGet                | Get all user fiat wallets          |
| `GET: /fiatwallets/transactions`     | FiatWalletsTransactionsGet    | Get all user fiat transactions     |
 
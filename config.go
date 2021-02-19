package client

import "net/url"

const DEFAULT_USER_AGENT string = "bitpanda-public-api-go-client"
const BASE_URL_LIVE string = "https://api.bitpanda.com/v1"

type Config struct {
	ApiKey    string
	BaseUrl   *url.URL
	UserAgent string
}

func NewConfig(apiKey string) *Config {
	appUrl, _ := url.Parse(BASE_URL_LIVE)
	config := &Config{
		ApiKey:    apiKey,
		BaseUrl:   appUrl,
		UserAgent: DEFAULT_USER_AGENT,
	}

	return config
}

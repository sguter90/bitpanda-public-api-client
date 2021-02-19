package client

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

type Client struct {
	Config     *Config
	HttpClient *http.Client
}

type ParamsPage struct {
	Page     int64 `json:"page"`
	PageSize int64 `json:"page_size"`
}

func (p *ParamsPage) ToQueryParams() map[string]string {
	params := map[string]string{}
	if p.Page > 0 {
		params["page"] = strconv.FormatInt(p.Page, 10)
	}
	if p.PageSize > 0 {
		params["page_size"] = strconv.FormatInt(p.PageSize, 10)
	}

	return params
}

type ResponseError struct {
	Errors []struct {
		Status int64  `json:"status"`
		Code   string `json:"code"`
		Title  string `json:"title"`
	} `json:"errors"`
}

type ResponseMeta struct {
	TotalCount int64 `json:"total_count"`
	Page       int64 `json:"page"`
	PageSize   int64 `json:"page_size"`
}

type ResponseLinks struct {
	Next string `json:"next"`
	Last string `json:"last"`
	Self string `json:"self"`
}

func NewClient(config *Config) *Client {
	hClient := http.Client{}
	c := &Client{
		Config:     config,
		HttpClient: &hClient,
	}

	return c
}

func (c *Client) newRequest(method, path string, queryParams map[string]string, body interface{}) (*http.Request, error) {
	query := c.Config.BaseUrl.Query()
	for param, value := range queryParams {
		query.Set(param, value)
	}

	rel := &url.URL{
		Path:     c.Config.BaseUrl.Path + path,
		RawQuery: query.Encode(),
	}
	u := c.Config.BaseUrl.ResolveReference(rel)

	var buf io.ReadWriter
	if body != nil {
		buf = new(bytes.Buffer)
		err := json.NewEncoder(buf).Encode(body)
		if err != nil {
			return nil, err
		}
	}

	req, err := http.NewRequest(method, u.String(), buf)
	if err != nil {
		return nil, err
	}

	if c.Config.ApiKey != "" {
		req.Header.Set("X-API-KEY", c.Config.ApiKey)
	}

	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}

	req.Header.Set("Accept", "application/json")
	req.Header.Set("User-Agent", c.Config.UserAgent)

	return req, nil
}

func (c *Client) newGetRequest(path string, queryParams map[string]string) (*http.Request, error) {
	return c.newRequest("GET", path, queryParams, nil)
}

func (c *Client) do(req *http.Request, v interface{}) (ResponseError, *http.Response, error) {
	errResp := ResponseError{}
	httpResp, err := c.HttpClient.Do(req)
	if err != nil {
		return errResp, nil, err
	}
	defer httpResp.Body.Close()

	respBytes, err := ioutil.ReadAll(httpResp.Body)
	if err != nil {
		return errResp, nil, err
	}

	err = json.Unmarshal(respBytes, &errResp)
	if err != nil {
		return errResp, nil, err
	}

	if len(errResp.Errors) > 0 {
		msgs := []string{}
		for _, errorRow := range errResp.Errors {
			msgs = append(msgs, strconv.FormatInt(errorRow.Status, 10)+" - "+errorRow.Code+" - "+errorRow.Title)
		}

		err = errors.New(strings.Join(msgs, ", "))
		return errResp, httpResp, err
	}

	//fmt.Println((string(respBytes)))

	err = json.Unmarshal(respBytes, &v)

	return errResp, httpResp, err
}

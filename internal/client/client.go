package client

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"reflect"

	"github.com/prs-watch/ygocards/types"
)

// API Client.
type Client struct {
	URL    *url.URL
	Params types.Params
}

// Clientのinitialize定義.
func CreateClient(p types.Params, u string) Client {
	up, _ := url.Parse(u)
	return Client{
		URL:    up,
		Params: p,
	}
}

// Client側でHTTPリクエストを実行.
func (c *Client) Run() (types.Response, error) {
	q := c.URL.Query()
	ts := reflect.TypeOf(c.Params)
	vs := reflect.ValueOf(c.Params)

	// URLパラメータ作成.
	for i := 0; i < vs.NumField(); i++ {
		value := vs.Field(i).String()
		if value != "" {
			key := toLowerSnake(ts.Field(i).Name)
			q.Set(key, value)
		}
	}

	// HTTPリクエスト.
	c.URL.RawQuery = q.Encode()
	res, err := http.Get(c.URL.String())
	if err != nil {
		return types.Response{}, fmt.Errorf("error: %v", err)
	}
	defer res.Body.Close()

	// レスポンス処理.
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return types.Response{}, fmt.Errorf("error: %v", err)
	}
	var response types.Response
	err = json.Unmarshal(body, &response)
	// API側でエラーが返却された場合はUnmarshalに失敗するため、返却されたエラーメッセージを返す.
	if err != nil {
		return types.Response{}, fmt.Errorf("error: %v", err)
	}
	if response.Error != "" {
		return types.Response{}, fmt.Errorf("error: %v", response.Error)
	}

	return response, nil
}

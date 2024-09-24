package ygocards

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"reflect"
)

// API Client.
type Client struct {
	URL    *url.URL
	Params Params
}

// APIにパスするパラメータを定義.
type Params struct {
	Name       string
	Fname      string
	Id         string
	Konamiid   string
	Cardtype   string
	Atk        string
	Def        string
	Level      string
	Race       string
	Attribute  string
	Link       string
	Linkmarker string
	Scale      string
	Cardset    string
	Archetype  string
	Banlist    string
	Sort       string
	Format     string
	HasEffect  string
	Staple     string
	Startdate  string
	Enddate    string
	Dateregion string
	Language   string
}

// Cardの定義.
type Card struct {
	Id                    int         `json:"id"`
	Name                  string      `json:"name"`
	Type                  string      `json:"type"`
	HumanReadableCardType string      `json:"humanReadableCardType"`
	FrameType             string      `json:"frameType"`
	Desc                  string      `json:"desc"`
	Race                  string      `json:"race"`
	Atk                   int         `json:"atk,omitempty"`
	Def                   int         `json:"def,omitempty"`
	Level                 int         `json:"level,omitempty"`
	Attribute             string      `json:"attribute,omitempty"`
	Linkval               int         `json:"linkval,omitempty"`
	Linkmarkers           []string    `json:"linkmarkers,omitempty"`
	YgoprodeckURL         string      `json:"ygoprodeck_url"`
	CardSets              []CardSet   `json:"card_sets"`
	CardImages            []CardImage `json:"card_images"`
	CardPrices            []CardPrice `json:"card_prices"`
}

// CardSetの定義.
type CardSet struct {
	SetName       string `json:"set_name"`
	SetCode       string `json:"set_code"`
	SetRarity     string `json:"set_rarity"`
	SetRarityCode string `json:"set_rarity_code"`
	SetPrice      string `json:"set_price"`
}

// CardImageの定義.
type CardImage struct {
	Id              int    `json:"id"`
	ImageURL        string `json:"image_url"`
	ImageURLSmall   string `json:"image_url_small"`
	ImageURLCropped string `json:"image_url_cropped"`
}

// CardPriceの定義.
type CardPrice struct {
	CardmarketPrice   string `json:"cardmarket_price"`
	TcgplayerPrice    string `json:"tcgplayer_price"`
	EbayPrice         string `json:"ebay_price"`
	AmazonPrice       string `json:"amazon_price"`
	CoolstuffincPrice string `json:"coolstuffinc_price"`
}

// Clientのレスポンス定義.
type Response struct {
	Data  []Card
	Error string `json:"error"`
}

// Paramのinitialize定義.
func CreateParams() Params {
	return Params{}
}

// Clientのinitialize定義.
func CreateClient(p Params, u string) Client {
	up, _ := url.Parse(u)
	return Client{
		URL:    up,
		Params: p,
	}
}

// Client側でHTTPリクエストを実行.
func (c *Client) Run() (Response, error) {
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
		return Response{}, fmt.Errorf("error: %v", err)
	}
	defer res.Body.Close()

	// レスポンス処理.
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return Response{}, fmt.Errorf("error: %v", err)
	}
	var response Response
	err = json.Unmarshal(body, &response)
	// API側でエラーが返却された場合はUnmarshalに失敗するため、返却されたエラーメッセージを返す.
	if err != nil {
		return Response{}, fmt.Errorf("error: %v", err)
	}
	if response.Error != "" {
		return Response{}, fmt.Errorf("error: %v", response.Error)
	}

	return response, nil
}

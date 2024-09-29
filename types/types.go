package types

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
	Data  []Card `json:"data"`
	Error string `json:"error"`
}

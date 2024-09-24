package ygocards

// APIエンドポイント.
const URL = "https://db.ygoprodeck.com/api/v7/cardinfo.php"

// API Clientを生成し、search処理を実行する.
func Run(p Params) (Response, error) {
	c := createClient(p, URL)
	return c.run()
}

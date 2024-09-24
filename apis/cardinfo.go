package apis

import "github.com/prs-watch/ygocards"

// APIエンドポイント.
const URL = "https://db.ygoprodeck.com/api/v7/cardinfo.php"

// API Clientを生成し、search処理を実行する.
func Run(p ygocards.Params) (ygocards.Response, error) {
	c := ygocards.CreateClient(p, URL)
	return c.Run()
}

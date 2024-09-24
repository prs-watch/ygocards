package apis

import (
	"github.com/prs-watch/ygocards/internal/client"
	"github.com/prs-watch/ygocards/types"
)

// APIエンドポイント.
const URL = "https://db.ygoprodeck.com/api/v7/cardinfo.php"

// API Clientを生成し、search処理を実行する.
func GetCardInfo(p types.Params) (types.Response, error) {
	c := client.CreateClient(p, URL)
	return c.Run()
}

package apis

import (
	"github.com/prs-watch/ygocards/internal/client"
)

// Paramsを生成するexposeしたエンドポイント.
func CreateParams() client.Params {
	return client.CreateParams()
}

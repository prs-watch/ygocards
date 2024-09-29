package apis

import (
	"github.com/prs-watch/ygocards/internal/client"
	"github.com/prs-watch/ygocards/types"
)

func GetRandomCard() (types.Response, error) {
	p := CreateParams()
	c := client.CreateClient(p, "https://db.ygoprodeck.com/api/v7/randomcard.php")
	return c.Run()
}

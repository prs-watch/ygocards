package apis

import (
	"fmt"
	"testing"

	"github.com/prs-watch/ygocards/internal/client"
)

func TestRun(t *testing.T) {
	type args struct {
		p client.Params
	}
	tests := []struct {
		name string
		args args
	}{
		{"Search existing card", args{client.Params{Name: "Dark Magician"}}},
		{"Search non-existing card", args{client.Params{Name: "Dark Magicians"}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v, err := GetCardInfo(tt.args.p)
			if err != nil {
				fmt.Println(err.Error())
			} else {
				fmt.Println(v)
			}
		})
	}
}

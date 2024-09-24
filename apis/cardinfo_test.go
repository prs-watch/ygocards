package apis

import (
	"fmt"
	"testing"

	"github.com/prs-watch/ygocards"
)

func TestRun(t *testing.T) {
	type args struct {
		p ygocards.Params
	}
	tests := []struct {
		name string
		args args
	}{
		{"Search existing card", args{ygocards.Params{Name: "Dark Magician"}}},
		{"Search non-existing card", args{ygocards.Params{Name: "Dark Magicians"}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v, err := Run(tt.args.p)
			if err != nil {
				fmt.Println(err.Error())
			} else {
				fmt.Println(v)
			}
		})
	}
}

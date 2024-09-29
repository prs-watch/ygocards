package apis

import (
	"fmt"
	"testing"

	"github.com/prs-watch/ygocards/types"
)

func TestGetCardInfo(t *testing.T) {
	type args struct {
		p types.Params
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name:    "Request to existing card",
			args:    args{p: types.Params{Name: "Tornado Dragon"}},
			want:    "Tornado Dragon",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetCardInfo(tt.args.p)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetCardInfo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got.Data[0].Name != tt.want {
				t.Errorf("GetCardInfo() = %v, want %v", got, tt.want)
			}
			// 正解データ確認用のprintln.
			fmt.Println(got)
		})
	}
}

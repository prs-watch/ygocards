package apis

import (
	"fmt"
	"testing"
)

func TestGetRandomCard(t *testing.T) {
	tests := []struct {
		name    string
		wantErr bool
	}{
		{
			name:    "Request random card.",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetRandomCard()
			if (err != nil) != tt.wantErr {
				t.Errorf("GetRandomCard() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			// 正解データ確認用のprintln.
			fmt.Println(got)
		})
	}
}

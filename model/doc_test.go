package model

import (
	"github.com/space-backend/config"
	"testing"
)

func TestBumpDocVersion(t *testing.T) {
	type args struct {
		sid int64
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "1",
			args: args{
				sid: 363640270979141,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := BumpDocVersion(config.DB, tt.args.sid); (err != nil) != tt.wantErr {
				t.Errorf("BumpDocVersion() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

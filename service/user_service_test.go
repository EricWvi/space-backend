package service

import (
	"github.com/space-backend/config"
	"testing"
)

func init() {
	config.InitForTest()
}

func TestSign(t *testing.T) {
	type args struct {
		id uint
	}
	tests := []struct {
		name            string
		args            args
		wantTokenString string
		wantErr         bool
	}{
		{
			name: "1",
			args: args{
				id: 1010,
			},
			wantTokenString: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MTAxMH0.l686fXi12bFMtVMMXjLwfSTj8Rsoix0s8EmT1KSdGFo",
			wantErr:         false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotTokenString, err := Sign(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("Sign() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotTokenString != tt.wantTokenString {
				t.Errorf("Sign() gotTokenString = %v, want %v", gotTokenString, tt.wantTokenString)
			}
		})
	}
}

func TestParseToken(t *testing.T) {
	type args struct {
		tokenString string
	}
	tests := []struct {
		name    string
		args    args
		wantId  uint
		wantErr bool
	}{
		{
			name: "1",
			args: args{
				tokenString: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MTAxMH0.l686fXi12bFMtVMMXjLwfSTj8Rsoix0s8EmT1KSdGFo",
			},
			wantId:  1010,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotId, err := ParseToken(tt.args.tokenString)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotId != tt.wantId {
				t.Errorf("ParseToken() gotId = %v, want %v", gotId, tt.wantId)
			}
		})
	}
}

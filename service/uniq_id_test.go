package service

import (
	"fmt"
	"testing"
)

func TestToSid(t *testing.T) {
	type args struct {
		id int64
	}
	tests := []struct {
		name    string
		args    args
		wantSid string
		wantErr bool
	}{
		{
			name: "1",
			args: args{
				id: 363565572022341,
			},
			wantSid: "fpulh3kvz9",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotSid, err := ToSid(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("ToSid() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotSid != tt.wantSid {
				t.Errorf("ToSid() gotSid = %v, want %v", gotSid, tt.wantSid)
			}
		})
	}
}

func TestZx(t *testing.T) {
	tokenString, _ := Sign(1010)
	fmt.Println(tokenString)
}

func TestParseSid(t *testing.T) {
	type args struct {
		sid string
	}
	tests := []struct {
		name    string
		args    args
		wantId  int64
		wantErr bool
	}{
		{
			name: "1",
			args: args{
				sid: "fpulh3kvz9",
			},
			wantId:  363565572022341,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotId, err := ParseSid(tt.args.sid)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseSid() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotId != tt.wantId {
				t.Errorf("ParseSid() gotId = %v, want %v", gotId, tt.wantId)
			}
		})
	}
}

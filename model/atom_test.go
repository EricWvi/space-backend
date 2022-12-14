package model

import (
	"github.com/space-backend/config"
	"testing"
)

func init() {
	config.InitForTest()
}

func TestGetAtomsByDocId(t *testing.T) {
	type args struct {
		docId int64
	}
	tests := []struct {
		name      string
		args      args
		wantAtoms []Atom
		wantErr   bool
	}{
		{
			name: "1",
			args: args{
				docId: 363565572022341,
			},
			wantAtoms: nil,
			wantErr:   false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := GetAtomViewsByDocId(tt.args.docId)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAtomsByDocId() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

package model

import (
	"fmt"
	"github.com/space-backend/config"
	"testing"
)

func init() {
	config.InitForTest()
}

func TestTx(t *testing.T) {
	var docs []AtomField
	config.DB.Table("atoms").Where(map[string]any{
		"sid": 363581145104453,
	}).Find(&docs)
	fmt.Println(docs)
}

func TestGetAtom(t *testing.T) {
	a, _ := GetAtom(config.DB, map[string]any{Atom_Sid: 363581145104453})
	fmt.Printf("%#v\n", a.AtomField)
}

func TestGetAtomsByDocId(t *testing.T) {
	type args struct {
		docId Sid
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
			views, _ := GetAtomViewsByDoc(config.DB, tt.args.docId, 3)
			for _, v := range views {
				fmt.Printf("%#v\n", *v)
			}
		})
	}
}

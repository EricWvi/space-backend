package model

import (
	"fmt"
	"github.com/space-backend/config"
	"testing"
)

func init() {
	config.InitForTest()
}

func TestGetCollectionViews(t *testing.T) {
	collections, _ := GetCollectionViews(config.DB)
	fmt.Println(collections[0].Sid)
	fmt.Println(collections[0].CollectionField.Sid)
}

func TestSlice(t *testing.T) {
	type Ts struct {
		A int
	}
	l := []Ts{
		{
			A: 1,
		},
		{
			A: 2,
		},
	}
	var x []*Ts
	x = append(x, &l[0])
	fmt.Printf("%p\n", &l[0])
	fmt.Printf("%p\n", x[0])
}

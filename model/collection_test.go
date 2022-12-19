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

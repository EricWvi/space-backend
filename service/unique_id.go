package service

import (
	"github.com/space-backend/model"
	"github.com/yitter/idgenerator-go/idgen"
)

func init() {
	options := idgen.NewIdGeneratorOptions(1)
	idgen.SetIdGenerator(options)
}

func NextId() model.Sid {
	return model.Sid(idgen.NextId())
}

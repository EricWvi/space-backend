package service

import (
	"errors"
	"github.com/yitter/idgenerator-go/idgen"
	"strconv"
)

func init() {
	options := idgen.NewIdGeneratorOptions(1)
	idgen.SetIdGenerator(options)
}

func NextId() int64 {
	return idgen.NextId()
}

func ParseSid(sid string) (id int64, err error) {
	// kurb1 3kr ol
	// 3krkurb1ol
	if len(sid) < 10 {
		err = errors.New("sid is invalid")
		return
	}
	a := sid[5:8]
	a += sid[0:5]
	a += sid[8:]
	return strconv.ParseInt(a, 36, 64)
}

func ToSid(id int64) (sid string, err error) {
	// 3krkurb1ol
	// kurb1 3kr ol
	a := strconv.FormatInt(id, 36)
	if len(a) < 10 {
		err = errors.New("sid is too short")
		return
	}
	sid = a[3:8]
	sid += a[0:3]
	sid += a[8:]
	return
}

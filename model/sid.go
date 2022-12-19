package model

import (
	"errors"
	"strconv"
	"strings"
)

type Sid int64

func (id *Sid) UnmarshalJSON(data []byte) error {
	raw := strings.Trim(string(data), "\"")
	i, err := ParseSid(raw)
	*id = i
	return err
}

func ParseSid(sid string) (id Sid, err error) {
	if sid == "" {
		id = 0
		return
	}
	// kurb1 3kr ol
	// 3krkurb1ol
	if len(sid) < 10 {
		err = errors.New("sid is invalid")
		return
	}
	a := sid[5:8]
	a += sid[0:5]
	a += sid[8:]
	i, err := strconv.ParseInt(a, 36, 64)
	return Sid(i), err
}

func (id *Sid) String() string {
	if *id == 0 {
		return ""
	}
	// 3krkurb1ol
	// kurb1 3kr ol
	a := strconv.FormatInt(int64(*id), 36)
	if len(a) < 10 {
		return ""
	}
	sid := a[3:8]
	sid += a[0:3]
	sid += a[8:]
	return sid
}

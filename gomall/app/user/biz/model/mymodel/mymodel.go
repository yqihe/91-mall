package mymodel

import (
	"strings"
	"time"
)

type MyTime time.Time

func (t *MyTime) UnmarshalJSON(data []byte) error {
	s := strings.Trim(string(data), `"`)
	tmp, err := time.Parse(time.DateTime, s)
	if err != nil {
		return err
	}
	*t = MyTime(tmp)
	return nil
}

func (t MyTime) MarshalJSON() ([]byte, error) {
	timeStr := time.Time(t).Format(time.DateTime)
	return []byte(`"` + timeStr + `"`), nil
}

package mymodel

import "time"

// MyTime is a custom time type that wraps time.Time
type MyTime time.Time

// String formats MyTime as a string (for example, "2006-01-02 15:04:05")
func (t MyTime) String() string {
	return time.Time(t).Format("2006-01-02 15:04:05")
}

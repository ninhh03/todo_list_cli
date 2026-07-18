package main

import (
	"fmt"
	"time"
)

func parseTimeString (tStr string) (time.Time, error) {
	n := time.Now()
	year := n.Year()
	month := n.Month()
	day := n.Day()
	location := n.Location()

	t, err := time.ParseInLocation("15:04", tStr, location)
	if err != nil {
		return time.Time{}, err
	}
	hour := t.Hour()
	minute := t.Minute()
	second := t.Second()
	nanosecond := t.Nanosecond()

	fullTime := time.Date(year, month, day, hour, minute, second, nanosecond, location)
	
	return fullTime, nil
}

func main() {
	fmt.Println("Hello, World!")
}
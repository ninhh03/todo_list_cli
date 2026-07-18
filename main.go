package main

import (
	"encoding/json"
	"fmt"
	"os"
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

func writeFile(taskList []Task) error {
	n := time.Now()
	filePath := fmt.Sprintf("data/%s.json", n.Format(time.DateOnly))

	file, err := os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0600)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "    ")
	err = encoder.Encode(taskList)
	if err != nil {
		return err
	}

	return nil
}

func readFile(filePath string) ([]Task, error) {
	var taskList []Task

	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	err = json.NewDecoder(file).Decode(&taskList)
	if err != nil {
		return nil, err
	}

	return taskList, nil
}

func main() {
	fmt.Println("Hello, World!")
}
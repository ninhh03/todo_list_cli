package main

import (
	"encoding/json"
	"fmt"
	"os"
	"text/tabwriter"
	"time"
)

func parseTimeString(tStr string) (time.Time, error) {
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

func convertDateFormat(dStr string) (string, error) {
	d, err := time.Parse("02/01/2006", dStr)
	if err != nil {
		return "", err
	}

	return d.Format(time.DateOnly), nil
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

func displayTaskList(taskList []Task) {
	w := tabwriter.NewWriter(os.Stdout, 4, 0, 4, ' ', 0)

	fmt.Fprintln(w, "ID\tName\tStatus\tPriority\tStart time\tEnd time")
	fmt.Fprintln(w, "----\t------------\t------------\t------------\t------------\t------------")

	if len(taskList) == 0 {
		fmt.Fprintln(w, "Task list is empty!")
		w.Flush()
		return
	}

	for _, task := range taskList {
		startTimeStr := fmt.Sprintf("%s - %s", task.StartTime.Format("15:04"), task.StartTime.Format("02/01/2006"))
		endTimeStr := fmt.Sprintf("%s - %s", task.EndTime.Format("15:04"), task.EndTime.Format("02/01/2006"))
		fmt.Fprintf(w, "%s\t%s\t%s\t%s\t%s\t%s\n", task.ID, task.Name, task.Status, task.Priority, startTimeStr, endTimeStr)
	}

	w.Flush()
}

func displayTaskListByDate(dStr string) error {
	fileName, err := convertDateFormat(dStr)
	if err != nil {
		return err
	}

	filePath := fmt.Sprintf("data/%s.json", fileName)

	taskList, err := readFile(filePath)
	if err != nil {
		return err
	}
	displayTaskList(taskList)
	return nil
}

func main() {
	err := displayTaskListByDate("18/07/2026")
	if err != nil {
		fmt.Println("Error: ", err)
	}
}
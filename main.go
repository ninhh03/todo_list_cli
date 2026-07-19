package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
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

func generateID(ordinalNumber int) (string, error) {
	return fmt.Sprintf("t%02d", ordinalNumber), nil
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

func displayTaskListByToday() error {
	n := time.Now()
	today := n.Format(time.DateOnly)

	filePath := fmt.Sprintf("data/%s.json",today)

	taskList, err := readFile(filePath)
	if err != nil {
		return err
	}

	displayTaskList(taskList)
	return nil
}

func inputFromKeyboard() (string, string, time.Time, time.Time, error) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print(">📝 Task name: ")
	name, err := reader.ReadString('\n')
	if err != nil {
		return "", "", time.Time{}, time.Time{}, err
	}
	name = strings.TrimSpace(name)

	fmt.Print(">⭐ Priority (High, Medium, Low): ")
	priority, err := reader.ReadString('\n')
	if err != nil {
		return "", "", time.Time{}, time.Time{}, err
	}
	priority = strings.TrimSpace(priority)

	fmt.Print(">⏰ Start time (hh:mm): ")
	startTimeStr, err := reader.ReadString('\n')
	if err != nil {
		return "", "", time.Time{}, time.Time{}, err
	}
	startTimeStr = strings.TrimSpace(startTimeStr)
	startTime, err := parseTimeString(startTimeStr)
	if err != nil {
		return "", "", time.Time{}, time.Time{}, err
	}

	fmt.Print(">⏰ End time (hh:mm): ")
	endTimeStr, err := reader.ReadString('\n')
	if err != nil {
		return "", "", time.Time{}, time.Time{}, err
	}
	endTimeStr = strings.TrimSpace(endTimeStr)
	endTime, err := parseTimeString(endTimeStr)
	if err != nil {
		return "", "", time.Time{}, time.Time{}, err
	}

	return name, priority, startTime, endTime, nil
}

func createTask(ordinalNumber int, name string, priority string, startTime time.Time, endTime time.Time) (Task, error) {
	id, err := generateID(ordinalNumber)
	if err != nil {
		return Task{}, err
	}
	if name == "" {
		return Task{}, fmt.Errorf("task name cannot be empty!")
	}
	status := "To-do"
	if priority == "" {
		return Task{}, fmt.Errorf("priority cannot be empty!")
	}
	if(!endTime.After(startTime)) {
		return Task{}, fmt.Errorf("end time must be after start time!")
	}
	
	newTask := Task{
		ID: id,
		Name: name,
		Status: status,
		Priority: priority,
		StartTime: startTime,
		EndTime: endTime,
	}

	return newTask, nil
}

func main() {
}
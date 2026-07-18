package main

import (
	"time"
)

type Task struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Status    string    `json:"status"`
	Priority  string    `json:"priority"`
	StartTime time.Time `json:"start_time"`
	EndTime   time.Time `json:"end_time"`
}

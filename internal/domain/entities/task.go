package entities

import "time"

type Task struct {
	Date    time.Time
	Title   string
	Comment string
	Repeat  RepeatRule
}

package entities

import "time"

type RepeatRule interface {
	NextExecution(current time.Time) (time.Time, error)
	String() string
}

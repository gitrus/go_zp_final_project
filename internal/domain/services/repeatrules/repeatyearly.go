package repeatrules

import (
	"time"
)

type YearlyRule struct{}

func (r *YearlyRule) NextExecution(current time.Time) (time.Time, error) {
	return current.AddDate(1, 0, 0), nil
}

func (r *YearlyRule) String() string {
	return "y"
}

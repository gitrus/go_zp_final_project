package repeatrules

import (
	"fmt"
	"time"
)

type DailyRule struct {
	Days int // d 1, d 7
}

func (r *DailyRule) NextExecution(current time.Time) (time.Time, error) {

	if r.Days <= 0 || r.Days > 400 {
		return time.Time{}, fmt.Errorf("invalid daily rule: %d days", r.Days)
	}
	return current.AddDate(0, 0, r.Days), nil
}

func (r *DailyRule) String() string {
	return fmt.Sprintf("d %d", r.Days)
}

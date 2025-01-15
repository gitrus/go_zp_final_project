package repeatrules

import (
	"fmt"
	"go_zp_final_project/internal/domain/entities"
	"strconv"
	"strings"
)

func ParseRule(rule string) (entities.RepeatRule, error) {
	parts := strings.Split(rule, " ")
	switch parts[0] {
	case "d":
		if len(parts) != 2 {
			return nil, fmt.Errorf("invalid daily rule format")
		}
		days, err := strconv.Atoi(parts[1])
		if err != nil {
			return nil, fmt.Errorf("invalid daily rule: %v", err)
		}
		return &DailyRule{Days: days}, nil
	case "y":
		return &YearlyRule{}, nil
	default:
		return nil, fmt.Errorf("unknown rule type: %s", parts[0])
	}
}

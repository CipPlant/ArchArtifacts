package builder

import (
	"fmt"
	"strings"
)

const (
	centurySymbols = "XVI"
)

func consistCenturyRune(description string) string {
	for _, symbol := range centurySymbols {
		if strings.ContainsRune(description, symbol) {
			return fmt.Sprintf("%s", description)
		}
	}
	return ""
}

func BuildingAge(description string, interval1, interval2 int16) string {
	switch {
	case interval1 == 0:
		if description == "" {
			return fmt.Sprintf("%d Год", interval2)
		}
		res := consistCenturyRune(description)
		if res == "" {
			return fmt.Sprintf("%s %d Год", description, interval2)
		}
		return res

	case interval2 == 0:
		if description == "" {
			return fmt.Sprintf("%d Год", interval1)
		}
		res := consistCenturyRune(description)
		if res == "" {
			return fmt.Sprintf("%s %d Год", description, interval1)
		}
		return res

	default:
		if description == "" {
			return fmt.Sprintf("%d - %d Года", interval1, interval2)
		}
		res := consistCenturyRune(description)
		if res == "" {
			return fmt.Sprintf("%s %d - %d Года", description, interval1, interval2)
		}
		return res
	}
}

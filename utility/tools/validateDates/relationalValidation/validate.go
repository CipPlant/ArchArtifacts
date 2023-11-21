package relationalValidation

import "errors"

func ValidateCentury(s string) ([]int, error) {
	switch s {
	case "start":
		return []int{1, 2, 3}, nil
	case "middle":
		return []int{4, 5, 6, 7}, nil
	case "end":
		return []int{8, 9, 10}, nil
	case "firstMiddle":
		return []int{1, 2, 3, 4, 5}, nil
	case "secondMiddle":
		return []int{6, 7, 8, 9, 10}, nil
	default:
		return nil, errors.New("error in validation")
	}
}

func ValidateDecade(s string) ([]int, error) {
	switch s {
	case "start":
		return []int{0, 1, 2}, nil
	case "middle":
		return []int{3, 4, 5, 6}, nil
	case "end":
		return []int{7, 8, 9}, nil
	case "firstMiddle":
		return []int{0, 1, 2, 3, 4}, nil
	case "secondMiddle":
		return []int{5, 6, 7, 8, 9}, nil
	default:
		return nil, errors.New("error in validation")
	}
}

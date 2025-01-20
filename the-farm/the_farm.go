package thefarm

import "fmt"

func DivideFood(fc FodderCalculator, cows int) (float64, error) {
	if fodder, err := fc.FodderAmount(cows); err != nil {
		return 0, err
	} else if fatteningFactor, err := fc.FatteningFactor(); err != nil {
		return 0, err
	} else {
		return fodder * fatteningFactor / float64(cows), nil
	}
}

func ValidateInputAndDivideFood(fc FodderCalculator, cows int) (float64, error) {
	if cows <= 0 {
		return 0, fmt.Errorf("invalid number of cows")
	}
	return DivideFood(fc, cows)
}

func ValidateNumberOfCows(cows int) error {
	if cows < 0 {
		return &InvalidCowsError{numberOfCows: cows, message: "there are no negative cows"}
	} else if cows == 0 {
		return &InvalidCowsError{numberOfCows: cows, message: "no cows don't need food"}
	} else {
		return nil
	}
}

type InvalidCowsError struct {
	numberOfCows int
	message      string
}

func (e *InvalidCowsError) Error() string {
	return fmt.Sprintf("%d cows are invalid: %s", e.numberOfCows, e.message)
}

package thefarm

import (
	"errors"
	"fmt"
)

// See types.go for the types defined for this exercise.

// TODO: Define the SillyNephewError type here.
type SillyNephewError struct {
	cows int
}

func (e *SillyNephewError) Error() string {
	return fmt.Sprintf("silly nephew, there cannot be %d cows", e.cows)
}

// DivideFood computes the fodder amount per cow for the given cows.
func DivideFood(weightFodder WeightFodder, cows int) (float64, error) {
	fodderAmount, err := weightFodder.FodderAmount()
	if err == ErrScaleMalfunction && fodderAmount > 0 {
		return (fodderAmount * 2) / float64(cows), nil
	} else if fodderAmount < 0 {
		return 0.0, errors.New("Negative fodder")
	} else if err != nil {
		return 0, err
	} else if cows == 0 {
		return 0.0, errors.New("Division by zero")
	} else if cows < 0 {
		return 0.0, &SillyNephewError{cows: cows}
	}

	return fodderAmount / float64(cows), err
}

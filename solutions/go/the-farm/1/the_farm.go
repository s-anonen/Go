package thefarm
import (
    "fmt"
    "errors"
       )
// TODO: define the 'DivideFood' function
func DivideFood(fc FodderCalculator, count int) (float64, error) {
    total, err := fc.FodderAmount(count)
    if err != nil {
        return 0, err
    }
    coef, err := fc.FatteningFactor()
    if err != nil {
        return 0, err
    }
    
    return (total/float64(count) * coef), nil
}

// TODO: define the 'ValidateInputAndDivideFood' function
func ValidateInputAndDivideFood(fc FodderCalculator, count int) (float64, error) {
    if count > 0 {
        return DivideFood(fc, count)
    }
    
    return 0, errors.New("invalid number of cows")
}

// TODO: define the 'ValidateNumberOfCows' function
type InvalidCowsError struct {
    count int
    str string
}

func (obj *InvalidCowsError) Error() string {
    return fmt.Sprintf("%d cows are invalid: %s", obj.count, obj.str)
}

func ValidateNumberOfCows(count int) error {
    if count < 0 {
        return &InvalidCowsError{count, "there are no negative cows"}
    } else if count == 0 {
        return &InvalidCowsError{count, "no cows don't need food"}
    } else {
        return nil
    }
}
// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.
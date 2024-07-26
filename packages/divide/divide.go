package divide

import "errors"

func DivideBy(x, y float64) (float64, error) {
	// This function is used by Divide, notice that it didn't have a capitalized name
	// The capitalization of this function is due to this function only belongs in another function and might not be used outside of this document

	if y == 0 {
		return 0, errors.New("can not divide by zero")
	}
	result := x / y
	return result, nil
}

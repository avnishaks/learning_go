package functions

import (
	"errors"
	"fmt"
)

func Function() {
	var a int8 = 10
	fmt.Println(a)
}

func IntDivFunc(num int, dem int) (int, int, error) {
	var err error
	if dem == 0 {
		err = errors.New("Cannot divide by zero")
		return 0, 0, err
	}
	var result int = num / dem
	fmt.Println("Result form int division function", result)
	var remainder int = num % dem
	return result, remainder, err
}

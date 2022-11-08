package main

import (
	"errors"
	"fmt"
)

func divide(q int, d int) (int, error) {
	if d == 0 {
		return 0, errors.New("divided by zero not possible")
	}

	return q / d, nil
}

func main() {
	fmt.Println(divide(10, 0))
}

package main

import "fmt"

func divide(q, d int) int {
	fmt.Println("Dividing it now")
	return q / d
}

func main() {
	fmt.Println("the division is:", divide(4, 0))
}

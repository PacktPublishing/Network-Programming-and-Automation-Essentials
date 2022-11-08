package main

import "fmt"

func divide(q, d int) int {
	fmt.Println("Dividing it now")
	return q / d
}

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Got a panic:", r)
		}
	}()
	fmt.Println("the division is:", divide(4, 0))

}

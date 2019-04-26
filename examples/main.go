package main

import "fmt"

func main() {
	val := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	for _, i := range val {
		if i%2 == 0 {
			fmt.Printf("%d is even\n", i)
		} else {
			fmt.Printf("%d is odd\n", i)
		}
	}
}

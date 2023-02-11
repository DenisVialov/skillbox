package main

import (
	"fmt"
)

func swap(x [10]int) (b [10]int) {
	i := 9
	for _, num := range x {
		b[i] = num
		i--
	}
	return
}

func main() {
	a := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(swap(a))
}

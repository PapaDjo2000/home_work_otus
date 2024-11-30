package main

import (
	"fmt"

	search "github.com/PapaDjo2000/hw08_binary_search/pkg"
)

func main() {
	mass := []int{1, 2, 4, 5, 7, 8}
	target := 5

	result := search.BinarySearch(target, mass)

	if result == -1 {
		fmt.Println("обьект не найден")
	} else {
		fmt.Printf("обьект %d находиться по индексу %d\n", target, result)
	}
}

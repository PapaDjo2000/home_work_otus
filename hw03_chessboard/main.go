package main

import "fmt"

func main() {
	var size_x int
	var size_y int

	fmt.Println("Введите Х")
	fmt.Scanf("%d", &size_x)

	fmt.Println("Введите У")
	fmt.Scanf("%d", &size_y)

	for i := 0; i < size_y; i++ {
		if i%2 == 0 {
			for i := 0; i < size_x; i++ {
				if i%2 == 0 {
					fmt.Print("#")
				} else {
					fmt.Print(" ")
				}
				if i == size_x-1 {
					fmt.Println("")
				}
			}
		} else {
			for i := 0; i < size_x; i++ {
				if i%2 == 0 {
					fmt.Print(" ")
				} else {
					fmt.Print("#")
				}
				if i == size_x-1 {
					fmt.Println("")
				}
			}
		}
	}
}

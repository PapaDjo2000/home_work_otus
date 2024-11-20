package main

import "fmt"

func X(index int, sizeX int) {
	if index%2 == 0 {
		fmt.Print(" ")
	} else {
		fmt.Print("#")
	}
	if index == sizeX-1 {
		fmt.Println("")
	}
}

func main() {

	var sizeX int
	var sizeY int

	fmt.Println("Введите Х")
	fmt.Scanf("%d", &sizeX)

	fmt.Println("Введите У")
	fmt.Scanf("%d", &sizeY)

	for i := 0; i < sizeY; i++ {
		if i%2 == 0 {
			for index := 0; index < sizeX; index++ {
				X(index, sizeX)
			}
		} else {
			for index := 1; index < sizeX; index++ {
				X(index, sizeX)
			}
		}
	}
}

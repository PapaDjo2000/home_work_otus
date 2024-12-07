package packageX

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

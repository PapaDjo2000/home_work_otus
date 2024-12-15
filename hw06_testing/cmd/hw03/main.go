package main

import (
	"fmt"

	packageX "github.com/PapaDjo2000/home_work_otus/hw06_testing/pkg/hw03"
)

func main() {
	var x, y int

	fmt.Println("Введите Х")
	fmt.Scanf("%d", &x)

	fmt.Println("Введите У")
	fmt.Scanf("%d", &y)
	fmt.Print(packageX.Chess(y, x))
}

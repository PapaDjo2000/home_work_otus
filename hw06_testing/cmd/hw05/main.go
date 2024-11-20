package main

import (
	"fmt"

	"github.com/PapaDjo2000/home_work_otus/hw06_testing/pkg/hw05/calculate"
	"github.com/PapaDjo2000/home_work_otus/hw06_testing/pkg/hw05/figura"
	"github.com/PapaDjo2000/home_work_otus/hw06_testing/pkg/hw05/shape"
)

func main() {
	var typeErr struct{}
	var circle shape.Shape = &figura.Circle{
		Radius: 5,
	}
	var rectangle shape.Shape = &figura.Rectangle{
		Wigth:  10,
		Height: 5,
	}
	var triangle shape.Shape = &figura.Triangle{
		Basis:  8,
		Height: 6,
	}

	pl, err := calculete.CalculateArea(circle)
	if err != nil {
		fmt.Print(err)
	} else {
		fmt.Printf("Площадь:%v\n", pl)
	}
	pl, err = calculete.CalculateArea(rectangle)
	if err != nil {
		fmt.Print(err)
	} else {
		fmt.Printf("Площадь:%v\n", pl)
	}
	_, err = calculete.CalculateArea(triangle)
	if err != nil {
		fmt.Print(err)
	} else {
		fmt.Printf("Площадь:%v\n", pl)
	}
	pl, err = calculete.CalculateArea(typeErr)
	if err != nil {
		fmt.Print(err)
	} else {
		fmt.Printf("Площадь:%v\n", pl)
	}
}

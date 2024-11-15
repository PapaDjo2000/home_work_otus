package main

import (
	"errors"
	"fmt"

	"github.com/PapaDjo2000/home_work_otus/hw05_shapes/figura"
	"github.com/PapaDjo2000/home_work_otus/hw05_shapes/shape"
)

func calculateArea(s any) (float64, error) {
	switch shape := s.(type) {
	case shape.Shape:
		return shape.Area(), nil
	default:
		err := errors.New("ошибка:переданный объект не является фигурой")
		return 0, err
	}
}

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

	pl, err := calculateArea(circle)
	if err != nil {
		fmt.Print(err)
	} else {
		fmt.Printf("Площадь:%v\n", pl)
	}
	pl, err = calculateArea(rectangle)
	if err != nil {
		fmt.Print(err)
	} else {
		fmt.Printf("Площадь:%v\n", pl)
	}
	_, err = calculateArea(triangle)
	if err != nil {
		fmt.Print(err)
	} else {
		fmt.Printf("Площадь:%v\n", pl)
	}
	pl, err = calculateArea(typeErr)
	if err != nil {
		fmt.Print(err)
	} else {
		fmt.Printf("Площадь:%v\n", pl)
	}
}

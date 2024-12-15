package calculete

import (
	"errors"

	"github.com/PapaDjo2000/home_work_otus/hw06_testing/pkg/hw05/shape"
)

func CalculateArea(s any) (float64, error) {
	switch shape1 := s.(type) {
	case shape.Shape:
		return shape1.Area(), nil
	default:
		err := errors.New("ошибка:переданный объект не является фигурой")
		return 0, err
	}
}

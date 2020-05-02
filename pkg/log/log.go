package log

import (
	"errors"
	"fmt"

	color "github.com/daviddengcn/go-colortext"
	colorfmt "github.com/daviddengcn/go-colortext/fmt"
)

func Status(bright bool, str string) {
	colorfmt.Println(color.Green, bright, str)
}

func Statusf(bright bool, str string, args ...interface{}) {
	formatted := fmt.Sprintf(str, args...)
	Status(bright, formatted)
}

func Error(bright bool, str string) error {
	err := errors.New(str)
	colorfmt.Println(color.Red, bright, str)
	return err
}

func Err(bright bool, err error) error {
	return Error(bright, err.Error())
}

func Errorf(bright bool, str string, args ...interface{}) error {
	formatted := fmt.Sprintf(str, args...)
	return Error(bright, formatted)
}

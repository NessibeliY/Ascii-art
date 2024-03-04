package internal

import (
	"errors"
	"fmt"
)

const height = 8

func Run() error {
	input, err := getInput()
	if err != nil {
		return err
	}

	if input == "" {
		return nil
	}

	ok, err := ValidInput(input)
	if !ok {
		return err
	}

	result, err := Convert(input, "standard.txt")
	if err != nil {
		return err
	}

	gotWidth, err := getWidth()
	if err != nil {
		return errors.New("there is no terminal to get width")
	}

	if width(result) > int(gotWidth) {
		return errors.New("please provide shorter text")
	}

	fmt.Print(result)

	return nil
}

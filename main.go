package main

import (
	"errors"
	"fmt"
)

func main() {
	basicErr := errors.New("This is a test error")
	fmt.Println(basicErr)

	a, b := 10, 0

	_, err := Divide(a, b)

	if err != nil {
		switch {
		case errors.Is(err, ErrDivideByZero):
			fmt.Printf("divide by zero error: %s\n", ErrDivideByZero)
		default:
			fmt.Printf("unexpected division error: %s\n", err)
		}
	}

	err = doWork()

	if err != nil {
		fmt.Println(err)
		var specificError *WorkError
		if errors.As(err, &specificError) {
			fmt.Println("it was a specific error")
		}

	}
}

var ErrDivideByZero = errors.New("divide by zero")

func Divide(a, b int) (int, error) {
	if b == 0 {
		return 0, ErrDivideByZero
	}
	return a / b, nil
}

func doWork() error {
	err := doMoreWork()

	if err != nil {
		return fmt.Errorf("doMoreWork: %w", err)
	}

	return nil
}

func doMoreWork() error {
	err := doTheMostWork()

	if err != nil {
		return fmt.Errorf("doMoreWork: %w", err)
	}

	return nil
}

func doTheMostWork() error {
	return &WorkError{
		SomeInt:    10,
		SomeString: "Sam",
		Msg:        "doTheMostWork: something went wrong",
	}
}

type WorkError struct {
	SomeInt    int
	SomeString string
	Msg        string
}

func (e *WorkError) Error() string {
	return e.Msg
}

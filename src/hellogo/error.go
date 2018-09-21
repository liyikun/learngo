package main

import (
	"time"
	"fmt"
)

type MyError struct {
	When time.Time
	What string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("at %v,%s", e.When, e.What)
}

func run() error {
	return &MyError{
		time.Now(),
		"it is not work",
	}
}


func Sqrt(x float64) (float64, error) {
	return 0, nil
}

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	if e < 0 {
		return fmt.Sprintf( "cannot Sqrt negative number: %v\n",float64(e))
	}
	return ""
}


func main()  {
	if err := run(); err != nil {
		fmt.Println(err)
	}


	newerr := ErrNegativeSqrt(2)
	fmt.Println(newerr)

	newerr = ErrNegativeSqrt(-2)
	fmt.Println(newerr)
}
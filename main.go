package main

import (
	"fmt"

	"github.com/pkg/errors"
)

func bar() error {
	err := foo()
	return err
}

func foo() error {
	return errors.New("foo")
}

func main() {
	err := bar()
	fmt.Printf("%+v", err)
}

package main

import (
	"fmt"

	"github.com/pkg/errors"
)

type Error string

func (e Error) Error() string {
	return string(e)
}

const e1 = Error("WTF")
const e2 = Error("WTF")

func main() {
	e3 := errors.Wrap(e1, "ololo")
	fmt.Println(e1 == e2)
	fmt.Println(e1 == e3)
	fmt.Println(errors.Cause(e3) == e1)
}

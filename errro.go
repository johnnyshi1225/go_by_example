package main

import (
	. "fmt"
	"errors"
)

func double(arg int) (int, error) {
	if arg % 2 != 0 {
		return -1, errors.New("I don't like odd!")
	}

	return arg * 2, nil
}

// Custom Error
type myError struct {
	code int
	msg string
}

func (e *myError) Error() string {
	return Sprintf("ERROR_CODE[%d], ERROR_MESSGE[%s]", e.code, e.msg)
}

func main() {
	for _, v := range []int{1, 2, 3, 4, 5, 6, 7, 8, 9} {
		Println("Current:", v)
		if result, err := double(v); err != nil {
			Println("Error:", err)
		} else {
			Println("Result:", result)
		}
		Println()
	}

	me := &myError{code: 404, msg: "Page not found."}
	Println(me)
}

package main

import "fmt"
import "errors"

func wontAcceptTheAnswer(value int) (int, error) {
	if value == 42 {
		return -1, errors.New("Won't accept 42")
	}

	return value * 2, nil
}

type argumentError struct {
	argument int
	problem string
}

func (e *argumentError) Error() string {
	return fmt.Sprintf("%d - %s", e.argument, e.problem)
}

func shallNotAcceptTheAnswer(value int) (int, error) {
	if value == 42 {
		return -1, &argumentError{value, "shall not work with it. <sniffs>"}
	}

	return value * 4, nil
}

func main() {
	for _, number := range []int{6, 42} {
		if returnValue, errorReturn := wontAcceptTheAnswer(number); errorReturn != nil {
			fmt.Println("wontAcceptTheAnswer failed:", errorReturn)
		} else {
			fmt.Println("wontAcceptTheAnswer succeeded:", returnValue)
		}
	}

	for _, number := range []int{6, 42} {
		if returnValue, errorReturn := shallNotAcceptTheAnswer(number); errorReturn != nil {
			fmt.Println("shallNotAcceptTheAnswer failed:", errorReturn)
		} else {
			fmt.Println("shallNotAcceptTheAnswer succeeded:", returnValue)
		}
	}

	_, errorReturn := shallNotAcceptTheAnswer(42)
	if errorInstance, ok := errorReturn.(*argumentError); ok {
		fmt.Println(errorInstance.argument)
		fmt.Println(errorInstance.problem)
	}
}

package main

import "fmt"

//def structs
type MyError struct {
	Code    int
	Message string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("Error %d: %s", e.Code, e.Message)
}

func Add(number1 float64, number2 float64) float64 {
	return number1 + number2
}

func Multiple(number1 float64, number2 float64) float64 {
	return number1 * number2
}

func Subtraction(number1 float64, number2 float64) float64 {
	return number1 - number2
}

func Divide(number1 float64, number2 float64) (float64, error) {
	if number2 == 0 {
		return 0, &MyError{Code: 400, Message: "Divide by zero"}
	}
	return number1 / number2, nil
}

func main() {
	fmt.Println(Add(3, 5))
	fmt.Println(Divide(5, 0))
}

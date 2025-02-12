package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/expr-lang/expr"
)

func evaluate(expression string) (interface{}, error) {
	env := map[string]interface{}{}
	program, err := expr.Compile(expression, expr.Env(env))
	if err != nil {
		return nil, err
	}
	output, err := expr.Run(program, env)
	if err != nil {
		return nil, err
	}
	return output, nil
}

func calculate(luku1, luku2 float64, op string) {
	expression := fmt.Sprintf("%f %s %f", luku1, op, luku2)
	result, _ := evaluate(expression)
	fmt.Println("Lasken:", luku1, op, luku2)
	fmt.Println("Tulos on", result)
}

func readInput[T int | float64 | string](msg string) T {
	fmt.Print(msg)
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	var result interface{}
	var t T
	switch any(t).(type) {
	case int:
		result, _ = strconv.Atoi(input)
	case float64:
		result, _ = strconv.ParseFloat(input, 64)
	case string:
		result = input
	}
	return result.(T)
}

func sendValue[T int | float64 | string](value T, channel chan T) {
	channel <- value
}

func main() {

	chan1 := make(chan float64)
	chan2 := make(chan float64)
	chan_op := make(chan string)

	value1 := readInput[float64]("Anna ensimmäinen luku: ")
	value2 := readInput[float64]("Anna toinen luku: ")
	value3 := readInput[string]("Anna laskutoimitus: ")
	go sendValue(value1, chan1)
	go sendValue(value2, chan2)
	go sendValue(value3, chan_op)

	luku1 := <-chan1
	luku2 := <-chan2
	operaatio := <-chan_op
	calculate(luku1, luku2, operaatio)
}

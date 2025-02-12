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

func readInput[T int | float64 | string](msg string) T {
	fmt.Print(msg)
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSuffix(input, "\n")
	input = strings.TrimSuffix(input, "\r")
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

func calculate(luku1, luku2 float64, op string) {
	expression := fmt.Sprintf("%f %s %f", luku1, op, luku2)
	result, _ := evaluate(expression)
	fmt.Println("Lasken: ", luku1, op, luku2)
	fmt.Println("Tulos on", result)
}

func main() {
	luku1 := readInput[float64]("Anna ensimmäinen luku: ")
	luku2 := readInput[float64]("Anna toinen luku: ")
	operaatio := readInput[string]("Anna laskutoimitus: ")
	calculate(luku1, luku2, operaatio)
}

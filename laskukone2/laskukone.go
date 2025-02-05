package main

import (
	"fmt"
	"os"

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

func readNumber(msg string) float64 {
	var number float64
	fmt.Print(msg)
	_, err := fmt.Scanln(&number)

	if err != nil {
		println("Virheellinen syöte.")
		os.Exit(0)
	}
	return number
}

func readOperation(msg string) string {
	var operation string
	fmt.Print(msg)
	fmt.Scan(&operation)
	return operation
}

func main() {
	luku1 := readNumber("Anna ensimmäinen luku: ")
	luku2 := readNumber("Anna toinen luku: ")
	operaatio := readOperation("Anna laskutoimitus: ")
	expression := fmt.Sprintf("%f %s %f", luku1, operaatio, luku2)
	result, _ := evaluate(expression)
	fmt.Println("Lasken: ", luku1, operaatio, luku2)
	fmt.Println("Tulos on", result)
}

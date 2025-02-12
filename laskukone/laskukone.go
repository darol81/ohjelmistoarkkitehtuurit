package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
)

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

func laske(luku1, luku2 float64, operaatio string) {
	var tulos float64
	var virhe string
	// Lasketaan laskutoimitus.

	fmt.Println("Lasken: ", luku1, operaatio, luku2)

	switch operaatio {
	case "+":
		tulos = luku1 + luku2
	case "-":
		tulos = luku1 - luku2
	case "*":
		tulos = luku1 * luku2
	case "/":
		if luku2 == 0 {
			fmt.Println("Nollalla ei voi jakaa.")
		} else {
			tulos = luku1 / luku2
		}
	default:
		tulos = luku1 + rand.Float64()*luku2
		virhe = "En tunnistanut laskutoimitusta '" + operaatio + "'. " +
			"Tässä sinulle satunnaisluku: " +
			strconv.FormatFloat(tulos, 'G', 3, 64)
	}
	// Tulostetaan laskutoimituksen tulos tai virhe.
	if virhe == "" {
		fmt.Println("Tulos on", tulos)
	} else {
		fmt.Println("Virhe:", virhe)
	}
}

func main() {
	// Kysy ensimmäinen luku.
	luku1 := readNumber("Anna ensimmäinen luku: ")
	luku2 := readNumber("Anna toinen luku: ")
	// Kysy laskutoimitus.
	operaatio := readOperation("Anna laskutoimitus: ")
	laske(luku1, luku2, operaatio)
}

package main

import (
	"fmt"
	"math/rand"
	"strconv"
)

func readNumber(msg string) float64 {
	var number float64
	fmt.Print(msg)
	fmt.Scan(&number)
	return number
}

func readOperation(msg string) string {
	var operation string
	fmt.Print(msg)
	fmt.Scan(&operation)
	return operation
}

func main() {
	// Laskettavat luvut
	var luku1, luku2 float64
	// Laskutoimitus
	var operaatio string

	// Kysy ensimmäinen luku.
	luku1 = readNumber("Anna ensimmäinen luku: ")
	luku2 = readNumber("Anna toinen luku: ")

	// Kysy laskutoimitus.
	operaatio = readOperation("Anna laskutoimitus: ")

	// Tulostetaan, mitä luettiin.
	fmt.Println("Lasken: ", luku1, operaatio, luku2)

	// Tulos sisältää laskutoimituksen tuloksen.
	var tulos float64

	// Virhe sisältää kuvauksen virheestä, mutta on tyhjä, jos virhettä ei tapahtunut.
	var virhe string

	// Suoritetaan laskutoimituksen mukainen laskutoimitus.
	switch operaatio {
	case "+":
		tulos = luku1 + luku2
		virhe = ""
	case "-":
		tulos = luku1 - luku2
		virhe = ""

	case "*":
		tulos = luku1 * luku2
		virhe = ""

	case "/":
		tulos = luku1 / luku2
		virhe = ""

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

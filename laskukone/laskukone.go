package main

import (
	"fmt"
	"math/rand"
	"strconv"
)

func main() {
	// Laskettavat luvut
	var luku1, luku2 float64

	// Laskutoimitus
	var lt string

	// Kysy ensimmäinen luku.
	fmt.Println("Anna ensimmäinen luku:")
	fmt.Scanf("%f", &luku1)

	// Kysy toinen luku.
	fmt.Println("Anna toinen luku:")
	fmt.Scanf("%f", &luku2)

	// Kysy laskutoimitus.
	fmt.Println("Anna laskutoimitus:")
	fmt.Scanf("%s", &lt)

	// Tulostetaan, mitä luettiin.
	fmt.Println("Lasken: ", luku1, lt, luku2)

	// Tulos sisältää laskutoimituksen tuloksen.
	var tulos float64

	// Virhe sisältää kuvauksen virheestä, mutta on tyhjä, jos virhettä ei tapahtunut.
	var virhe string

	// Suoritetaan laskutoimituksen mukainen laskutoimitus.
	switch lt {
	case "+":
		tulos = luku1 + luku2
		virhe = ""
	default:
		tulos = luku1 + rand.Float64()*luku2
		virhe = "En tunnistanut laskutoimitusta '"+lt+"'. "+
			"Tässä sinulle satunnaisluku: "+
			strconv.FormatFloat(tulos, 'G', 3, 64)
	}

	// Tulostetaan laskutoimituksen tulos tai virhe.
	if virhe=="" {
		fmt.Println("Tulos on", tulos)
	} else {
		fmt.Println("Virhe:", virhe)
	}
}

 
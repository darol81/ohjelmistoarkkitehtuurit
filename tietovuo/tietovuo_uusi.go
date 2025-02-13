
// Pseudokoodina. Esimerkki uuden tietovuon luomisesta.
// Pakettia käsitellään ja ulostuloa käytetään sisääntulona seuraavalle operaatiolle.
package main	

import (
	"fmt"
)

type Paketti struct {
	// Paketin sisältö
}

int main() {	
	paketti := Paketti{}
	paketti = handle_operation1(paketti)
	paketti = handle_operation2(paketti)
}
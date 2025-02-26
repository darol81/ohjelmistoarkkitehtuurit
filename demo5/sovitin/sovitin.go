/*

Oih, eihän Acme isona firmana tietenkään itse mitään tee, vaan hankkii toteutuksen alihankintana Pekan Pittipajalta. Ongelmana
 vain on, ettei PP tee mitään samalla rajapinnalla kuin Acmelta vaaditaan, vaan heillä on tarjolla operaatiot:

    setConnectionOpen(state boolean), joka vastaa edellisen tehtävän rajapinnan yhteyden avausta (state = true) ja	sulkemista (state = false), sekä
    communicate(data *[]byte, readFlag boolean), joka täyttää data-parametrin luettavilla tavuilla, jos readFlag = true, tai kirjoittaa sen arvon
	verkkoon, jos readFlag = false.

Autapa Acmea suunnittelemalla ja toteuttamalla Sovitin (Adapter, kurssin oppien mukaan) Acmen tarjoaman ja
PP:n tarjoaman rajapinnan välille.

*/

package main

// Määritellään Tietoverkko-interface
type Tietoverkko interface {
	AvaaYhteys(kohde string)
	SuljeYhteys(kohde string)
	LahetaDataa(kohde string, data string)
	VastaanotaDataa(kohde string) string
}

type PekanPittipaja struct {
}

func (p *PekanPittipaja) setConnectionOpen(state bool) {
	if state {
		println("PekanPittipaja: Avataan yhteys")
	} else {
		println("PekanPittipaja: Suljetaan yhteys")
	}
}

func (p *PekanPittipaja) communicate(data *[]byte, readFlag bool) {
	if readFlag {
		println("PekanPittipaja: Luetaan dataa")
	} else {
		println("PekanPittipaja: Kirjoitetaan dataa: " + string(*data))
	}
}

// Sovitin Acme ja PekanPittipaja välille
type Sovitin struct {
	PP PekanPittipaja
}

func (s *Sovitin) AvaaYhteys(kohde string) {
	s.PP.setConnectionOpen(true)
}

func (s *Sovitin) SuljeYhteys(kohde string) {
	s.PP.setConnectionOpen(false)
}

func (s *Sovitin) LahetaDataa(kohde string, data string) {
	d := []byte(data)
	s.PP.communicate(&d, false)
}

func (s *Sovitin) VastaanotaDataa(kohde string) string {
	var d []byte
	s.PP.communicate(&d, true)
	return string(d)
}

// Määritellään Acme-toteutus
type Acme struct {
}

func (a *Acme) AvaaYhteys(kohde string) {
	println("Acme: Avataan yhteys kohteeseen: ", kohde)
}

func (a *Acme) SuljeYhteys(kohde string) {
	println("Acme: Suljetaan yhteys kohteeseen: ", kohde)
}

func (a *Acme) LahetaDataa(kohde string, data string) {
	println("Acme: Lähetetään dataa kohteeseen: ", kohde, "Data: ", data)
}

func (a *Acme) VastaanotaDataa(kohde string) string {
	println("Acme: Vastaanotetaan dataa kohteesta: ", kohde)
	return "Dataa"
}

type Kayttojarjestelma struct {
	Tietoverkko Tietoverkko
}

func (k *Kayttojarjestelma) AsetaTietoverkko(t Tietoverkko) {
	k.Tietoverkko = t
}

func main() {
	// Luodaan käyttöjärjestelmä
	kayttis := Kayttojarjestelma{}
	verkko := Sovitin{}
	kayttis.AsetaTietoverkko(&verkko)
	kayttis.Tietoverkko.AvaaYhteys("www.google.com")
	kayttis.Tietoverkko.LahetaDataa("www.google.com", "Dataa tässä menee paljon...")
	kayttis.Tietoverkko.VastaanotaDataa("www.google.com")
	kayttis.Tietoverkko.SuljeYhteys("www.google.com")
}

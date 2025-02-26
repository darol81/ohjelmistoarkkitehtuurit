/*

Käyttöjärjestelmässä tulee olla yhteys Tietoverkkoon. Yhteyden toteutus voi tulla firmalta Acme tai Bonc, mutta huhuja on myös
kolmannesta firmasta. Käyttöjärjestelmä tarvitsee yhteydeltä seuraavat toiminnot: avaa yhteys kohteeseen, sulje yhteys
kohteeseen, lähetä dataa ja vastaanota dataa. Käyttöjärjestelmän ei tule olla tietoinen käytettävästä yhteyden toteutuksesta.

    Suunnittele ja kuvaa UML-kaavioina yllä selitetty tilanne.
    Mielellään myös toteutat vastaavan suunnitelman valitsemallasi kielellä.

Käyttöjärjestelmää ei todellakaan tarvitse toteuttaa, ei edes tietoverkkoyhteyttä, vaan ainoastaan se,
kuinka tietoverkkoyhteyden toteutus voidaan antaa käyttöjärjestelmän käytettäväksi. Vinkki: rajapinnat.

*/

package main

// Määritellään Tietoverkko-interface
type Tietoverkko interface {
	AvaaYhteys(kohde string)
	SuljeYhteys(kohde string)
	LahetaDataa(kohde string, data string)
	VastaanotaDataa(kohde string) string
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
	verkko := Acme{}
	kayttis.AsetaTietoverkko(&verkko)
	kayttis.Tietoverkko.AvaaYhteys("www.google.com")
	kayttis.Tietoverkko.LahetaDataa("www.google.com", "Dataa tässä menee paljon...")
	kayttis.Tietoverkko.VastaanotaDataa("www.google.com")
	kayttis.Tietoverkko.SuljeYhteys("www.google.com")
}

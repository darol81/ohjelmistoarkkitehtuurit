/* Tiedostojärjestelmäsimulaattori, joka käyttää rekursiokoostetta */

package main

import (
	"fmt"
	"strings"
)

type Document struct {
	name     string
	children []Document
}

func (d *Document) printDocument(level int) {

	indent := strings.Repeat(" ", level*2)
	fmt.Println(indent + "+ " + d.name)
	for _, child := range d.children {
		child.printDocument(level + 1)
	}
}

func (d *Document) addDocument(child Document) {
	d.children = append(d.children, child)
}

func newDocument(name string) Document {
	return Document{name: name}
}

func main() {
	MainDoc := newDocument("Main Doc")
	Subdoc1 := newDocument("Subdoc 1")
	Subdoc2 := newDocument("Subdoc 2")

	child_a := newDocument("Child A")
	child_b := newDocument("Child B")
	child_c := newDocument("Child C")
	child_d := newDocument("Child D")

	Subdoc1.addDocument(child_a)
	Subdoc1.addDocument(child_b)
	Subdoc2.addDocument(child_c)
	Subdoc2.addDocument(child_d)

	MainDoc.addDocument(Subdoc1)
	MainDoc.addDocument(Subdoc2)

	MainDoc.printDocument(0)
}

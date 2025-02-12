package main

import (
        "fmt"
        "time"
)

// ---------------   Rajapinnat

type PowerSource interface {
        Start()
        Stop()
        SetUser(EngineUser)
}

type EngineUser interface {
        Warn(string)
}

// ---------------   Moottori

type engine struct {
        oilPressure int
        user        EngineUser
}

// Erikoinen määrittely, jotta Start yhdistyy engine-tietueeseen
// muodostaen tavallaan moduulin.

func (eng engine) Start() {
        fmt.Println("byrrrn!")
        eng.Run()
}

func (eng engine) Stop() {
        fmt.Println("köh-köh")
}

// Tähti enginen edessä kertoo, että metodi muuttaa engine-tietueen
// sisältämää tietoa (viiteparametri, ei arvoparametri).
func (eng *engine) SetUser(u EngineUser) {
        eng.user = u
        fmt.Println("  use a ok ")
}

// Run ei kuulu PowerSource-rajapintaan

func (eng engine) Run() {
        fmt.Println("...put...")
        if eng.oilPressure < 23 {
                eng.user.Warn("low oil pressure")
        }
        fmt.Println("...put...")
}

// Tuon edeltävän voisi määritellä näin
//
// func Run(eng engine) { .... }
//
// ja kutsua sitä myöhemmin näin
//
//     Run(e)
//
// jolloin ohjelma toimisi samoin.  Saman voisi
// tehdä muille, rajapintoihin kuulumattomille
// metodeille.

// ---------------   Auto

type Car struct {
        lp    string
        myEng PowerSource
}

func (c *Car) SetEngine(p PowerSource) {
        c.myEng = p
}

func (c Car) Warn(msg string) {
        fmt.Println(msg)
}

// ---------------   Pääohjelma

func main() {
        fmt.Println("hello, world")
        var c = Car{lp: "313"}
        var e = engine{oilPressure: 24}
        e.SetUser(&c)   // & kertoo, että viedään viite,
        c.SetEngine(&e) //           eikä kopiota arvosta
        e.Start()
        time.Sleep(2 * time.Second)
        e.Stop()
}
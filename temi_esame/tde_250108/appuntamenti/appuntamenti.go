package main

import (
	"fmt"
	"strings"
)

type Appuntamento struct {
	giorno             int
	oraInizio, oraFine int
}

func (a Appuntamento) String() (s string) {
	prima := strings.Repeat("-", a.oraInizio)
	app := strings.Repeat("*", a.oraFine-a.oraInizio)
	dopo := strings.Repeat("-", 24-a.oraFine)
	return fmt.Sprintf(
		"[il giorno %d dalle %02d alle %02d: %s%s%s]",
		a.giorno, a.oraInizio, a.oraFine, prima, app, dopo)
}

func main() {
	var gg, h1, h2 int
	agenda := make([]Appuntamento, 0)
	for {
		_, err := fmt.Scan(&gg, &h1, &h2)
		if err != nil {
			break
		}
		a, ok := NuovoAppuntamento(gg, h1, h2)
		if ok {
			AggiungiAppuntamento(a, &agenda)
		}
	}
	for _, ap := range agenda {
		fmt.Println(ap)
	}
	//fmt.Scan(&gg)
	//fmt.Println(AppuntamentiGiornata(gg, agenda))
}

func GiornoValido(gg int) bool {
	return 1 <= gg && gg <= 366
}

func OraValida(h int) bool {
	return 0 <= h && h <= 24
}

func NuovoAppuntamento(gg, h1, h2 int) (a Appuntamento, ok bool) {
	if GiornoValido(gg) && OraValida(h1) && OraValida(h2) && h1 <= h2 {
		a = Appuntamento{gg, h1, h2}
		ok = true
	} else {
		ok = false
	}
	return
}

func AggiungiAppuntamento(a Appuntamento, agenda *[]Appuntamento) bool {
	for _, app := range *agenda {
		if SovrapposizioneAppuntamenti(a, app) {
			return false
		}
	}
	*agenda = append(*agenda, a)
	return true
}

func SovrapposizioneAppuntamenti(a, b Appuntamento) bool {
	if a.giorno != b.giorno {
		return false
	}

	return a.oraInizio < b.oraFine &&
		b.oraInizio < a.oraFine
}

/*
func AppuntamentiGiornata(gg int, agenda []Appuntamento) (giornata []Appuntamento) {
	for _, a := range agenda {
		if a.giorno == gg {
			giornata = append(giornata, a)
		}
	}
	return
}
*/

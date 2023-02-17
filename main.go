package main

import (
	"bytes"
	"flag"
	"fmt"
	"strconv"

	"github.com/NorwegianKiwi-glitch/funtemps/conv"
)

// Definerer flag-variablene i hoved-"scope"
var fahr float64
var celsius float64
var kelvin float64

var out string
var Svar float64
var funfacts string
var funfactsunit string

// Bruker init (som anbefalt i dokumentasjonen) for å sikre at flagvariablene
// er initialisert.
func init() {

	// Definerer og initialiserer flagg-variablene
	flag.Float64Var(&fahr, "F", 0.0, "temperatur i grader fahrenheit")
	flag.Float64Var(&celsius, "C", 0.0, "temperatur i grader celsius")
	flag.Float64Var(&kelvin, "K", 0.0, "temperatur i grader kelvin")
	// Du må selv definere flag-variablene for "C" og "K"
	flag.StringVar(&out, "out", "C", "beregne temperatur i C - celsius, F - farhenheit, K- Kelvin")

	flag.StringVar(&funfacts, "funfacts", "sun", "\"fun-facts\" om sun - Solen, luna - Månen og terra - Jorden")
	flag.StringVar(&funfactsunit, "t", "C", "bruker setter input verdi, hvis ikke setter programmet til deafult verdi Celsius")

	// Du må selv definere flag-variabelen for -t flagget, som bestemmer
	// hvilken temperaturskala skal brukes når funfacts skal vises

}

func addSpaces(s string) string { // legger til mellomrom mellom hvert 3 siffer
	var buf bytes.Buffer
	n := len(s)
	for i, c := range s {
		buf.WriteRune(c)
		if i != n-1 && (n-i-1)%3 == 0 {
			buf.WriteRune(' ')
		}
	}
	return buf.String()
}

func main() {

	flag.Parse()

	// Her er noen eksempler du kan bruke i den manuelle testingen
	fmt.Println(fahr, out, funfacts)

	fmt.Println("len(flag.Args())", len(flag.Args()))
	fmt.Println("flag.NFlag()", flag.NFlag())

	fmt.Println(isFlagPassed("out"))

	Erlik := "="
	F := "°F"
	C := "°C"
	K := "°K"

	// FahrenheitToCelsius
	if out == "C" && isFlagPassed("F") {
		Svar = conv.FarhenheitToCelsius(fahr)

		fmt.Printf("%.12g %s %s ", fahr, F, Erlik)
		if Svar == float64(int(Svar)) {
			fmt.Printf("%s %s\n", addSpaces(strconv.Itoa(int(Svar))), C)
		} else {
			fmt.Printf("%s %s\n", addSpaces(strconv.FormatFloat(Svar, 'f', 2, 64)), C)
		}
	}

	// CelsiusToFahrenheit
	if out == "F" && isFlagPassed("C") {
		Svar = conv.CelsiusToFarenheit(celsius)

		fmt.Printf("%.12g %s %s ", celsius, C, Erlik)
		if Svar == float64(int(Svar)) {
			fmt.Printf("%s %s\n", addSpaces(strconv.Itoa(int(Svar))), F)
		} else {
			fmt.Printf("%s %s\n", addSpaces(strconv.FormatFloat(Svar, 'f', 2, 64)), F)
		}
	}

	// CelsiusToKelvin
	if out == "K" && isFlagPassed("C") {
		Svar = conv.CelsiusToKelvin(celsius)

		fmt.Printf("%.12g %s %s ", celsius, C, Erlik)
		if Svar == float64(int(Svar)) {
			fmt.Printf("%s %s\n", addSpaces(strconv.Itoa(int(Svar))), K)
		} else {
			fmt.Printf("%s %s\n", addSpaces(strconv.FormatFloat(Svar, 'f', 2, 64)), K)
		}
	}

	// KelvinToCelsius
	if out == "C" && isFlagPassed("K") {
		Svar = conv.KelvinToCelsius(kelvin)

		fmt.Printf("%.12g %s %s ", kelvin, K, Erlik)
		if Svar == float64(int(Svar)) {
			fmt.Printf("%s %s\n", addSpaces(strconv.Itoa(int(Svar))), C)
		} else {
			fmt.Printf("%s %s\n", addSpaces(strconv.FormatFloat(Svar, 'f', 2, 64)), C)
		}

	// FahrenheitToKelvin
	if out == "K" && isFlagPassed("F") {
		Svar = conv.FarhenheitToKelvin(fahr)

		fmt.Printf("%.12g %s %s ", fahr, F, Erlik)
		if Svar == float64(int(Svar)) {
			fmt.Printf("%s %s\n", addSpaces(strconv.Itoa(int(Svar))), K)
		} else {
			fmt.Printf("%s %s\n", addSpaces(strconv.FormatFloat(Svar, 'f', 2, 64)), K)
		}
	}

	// KelvinToFahrenheit
	if out == "F" && isFlagPassed("K") {
		Svar = conv.FarhenheitToCelsius(kelvin)

		fmt.Printf("%.12g %s %s ", kelvin, K, Erlik)
		if Svar == float64(int(Svar)) {
			fmt.Printf("%s %s\n", addSpaces(strconv.Itoa(int(Svar))), F)
		} else {
			fmt.Printf("%s %s\n", addSpaces(strconv.FormatFloat(Svar, 'f', 2, 64)), F)
		}
	}

}

// Funksjonen sjekker om flagget er spesifisert på kommandolinje
// Du trenger ikke å bruke den, men den kan hjelpe med logikken
func isFlagPassed(name string) bool {
	found := false
	flag.Visit(func(f *flag.Flag) {
		if f.Name == name {
			found = true
		}
	})
	return found
}

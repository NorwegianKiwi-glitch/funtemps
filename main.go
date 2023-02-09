package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"

	"golang.org/x/text/language"
	"golang.org/x/text/message"

	"github.com/NorwegianKiwi-glitch/funtemps/conv"
)

// Definerer flag-variablene i hoved-"scope"
var fahr string
var celsius string
var kelvin string

var fahrFloat float64
var celsiusFloat float64
var kelvinFloat float64

var out string
var funfacts string
var scale string

// Bruker init (som anbefalt i dokumentasjonen) for å sikre at flagvariablene
// er initialisert.
func init() {

	/*
	   Her er eksempler på hvordan man implementerer parsing av flagg.
	   For eksempel, kommando
	       funtemps -F 0 -out C
	   skal returnere output: 0°F er -17.78°C
	*/

	// Definerer og initialiserer flagg-variablene
	flag.StringVar(&fahr, "F", "0.0", "Temperatur i grader fahrenheit")
	flag.StringVar(&celsius, "C", "0.0", "Temperatur i grader celsius")
	flag.StringVar(&kelvin, "K", "0.0", "Temperatur i Kelvin")

	// Du må selv definere flag-variablene for "C" og "K"
	flag.StringVar(&out, "out", "", "Beregne temperatur i C - celsius, F - farhenheit, K- Kelvin")
	flag.StringVar(&funfacts, "funfacts", "", "\"fun-facts\" om sun - Solen, luna - Månen og terra - Jorden")
	// Du må selv definere flag-variabelen for -t flagget, som bestemmer
	// hvilken temperaturskala skal brukes når funfacts skal vises
	flag.StringVar(&scale, "t", "", "'Funfacts' temperaturskala i °C, °F, K.")
}

func main() {
	flag.Parse()

	// Making float versions of input. We need to keep original input as string in order to print it again as it was inputted. "Input skal alltid presenteres slik som input er på kommandolinje"
	celsiusFloat = toFloat(celsius)
	fahrFloat = toFloat(fahr)
	kelvinFloat = toFloat(kelvin)

	checkInvalidArgs() // Check for invalid arguments. Print and exit if exist

	// Are we going to show funfacts?
	if isFlagPassed("funfacts") || isFlagPassed("t") {
		switch {
		case isFlagPassed("funfacts") || !isFlagPassed("t"): // Check if another flag than 't' is passed
			fmt.Println("'funfacts' can only be used with 't'")
			os.Exit(1)
		case isFlagPassed("t") || !isFlagPassed("funfacts"): // Check if another flag than 'funfacts' is passed
			fmt.Println("'t' can only be used with 'funfacts'")
			os.Exit(1)
		default: // If both 't' and 'funfacts' are true
			fmt.Println("show some fun facts", funfacts, scale)

		}
		// no funfacts, just temperature conversion
	} else if isFlagPassed(out) { // Check to see if conversion is sane
		fmt.Println("Can not convert from", out, "to", out+".")
		os.Exit(1)
	} else {
		switch {
		case out == "C" && isFlagPassed("F"): // F to C
			fmt.Printf("%v°F er %v°C", fahr, formatOutput(conv.FarhenheitToCelsius(fahrFloat)))

		case out == "K" && isFlagPassed("F"): // F to K
			fmt.Printf("%v°F er %vK", fahr, formatOutput(conv.FahrenheitToKelvin(fahrFloat)))

		case out == "F" && isFlagPassed("C"): // C to F
			fmt.Printf("%v°C er %v°F", celsius, formatOutput(conv.CelsiusToFahrenheit(celsiusFloat)))

		case out == "K" && isFlagPassed("C"): // C to K
			fmt.Printf("%v°C er %vK", celsius, formatOutput(conv.CelsiusToKelvin(celsiusFloat)))

		case out == "C" && isFlagPassed("K"): // K to C
			fmt.Printf("%vK er %v°C", kelvin, formatOutput(conv.KelvinToCelsius(kelvinFloat)))

		case out == "F" && isFlagPassed("K"): // K to F
			fmt.Printf("%vK er %v°F", kelvin, formatOutput(conv.KelvinToFahrenheit(kelvinFloat)))

		default: // Catch if conditions are not met
			fmt.Println("Error")
			flag.PrintDefaults()
		}
	}

	/**
	    Her må logikken for flaggene og kall til funksjoner fra conv og funfacts
	    pakkene implementeres.

	    Det er anbefalt å sette opp en tabell med alle mulige kombinasjoner
	    av flagg. flag-pakken har funksjoner som man kan bruke for å teste
	    hvor mange flagg og argumenter er spesifisert på kommandolinje.

	        fmt.Println("len(flag.Args())", len(flag.Args()))
			    fmt.Println("flag.NFlag()", flag.NFlag())

	    Enkelte kombinasjoner skal ikke være gyldige og da må kontrollstrukturer
	    brukes for å utelukke ugyldige kombinasjoner:
	    -F, -C, -K kan ikke brukes samtidig
	    disse tre kan brukes med -out, men ikke med -funfacts
	    -funfacts kan brukes kun med -t
	    ...
	    Jobb deg gjennom alle tilfellene. Vær obs på at det er en del sjekk
	    implementert i flag-pakken og at den vil skrive ut "Usage" med
	    beskrivelsene av flagg-variablene, som angitt i parameter fire til
	    funksjonene Float64Var og StringVar
	*/

	// Her er noen eksempler du kan bruke i den manuelle testingen
	//	fmt.Println(fahr, out, funfacts)

	//	fmt.Println("len(flag.Args())", len(flag.Args()))
	//	fmt.Println("flag.NFlag()", flag.NFlag())

	//	fmt.Println(isFlagPassed("out"))

	// Eksempel på enkel logikk
	//	if out == "C" && isFlagPassed("F") {
	// Kalle opp funksjonen FahrenheitToCelsius(fahr), som da
	// skal returnere °C

	//		fmt.Println("0°F er °C", conv.FarhenheitToCelsius(fahr))
	//		fmt.Printf("%v°F er %.2f °C", fahr, conv.FarhenheitToCelsius(fahr))
	//	}

	fmt.Println() //newline
}

// Funksjonen sjekker om flagget er spesifisert på kommandolinje
func isFlagPassed(name string) bool {
	found := false
	flag.Visit(func(f *flag.Flag) {
		if f.Name == name {
			found = true
		}
	})
	return found
}

func checkInvalidArgs() {
	// Check for invalid arguments. Print and exit if exist
	if flag.NArg() > 0 {
		fmt.Println("Error")
		for unknownArgs := range flag.Args() {
			fmt.Println("'"+flag.Args()[unknownArgs]+"'", "is not recognised")
		}
		fmt.Println("Exiting.")
		os.Exit(1)
	}
}

func toFloat(input string) float64 { // From: https://www.tutorialkart.com/golang-tutorial/golang-convert-string-to-float/
	value, err := strconv.ParseFloat(input, 64)
	if err == nil { // If there was no error
		return value
	} else { // Tell user that we could not convert to float, input was probably not a number
		fmt.Println("'"+input+"'", "is not a number.")
		os.Exit(1)
	}

	return value
}

func formatOutput(input float64) string {
	p := message.NewPrinter(language.Ukrainian) // Source: https://gosamples.dev/print-number-thousands-separator/
	p.Printf("%f", input)
	formatted := strconv.FormatFloat(input, 'f', 2, 64)

	fmt.Println()

	return strings.TrimRight(strings.TrimRight(formatted, "0"), ".")
}

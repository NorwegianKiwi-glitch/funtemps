package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/NorwegianKiwi-glitch/funtemps/conv"
	"github.com/NorwegianKiwi-glitch/funtemps/funfacts"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: temperature [value] [unit]")
		os.Exit(1)
	}

	value, err := strconv.ParseFloat(os.Args[1], 64)
	if err != nil {
		fmt.Printf("Invalid value: %s\n", os.Args[1])
		os.Exit(1)
	}

	unit := os.Args[2]
	var result float64
	var unitName string

	switch unit {
	case "C", "c":
		result = conv.CelsiusToFarenheit(value)
		unitName = "Farenheit"
	case "F", "f":
		result = conv.FarenheitToCelsius(value)
		unitName = "Celsius"
	default:
		fmt.Printf("Invalid unit: %s\n", unit)
		os.Exit(1)
	}

	fmt.Printf("%.2f %s = %.2f %s\n", value, unit, result, unitName)

	facts := funfacts.GetFunFacts("Terra")
	fmt.Println("Fun facts about temperature on Earth:")
	for i, fact := range facts {
		fmt.Printf("%d. %s\n", i+1, fact)
	}
}

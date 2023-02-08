package funfacts

import (
	"reflect"
	"testing"
)

/*
*

	Mal for TestGetFunFacts funksjonen.
	Definer korrekte typer for input og want,
	og sette korrekte testverdier i slice tests.
*/
func TestGetFunFacts(t *testing.T) {
	type test struct {
		input string // her må du skrive riktig type for input
		want  string // her må du skrive riktig type for returverdien
	}

	// Her må du legge inn korrekte testverdier
	// tests := []test{
	// {input: , want: },
	//
	tests := []test{
		{input: "Sun", want: "Tempraturen i solens kjerne er 15 000 000C"},
		{input: "Luna", want: "Tempraturen på månens overflate om natten er -183C"},
		{input: "Terra", want: "Høyeste tempraturen på jordens overflate er 56,7C"},
	}

	for _, tc := range tests {
		got := GetFunFacts(tc.input)
		if !reflect.DeepEqual(tc.want, got) {
			t.Errorf("expected: %v, got: %v", tc.want, got)
		}
	}
}

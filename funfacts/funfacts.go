package funfacts

/**
  Implementer funfacts-funksjon:
    GetFunFacts(about string) []string
      hvor about kan ha en av tre testverdier, -
        sun, luna eller terra

  Sett inn alle Funfucts i en struktur
  type FunFacts struct {
      Sun []string
      Luna []string
      Terra []string
  }
*/

type funfacts struct {
	Sun   []string
	Luna  []string
	Terra []string
}

func GetFunFacts(about string) []string {
	var facts funfacts
	facts = funfacts{
		Sun: []string{
			"Tempraturen i solens kjerne er 15 000 000C*",
			"Tempraturen på ytre laget av Solen er 5 500C*",
		},
		Luna: []string{
			"Tempraturen på månens overflate om natten er -183C*",
			"Temrapturen på månens overflate om dagen er 379C*",
		},
		Terra: []string{
			"Høyeste tempraturen på jordens overflate er 56,7C*",
			"Laveste tempratur målt på jordens overlate er -89,9C*",
			"Temrpatur i jordens indre kjerne er 9 108C*",
		},
	}

	switch about {
	case "Sun":
		return facts.Sun
	case "Luna":
		return facts.Luna
	case "Terra":
		return facts.Terra
	default:
		return []string{"Ugyldig input. Venligst velg mellom sun, luna eller terra"}
	}
}

package russian

func rublesByCase(nnc numeralNumberCase) string {
	switch nnc {
	case singularNominative:
		return "рубль"
	case singularGenitive:
		return "рубля"
	case pluralGenitive:
		return "рублей"
	default:
		return ""
	}
}

// Rubles returns russian for "ruble" corresponding to 'amount'.
func Rubles(amount int64) string {
	return rublesByCase(getNumeralNumberCase(amount))
}

// NumRubles returns string containing number 'amount' and russian for "ruble" corresponding to 'amount'.
func NumRubles(amount int64, showZero bool) string {
	return numberAndRussianItems(amount, showZero, Rubles(amount))
}

func kopecksByCase(nnc numeralNumberCase) string {
	switch nnc {
	case singularNominative:
		return "копейка"
	case singularGenitive:
		return "копейки"
	case pluralGenitive:
		return "копеек"
	default:
		return ""
	}
}

// Kopecks returns russian for "kopeck" corresponding to 'amount'.
func Kopecks(amount int64) string {
	return kopecksByCase(getNumeralNumberCase(amount))
}

// NumKopecks returns string containing number 'amount' and russian for "kopeck" corresponding to 'amount'.
func NumKopecks(amount int64, showZero bool) string {
	return numberAndRussianItems(amount, showZero, Kopecks(amount))
}

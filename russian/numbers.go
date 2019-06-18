package russian

func thousandsByCase(nnc numeralNumberCase) string {
	switch nnc {
	case singularNominative:
		return "тысяча"
	case singularGenitive:
		return "тысячи"
	case pluralGenitive:
		return "тысяч"
	default:
		return ""
	}
}

// Thousands returns russian for "thousand" corresponding to 'i'.
func Thousands(i int64) string {
	return thousandsByCase(getNumeralNumberCase(i))
}

func millionsByCase(nnc numeralNumberCase) string {
	switch nnc {
	case singularNominative:
		return "миллион"
	case singularGenitive:
		return "миллиона"
	case pluralGenitive:
		return "миллионов"
	default:
		return ""
	}
}

// Millions returns russian for "million" corresponding to 'i'.
func Millions(i int64) string {
	return millionsByCase(getNumeralNumberCase(i))
}

func milliardsByCase(nnc numeralNumberCase) string {
	switch nnc {
	case singularNominative:
		return "миллиард"
	case singularGenitive:
		return "миллиарда"
	case pluralGenitive:
		return "миллиардов"
	default:
		return ""
	}
}

// Milliards returns russian for "milliard" corresponding to 'i'.
func Milliards(i int64) string {
	return milliardsByCase(getNumeralNumberCase(i))
}

// Billions returns russian for "milliard" corresponding to 'i'.
// (Billion in russian is called milliard.)
func Billions(i int64) string {
	return Milliards(i)
}

func trillionsByCase(nnc numeralNumberCase) string {
	switch nnc {
	case singularNominative:
		return "триллион"
	case singularGenitive:
		return "триллиона"
	case pluralGenitive:
		return "триллионов"
	default:
		return ""
	}
}

// Trillions returns russian for "trillion" corresponding to 'i'.
func Trillions(i int64) string {
	return trillionsByCase(getNumeralNumberCase(i))
}

func quadrillionsByCase(nnc numeralNumberCase) string {
	switch nnc {
	case singularNominative:
		return "квадриллион"
	case singularGenitive:
		return "квадриллиона"
	case pluralGenitive:
		return "квадриллионов"
	default:
		return ""
	}
}

// Quadrillions returns russian for "quadrillion" corresponding to 'i'.
func Quadrillions(i int64) string {
	return quadrillionsByCase(getNumeralNumberCase(i))
}

func quintillionsByCase(nnc numeralNumberCase) string {
	switch nnc {
	case singularNominative:
		return "квинтиллион"
	case singularGenitive:
		return "квинтиллиона"
	case pluralGenitive:
		return "квинтиллионов"
	default:
		return ""
	}
}

// Quintillions returns russian for "quintillion" corresponding to 'i'.
func Quintillions(i int64) string {
	return quintillionsByCase(getNumeralNumberCase(i))
}

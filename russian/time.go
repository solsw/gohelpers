package russian

func daysByCase(nnc numeralNumberCase) string {
	switch nnc {
	case singularNominative:
		return "день"
	case singularGenitive:
		return "дня"
	case pluralGenitive:
		return "дней"
	default:
		return ""
	}
}

// Days returns russian for 'day' corresponding to i.
func Days(i int64) string {
	return daysByCase(getNumeralNumberCase(i))
}

// NumDays returns string containing number i and russian for 'day' corresponding to i.
func NumDays(i int64, showZero bool) string {
	return numberAndRussianItems(i, showZero, Days(i))
}

func hoursByCase(nnc numeralNumberCase) string {
	switch nnc {
	case singularNominative:
		return "час"
	case singularGenitive:
		return "часа"
	case pluralGenitive:
		return "часов"
	default:
		return ""
	}
}

// Hours returns russian for 'hour' corresponding to i.
func Hours(i int64) string {
	return hoursByCase(getNumeralNumberCase(i))
}

// NumHours returns string containing number i and russian for 'hour' corresponding to i.
func NumHours(i int64, showZero bool) string {
	return numberAndRussianItems(i, showZero, Hours(i))
}

package russian

import (
	"fmt"

	"github.com/solsw/gohelpers/mathhelper"
)

type (
	// GrammaticalGender - russian grammatical gender.
	GrammaticalGender int
)

const (
	// Neuter gender.
	Neuter GrammaticalGender = iota
	// Masculine gender.
	Masculine
	// Feminine gender.
	Feminine
)

type (
	// русские число/падеж, используемые с числами
	// russian number/case used with numbers
	numeralNumberCase int
)

// numeralNumberCase choices
const (
	// singular, nominative case
	// единственное число, именительный падеж (1, 21 час (но 11 часов))
	singularNominative numeralNumberCase = iota
	// singular, genitive case
	// единственное число, родительный падеж (2, 3, 4, 22 часа (но 12, 13, 14 часов))
	singularGenitive
	// plural, genitive case
	// множественное число, родительный падеж (0, 5 (и всё остальное) часов)
	pluralGenitive
)

func getNumeralNumberCasePrim(last2 int64) numeralNumberCase {
	// 0 <= last2 <= 99
	// depends on two last digits in general
	// в общем случае определяется двумя последними цифрами
	if last2 == 11 || last2 == 12 || last2 == 13 || last2 == 14 {
		return pluralGenitive
	}
	// depends on one last digit now
	// теперь определяется одной последней цифрой
	switch last2 % 10 {
	case 1:
		return singularNominative
	case 2, 3, 4:
		return singularGenitive
	default:
		return pluralGenitive
	}
}

func getNumeralNumberCase(n int64) numeralNumberCase {
	return getNumeralNumberCasePrim(mathhelper.AbsInt(n) % 100)
}

// NumberAndItems returns string containing number 'n' and 'items'.
// If 'n' is 0 and 'showZero' is false, empty string is returned.
func NumberAndItems(n int64, showZero bool, items string) string {
	if n == 0 && !showZero {
		return ""
	}
	return fmt.Sprintf("%d %s", n, items)
}

// NumberInWordsAndItems returns string containing number 'n' in russian words and 'items'.
// If 'n' is 0 and 'showZero' is false, empty string is returned.
// If 'withZero' is false, zero triples will be omitted.
func NumberInWordsAndItems(n int64, showZero, withZero bool, gender GrammaticalGender, items string) string {
	if n == 0 && !showZero {
		return ""
	}
	return NumberInWords(n, withZero, gender) + " " + items
}

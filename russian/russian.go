// Package russian implements functions to manipulate russian words.
package russian

import (
	"fmt"

	"github.com/solsw/gohelpers/mathhelper"
)

// GrammaticalGender - russian grammatical gender.
type GrammaticalGender int

const (
	// Neuter gender.
	Neuter GrammaticalGender = iota
	// Masculine gender.
	Masculine
	// Feminine gender.
	Feminine
)

// numeralNumberCase - русские число/падеж, используемые с числами.
type numeralNumberCase int

// numeralNumberCase choises
const (
	// единственное число, именительный падеж (1, 21 час (но 11 часов))
	singularNominative numeralNumberCase = iota
	// единственное число, родительный падеж (2, 3, 4, 22 часа (но 12, 13, 14 часов))
	singularGenitive
	// множественное число, родительный падеж (0, 5 (и всё остальное) часов)
	pluralGenitive
)

func getNumeralNumberCasePrim(last2 int64) numeralNumberCase {
	// last2 от 0 до 99
	// в общем случае определяется двумя последними цифрами
	if last2 == 11 || last2 == 12 || last2 == 13 || last2 == 14 {
		return pluralGenitive
	}
	// теперь определяется одной последней цифрой
	last1 := last2 % 10
	if last1 == 1 {
		return singularNominative
	}
	if last1 == 2 || last1 == 3 || last1 == 4 {
		return singularGenitive
	}
	return pluralGenitive
}

func getNumeralNumberCase(i int64) numeralNumberCase {
	return getNumeralNumberCasePrim(mathhelper.AbsInt(i) % 100)
}

func numberAndRussianItems(items int64, showZero bool, russianItems string) string {
	if !showZero && items == 0 {
		return ""
	}
	return fmt.Sprintf("%d %s", items, russianItems)
}

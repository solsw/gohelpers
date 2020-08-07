package russian

import (
	"math"
	"strings"
)

// Binder of whole and fractional parts.
type (
	Binder int
)

// Supported binders.
const (
	None Binder = iota
	And
	Whole
)

var (
	wholeNumberCase = [3]string{"целая", "целых", "целых"}
)

// FloatInWords returns 'f' represented in russian words.
// 'withZeroInt' is used by integer part of 'f' (see help for IntInWords function).
func FloatInWords(f float64, fraction Fraction, binder Binder,
	showZeroInt, showZeroFrac, withZeroInt bool) string {
	absF := math.Abs(f)
	fint, frac := math.Modf(absF)
	ifint := int64(fint)
	sfrac := FracInWords(frac, fraction, showZeroFrac)
	var sb strings.Builder
	if ifint == 0 && !showZeroInt && len(sfrac) > 0 {
		sb.WriteString(sfrac)
	} else {
		gender := Masculine
		if binder == Whole {
			gender = Feminine
		}
		sb.WriteString(IntInWords(ifint, withZeroInt, gender))
		switch binder {
		case And:
			if len(sfrac) > 0 {
				sb.WriteString(" и")
			}
		case Whole:
			sb.WriteString(" ")
			sb.WriteString(wholeNumberCase[getNumeralNumberCase(ifint)])
		}
		if len(sfrac) > 0 {
			sb.WriteString(" ")
			sb.WriteString(sfrac)
		}
	}
	if f < 0 {
		return "минус " + sb.String()
	}
	return sb.String()
}

package russian

import (
	"errors"
	"math"
	"strings"

	"github.com/solsw/gohelpers/mathhelper"
)

const (
	zeroInWords string = "ноль"
)

func tripleInWords(triple int64, gender GrammaticalGender) (string, error) {
	if !(0 <= triple && triple <= 999) {
		return "", errors.New("triple out of range")
	}
	if triple == 0 {
		return zeroInWords, nil
	}
	ones := triple % 10           // number of ones
	tens := triple / 10 % 10      // number of tens
	hundreds := triple / 100 % 10 // number of hundreds
	var o string
	if ones > 0 {
		if tens == 1 {
			switch ones {
			case 1:
				o = "одиннадцать"
			case 2:
				o = "двенадцать"
			case 3:
				o = "тринадцать"
			case 4:
				o = "четырнадцать"
			case 5:
				o = "пятнадцать"
			case 6:
				o = "шестнадцать"
			case 7:
				o = "семнадцать"
			case 8:
				o = "восемнадцать"
			case 9:
				o = "девятнадцать"
			}
		} else { // tens != 1
			switch ones {
			case 1:
				switch gender {
				case Neuter:
					o = "одно"
				case Masculine:
					o = "один"
				case Feminine:
					o = "одна"
				}
			case 2:
				switch gender {
				case Feminine:
					o = "две"
				default:
					o = "два"
				}
			case 3:
				o = "три"
			case 4:
				o = "четыре"
			case 5:
				o = "пять"
			case 6:
				o = "шесть"
			case 7:
				o = "семь"
			case 8:
				o = "восемь"
			case 9:
				o = "девять"
			}
		}
	} // if ones > 0
	var t string
	if (tens > 1) || ((tens == 1) && (ones == 0)) {
		switch tens {
		case 1:
			t = "десять"
		case 2:
			t = "двадцать"
		case 3:
			t = "тридцать"
		case 4:
			t = "сорок"
		case 5:
			t = "пятьдесят"
		case 6:
			t = "шестьдесят"
		case 7:
			t = "семьдесят"
		case 8:
			t = "восемьдесят"
		case 9:
			t = "девянoсто"
		}
	} // if (tens > 1) || ((tens == 1) && (ones == 0))
	var h string
	if hundreds > 0 {
		switch hundreds {
		case 1:
			h = "сто"
		case 2:
			h = "двести"
		case 3:
			h = "триста"
		case 4:
			h = "четыреста"
		case 5:
			h = "пятьсот"
		case 6:
			h = "шестьсот"
		case 7:
			h = "семьсот"
		case 8:
			h = "восемьсот"
		case 9:
			h = "девятьсот"
		}
	} // if hundreds > 0
	var sb strings.Builder
	if len(h) > 0 {
		sb.WriteString(h)
	}
	if len(t) > 0 {
		if sb.Len() > 0 {
			sb.WriteRune(' ')
		}
		sb.WriteString(t)
	}
	if len(o) > 0 {
		if sb.Len() > 0 {
			sb.WriteRune(' ')
		}
		sb.WriteString(o)
	}
	return sb.String(), nil
}

const minInt64InWords string = // -9 223 372 036 854 775 808
"минус девять квинтиллионов двести двадцать три квадриллиона триста семьдесят два триллиона " +
	"тридцать шесть миллиардов восемьсот пятьдесят четыре миллиона семьсот семьдесят пять тысяч восемьсот восемь"

// IntInWords returns 'n' represented in russian words.
// If 'withZero' is false, zero triples are omitted.
// 'gender' determines russian grammatical gender for ones.
func IntInWords(n int64, withZero bool, gender GrammaticalGender) string {
	if n == 0 {
		return zeroInWords
	}
	// since -math.MinInt64 is out of int64 range
	if n == math.MinInt64 {
		return minInt64InWords
	}
	absN := mathhelper.AbsInt(n)
	var res string
	var tripleNumber int
	for absN > 0 {
		tripleNumber++
		triple := absN % 1000
		if triple > 0 || (triple == 0 && withZero) {
			if tripleNumber == 1 {
				res, _ = tripleInWords(triple, gender)
			} else {
				var newRes string
				if tripleNumber == 2 {
					newRes, _ = tripleInWords(triple, Feminine)
				} else {
					newRes, _ = tripleInWords(triple, Masculine)
				}
				if len(newRes) > 0 {
					var rank string
					switch tripleNumber {
					case 2:
						rank = Thousands(triple)
					case 3:
						rank = Millions(triple)
					case 4:
						rank = Milliards(triple)
					case 5:
						rank = Trillions(triple)
					case 6:
						rank = Quadrillions(triple)
					case 7:
						rank = Quintillions(triple)
					}
					newRes += " " + rank
					if len(res) > 0 {
						res = newRes + " " + res
					} else {
						res = newRes
					}
				}
			}
		}
		absN /= 1000
	}
	if n < 0 {
		return "минус " + res
	}
	return res
}

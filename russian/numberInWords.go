package russian

import (
	"errors"
	"math"
	"strings"
)

const (
	zeroInWords string = "ноль"
)

func digit3InWords(digit3 int64, gender GrammaticalGender) (string, error) {
	if !(0 <= digit3 && digit3 <= 999) {
		return "", errors.New("digit3 out of range")
	}
	if digit3 == 0 {
		return zeroInWords, nil
	}
	ones := digit3 % 10           // number of ones
	tens := digit3 / 10 % 10      // number of tens
	hundreds := digit3 / 100 % 10 // number of hundreds
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
				if gender == Masculine {
					o = "один"
				} else {
					if gender == Feminine {
						o = "одна"
					} else {
						o = "одно"
					}
				}
			case 2:
				if gender == Feminine {
					o = "две"
				} else {
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
	} // if ones > 0 {
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
	} //
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
	} // if hundreds > 0 {
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

// NumberInWords returns 'number' represented in russian words.
// If 'withZero' is false, zero 3 digit groupings will be omitted.
// 'gender' determines russian grammatical gender for ones.
func NumberInWords(number int64, withZero bool, gender GrammaticalGender) string {
	if number == 0 {
		return zeroInWords
	}
	// since -math.MinInt64 is out of int64 range
	if number == math.MinInt64 {
		return minInt64InWords
	}
	var absN int64
	if number < 0 {
		absN = -number
	} else {
		absN = number
	}
	var res string
	var numberOf3 int
	for absN > 0 {
		numberOf3++
		digit3 := absN % 1000
		if digit3 > 0 || (digit3 == 0 && withZero) {
			if numberOf3 == 1 {
				res, _ = digit3InWords(digit3, gender)
			} else {
				var newRes string
				if numberOf3 == 2 {
					newRes, _ = digit3InWords(digit3, Feminine)
				} else {
					newRes, _ = digit3InWords(digit3, Masculine)
				}
				if len(newRes) > 0 {
					var razryad string
					switch numberOf3 {
					case 2:
						razryad = Thousands(digit3)
					case 3:
						razryad = Millions(digit3)
					case 4:
						razryad = Milliards(digit3)
					case 5:
						razryad = Trillions(digit3)
					case 6:
						razryad = Quadrillions(digit3)
					case 7:
						razryad = Quintillions(digit3)
					}
					newRes += " " + razryad
				}
				if len(newRes) > 0 && len(res) > 0 {
					res = newRes + " " + res
				} else if len(newRes) > 0 {
					res = newRes
				}
			}
		}
		absN /= 1000
	}
	if number < 0 {
		return "минус " + res
	}
	return res
}

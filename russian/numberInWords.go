package russian

import (
	"errors"
	"math"
	"strings"
)

const nol string = "ноль"

func threeDigitNumberInWords(digits3 int64, gender GrammaticalGender) (string, error) {
	if digits3 < 0 || 999 < digits3 {
		return "", errors.New("digits3 out of range")
	}
	if digits3 == 0 {
		return nol, nil
	}
	// digits3 - от 1 до 999
	ones := digits3 % 10           // number of ones
	tens := digits3 / 10 % 10      // number of tens
	hundreds := digits3 / 100 % 10 // number of hundreds
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

// NumberInWords returns 'number' represented by russian words.
func NumberInWords(number int64, gender GrammaticalGender, withZeros bool) string {
	if number == 0 {
		return nol
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
	var nomerTroyki int
	for absN > 0 {
		nomerTroyki++
		digits3 := absN % 1000
		if digits3 > 0 || (withZeros && digits3 == 0) {
			if nomerTroyki == 1 {
				res, _ = threeDigitNumberInWords(digits3, gender)
			} else {
				var newRes string
				if nomerTroyki == 2 {
					newRes, _ = threeDigitNumberInWords(digits3, Feminine)
				} else {
					newRes, _ = threeDigitNumberInWords(digits3, Masculine)
				}
				if len(newRes) > 0 {
					var razryad string
					switch nomerTroyki {
					case 2:
						razryad = Thousands(digits3)
					case 3:
						razryad = Millions(digits3)
					case 4:
						razryad = Milliards(digits3)
					case 5:
						razryad = Trillions(digits3)
					case 6:
						razryad = Quadrillions(digits3)
					case 7:
						razryad = Quintillions(digits3)
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

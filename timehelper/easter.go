package timehelper

import (
	"time"
)

// CatholicEaster returns Gregorian date of catholic Easter for 'year'.
func CatholicEaster(year int) time.Time {
	// вычисляет month, day для года year григорианского календаря по книге Jean Meeus Astronomical Algorithms
	// см. http://ru.wikipedia.org/wiki/%D0%9F%D0%B0%D1%81%D1%85%D0%B0%D0%BB%D0%B8%D1%8F
	a := year % 19
	b := year / 100
	c := year % 100
	d := b / 4
	e := b % 4
	f := (b + 8) / 25
	g := (b - f + 1) / 3
	h := (19*a + b - d - g + 15) % 30
	i := c / 4
	k := c % 4
	l := (32 + 2*e + 2*i - h - k) % 7
	m := (a + 11*h + 22*l) / 451
	month := (h + l - 7*m + 114) / 31
	day := ((h + l - 7*m + 114) % 31) + 1

	return time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)
}

// OrthodoxEaster returns Gregorian date of orthodox Easter for 'year'.
func OrthodoxEaster(year int) time.Time {
	// вычисляет month, day для года year юлианского календаря по книге Jean Meeus Astronomical Algorithms
	// см. http://ru.wikipedia.org/wiki/%D0%9F%D0%B0%D1%81%D1%85%D0%B0%D0%BB%D0%B8%D1%8F
	a := year % 4
	b := year % 7
	c := year % 19
	d := (19*c + 15) % 30
	e := (2*a + 4*b - d + 34) % 7
	month := (d + e + 114) / 31
	day := ((d + e + 114) % 31) + 1

	res, _ := JulianToGregorian(time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC))
	return res
}

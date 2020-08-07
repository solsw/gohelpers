package russian

var (
	thousandNumberCase    = [3]string{"тысяча", "тысячи", "тысяч"}
	millionNumberCase     = [3]string{"миллион", "миллиона", "миллионов"}
	milliardNumberCase    = [3]string{"миллиард", "миллиарда", "миллиардов"}
	trillionNumberCase    = [3]string{"триллион", "триллиона", "триллионов"}
	quadrillionNumberCase = [3]string{"квадриллион", "квадриллиона", "квадриллионов"}
	quintillionNumberCase = [3]string{"квинтиллион", "квинтиллиона", "квинтиллионов"}
)

// Thousands returns russian for "thousand" corresponding to 'n'.
func Thousands(n int64) string {
	return thousandNumberCase[getNumeralNumberCase(n)]
}

// NThousands returns string containing 'n' and corresponding russian for "thousand".
// If 'n' is 0 and 'showZero' is false, empty string is returned.
func NThousands(n int64, showZero bool) string {
	return IntAndItems(n, showZero, Thousands(n))
}

// Millions returns russian for "million" corresponding to 'n'.
func Millions(n int64) string {
	return millionNumberCase[getNumeralNumberCase(n)]
}

// NMillions returns string containing 'n' and corresponding russian for "million".
// If 'n' is 0 and 'showZero' is false, empty string is returned.
func NMillions(n int64, showZero bool) string {
	return IntAndItems(n, showZero, Millions(n))
}

// Milliards returns russian for "milliard" corresponding to 'n'.
func Milliards(n int64) string {
	return milliardNumberCase[getNumeralNumberCase(n)]
}

// NMilliards returns string containing 'n' and corresponding russian for "milliard".
// If 'n' is 0 and 'showZero' is false, empty string is returned.
func NMilliards(n int64, showZero bool) string {
	return IntAndItems(n, showZero, Milliards(n))
}

// Billions returns russian for "milliard" corresponding to 'n'.
// (There is no "billion" in russian.)
func Billions(n int64) string {
	return Milliards(n)
}

// NBillions returns string containing 'n' and corresponding russian for "milliard".
// (There is no "billion" in russian.)
// If 'n' is 0 and 'showZero' is false, empty string is returned.
func NBillions(n int64, showZero bool) string {
	return IntAndItems(n, showZero, Billions(n))
}

// Trillions returns russian for "trillion" corresponding to 'n'.
func Trillions(n int64) string {
	return trillionNumberCase[getNumeralNumberCase(n)]
}

// NTrillions returns string containing 'n' and corresponding russian for "trillion".
// If 'n' is 0 and 'showZero' is false, empty string is returned.
func NTrillions(n int64, showZero bool) string {
	return IntAndItems(n, showZero, Trillions(n))
}

// Quadrillions returns russian for "quadrillion" corresponding to 'n'.
func Quadrillions(n int64) string {
	return quadrillionNumberCase[getNumeralNumberCase(n)]
}

// NQuadrillions returns string containing 'n' and corresponding russian for "quadrillion".
// If 'n' is 0 and 'showZero' is false, empty string is returned.
func NQuadrillions(n int64, showZero bool) string {
	return IntAndItems(n, showZero, Quadrillions(n))
}

// Quintillions returns russian for "quintillion" corresponding to 'n'.
func Quintillions(n int64) string {
	return quintillionNumberCase[getNumeralNumberCase(n)]
}

// NQuintillions returns string containing 'n' and corresponding russian for "quintillion".
// If 'n' is 0 and 'showZero' is false, empty string is returned.
func NQuintillions(n int64, showZero bool) string {
	return IntAndItems(n, showZero, Quintillions(n))
}

package russian

var (
	rublesByCase  = [3]string{"рубль", "рубля", "рублей"}
	kopecksByCase = [3]string{"копейка", "копейки", "копеек"}
)

// Rubles returns russian for "ruble" corresponding to 'n'.
func Rubles(n int64) string {
	return rublesByCase[getNumeralNumberCase(n)]
}

// NRubles returns string containing number 'n' and corresponding russian for "ruble".
// If 'n' is 0 and 'showZero' is false, empty string is returned.
func NRubles(n int64, showZero bool) string {
	return NumberAndItems(n, showZero, Rubles(n))
}

// NInWordsRubles returns string containing number 'n' in russian words and corresponding russian for "ruble".
// If 'n' is 0 and 'showZero' is false, empty string is returned.
// If 'withZero' is false, zero triples will be omitted.
func NInWordsRubles(n int64, showZero, withZero bool) string {
	return NumberInWordsAndItems(n, showZero, withZero, Masculine, Rubles(n))
}

// Kopecks returns russian for "kopeck" corresponding to 'n'.
func Kopecks(n int64) string {
	return kopecksByCase[getNumeralNumberCase(n)]
}

// NKopecks returns string containing number 'n' and corresponding russian for "kopeck".
// If 'n' is 0 and 'showZero' is false, empty string is returned.
func NKopecks(n int64, showZero bool) string {
	return NumberAndItems(n, showZero, Kopecks(n))
}

// NInWordsKopecks returns string containing number 'n' in russian words and corresponding russian for "kopeck".
// If 'n' is 0 and 'showZero' is false, empty string is returned.
// If 'withZero' is false, zero triples will be omitted.
func NInWordsKopecks(n int64, showZero, withZero bool) string {
	return NumberInWordsAndItems(n, showZero, withZero, Feminine, Kopecks(n))
}

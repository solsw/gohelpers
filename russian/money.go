package russian

var (
	rublesNumberCase  = [3]string{"рубль", "рубля", "рублей"}
	kopecksNumberCase = [3]string{"копейка", "копейки", "копеек"}
)

// Rubles returns russian for "ruble" corresponding to 'n'.
func Rubles(n int64) string {
	return rublesNumberCase[getNumeralNumberCase(n)]
}

// NRubles returns string containing 'n' and corresponding russian for "ruble".
// If 'n' is 0 and 'showZero' is false, empty string is returned.
func NRubles(n int64, showZero bool) string {
	return IntAndItems(n, showZero, Rubles(n))
}

// NInWordsRubles returns string containing 'n' in russian words and corresponding russian for "ruble".
// If 'n' is 0 and 'showZero' is false, empty string is returned.
// If 'withZero' is false, zero triples are omitted.
func NInWordsRubles(n int64, showZero, withZero bool) string {
	return IntInWordsAndItems(n, showZero, withZero, Masculine, Rubles(n))
}

// Kopecks returns russian for "kopeck" corresponding to 'n'.
func Kopecks(n int64) string {
	return kopecksNumberCase[getNumeralNumberCase(n)]
}

// NKopecks returns string containing 'n' and corresponding russian for "kopeck".
// If 'n' is 0 and 'showZero' is false, empty string is returned.
func NKopecks(n int64, showZero bool) string {
	return IntAndItems(n, showZero, Kopecks(n))
}

// NInWordsKopecks returns string containing 'n' in russian words and corresponding russian for "kopeck".
// If 'n' is 0 and 'showZero' is false, empty string is returned.
// If 'withZero' is false, zero triples are omitted.
func NInWordsKopecks(n int64, showZero, withZero bool) string {
	return IntInWordsAndItems(n, showZero, withZero, Feminine, Kopecks(n))
}

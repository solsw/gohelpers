package russian

var (
	centuryNumberCase     = [3]string{"век", "века", "веков"}
	yearNumberCase        = [3]string{"год", "года", "лет"}
	monthNumberCase       = [3]string{"месяц", "месяца", "месяцев"}
	weekNumberCase        = [3]string{"неделя", "недели", "недель"}
	dayNumberCase         = [3]string{"день", "дня", "дней"}
	hourNumberCase        = [3]string{"час", "часа", "часов"}
	minuteNumberCase      = [3]string{"минута", "минуты", "минут"}
	secondNumberCase      = [3]string{"секунда", "секунды", "секунд"}
	millisecondNumberCase = [3]string{"миллисекунда", "миллисекунды", "миллисекунд"}
	microsecondNumberCase = [3]string{"микросекунда", "микросекунды", "микросекунд"}
	nanosecondNumberCase  = [3]string{"наносекунда", "наносекунды", "наносекунд"}
)

// Centuries returns russian for "century" corresponding to 'n'.
func Centuries(n int64) string {
	return centuryNumberCase[getNumeralNumberCase(n)]
}

// NCenturies returns string containing 'n' and corresponding russian for "century".
// If 'n' is 0 and 'showZero' is false, empty string is returned.
func NCenturies(n int64, showZero bool) string {
	return IntAndItems(n, showZero, Centuries(n))
}

// NInWordsCenturies returns string containing 'n' in russian words and corresponding russian for "century".
// If 'n' is 0 and 'showZero' is false, empty string is returned.
// If 'withZero' is false, zero triples are omitted.
func NInWordsCenturies(n int64, showZero, withZero bool) string {
	return IntInWordsAndItems(n, showZero, withZero, Masculine, Centuries(n))
}

// Years returns russian for "year" corresponding to 'n'.
func Years(n int64) string {
	return yearNumberCase[getNumeralNumberCase(n)]
}

// NYears returns string containing 'n' and corresponding russian for "year".
// If 'n' is 0 and 'showZero' is false, empty string is returned.
func NYears(n int64, showZero bool) string {
	return IntAndItems(n, showZero, Years(n))
}

// NInWordsYears returns string containing 'n' in russian words and corresponding russian for "year".
// If 'n' is 0 and 'showZero' is false, empty string is returned.
// If 'withZero' is false, zero triples are omitted.
func NInWordsYears(n int64, showZero, withZero bool) string {
	return IntInWordsAndItems(n, showZero, withZero, Masculine, Years(n))
}

// Months returns russian for "month" corresponding to 'n'.
func Months(n int64) string {
	return monthNumberCase[getNumeralNumberCase(n)]
}

// NMonths returns string containing 'n' and corresponding russian for "month".
// If 'n' is 0 and 'showZero' is false, empty string is returned.
func NMonths(n int64, showZero bool) string {
	return IntAndItems(n, showZero, Months(n))
}

// NInWordsMonths returns string containing 'n' in russian words and corresponding russian for "month".
// If 'n' is 0 and 'showZero' is false, empty string is returned.
// If 'withZero' is false, zero triples are omitted.
func NInWordsMonths(n int64, showZero, withZero bool) string {
	return IntInWordsAndItems(n, showZero, withZero, Masculine, Months(n))
}

// Weeks returns russian for "week" corresponding to 'n'.
func Weeks(n int64) string {
	return weekNumberCase[getNumeralNumberCase(n)]
}

// NWeeks returns string containing 'n' and corresponding russian for "week".
// If 'n' is 0 and 'showZero' is false, empty string is returned.
func NWeeks(n int64, showZero bool) string {
	return IntAndItems(n, showZero, Weeks(n))
}

// NInWordsWeeks returns string containing 'n' in russian words and corresponding russian for "week".
// If 'n' is 0 and 'showZero' is false, empty string is returned.
// If 'withZero' is false, zero triples are omitted.
func NInWordsWeeks(n int64, showZero, withZero bool) string {
	return IntInWordsAndItems(n, showZero, withZero, Feminine, Weeks(n))
}

// Days returns russian for "day" corresponding to 'n'.
func Days(n int64) string {
	return dayNumberCase[getNumeralNumberCase(n)]
}

// NDays returns string containing 'n' and corresponding russian for "day".
// If 'n' is 0 and 'showZero' is false, empty string is returned.
func NDays(n int64, showZero bool) string {
	return IntAndItems(n, showZero, Days(n))
}

// NInWordsDays returns string containing 'n' in russian words and corresponding russian for "day".
// If 'n' is 0 and 'showZero' is false, empty string is returned.
// If 'withZero' is false, zero triples are omitted.
func NInWordsDays(n int64, showZero, withZero bool) string {
	return IntInWordsAndItems(n, showZero, withZero, Masculine, Days(n))
}

// Hours returns russian for "hour" corresponding to 'n'.
func Hours(n int64) string {
	return hourNumberCase[getNumeralNumberCase(n)]
}

// NHours returns string containing 'n' and corresponding russian for "hour".
// If 'n' is 0 and 'showZero' is false, empty string is returned.
func NHours(n int64, showZero bool) string {
	return IntAndItems(n, showZero, Hours(n))
}

// NInWordsHours returns string containing 'n' in russian words and corresponding russian for "hour".
// If 'n' is 0 and 'showZero' is false, empty string is returned.
// If 'withZero' is false, zero triples are omitted.
func NInWordsHours(n int64, showZero, withZero bool) string {
	return IntInWordsAndItems(n, showZero, withZero, Masculine, Hours(n))
}

// Minutes returns russian for "minute" corresponding to 'n'.
func Minutes(n int64) string {
	return minuteNumberCase[getNumeralNumberCase(n)]
}

// NMinutes returns string containing 'n' and corresponding russian for "minute".
// If 'n' is 0 and 'showZero' is false, empty string is returned.
func NMinutes(n int64, showZero bool) string {
	return IntAndItems(n, showZero, Minutes(n))
}

// NInWordsMinutes returns string containing 'n' in russian words and corresponding russian for "minute".
// If 'n' is 0 and 'showZero' is false, empty string is returned.
// If 'withZero' is false, zero triples are omitted.
func NInWordsMinutes(n int64, showZero, withZero bool) string {
	return IntInWordsAndItems(n, showZero, withZero, Feminine, Minutes(n))
}

// Seconds returns russian for "second" corresponding to 'n'.
func Seconds(n int64) string {
	return secondNumberCase[getNumeralNumberCase(n)]
}

// NSeconds returns string containing 'n' and corresponding russian for "second".
// If 'n' is 0 and 'showZero' is false, empty string is returned.
func NSeconds(n int64, showZero bool) string {
	return IntAndItems(n, showZero, Seconds(n))
}

// NInWordsSeconds returns string containing 'n' in russian words and corresponding russian for "second".
// If 'n' is 0 and 'showZero' is false, empty string is returned.
// If 'withZero' is false, zero triples are omitted.
func NInWordsSeconds(n int64, showZero, withZero bool) string {
	return IntInWordsAndItems(n, showZero, withZero, Feminine, Seconds(n))
}

// Milliseconds returns russian for "millisecond" corresponding to 'n'.
func Milliseconds(n int64) string {
	return millisecondNumberCase[getNumeralNumberCase(n)]
}

// NMilliseconds returns string containing 'n' and corresponding russian for "millisecond".
// If 'n' is 0 and 'showZero' is false, empty string is returned.
func NMilliseconds(n int64, showZero bool) string {
	return IntAndItems(n, showZero, Milliseconds(n))
}

// NInWordsMilliseconds returns string containing 'n' in russian words and corresponding russian for "millisecond".
// If 'n' is 0 and 'showZero' is false, empty string is returned.
// If 'withZero' is false, zero triples are omitted.
func NInWordsMilliseconds(n int64, showZero, withZero bool) string {
	return IntInWordsAndItems(n, showZero, withZero, Feminine, Milliseconds(n))
}

// Microseconds returns russian for "microsecond" corresponding to 'n'.
func Microseconds(n int64) string {
	return microsecondNumberCase[getNumeralNumberCase(n)]
}

// NMicroseconds returns string containing 'n' and corresponding russian for "microsecond".
// If 'n' is 0 and 'showZero' is false, empty string is returned.
func NMicroseconds(n int64, showZero bool) string {
	return IntAndItems(n, showZero, Microseconds(n))
}

// NInWordsMicroseconds returns string containing 'n' in russian words and corresponding russian for "microsecond".
// If 'n' is 0 and 'showZero' is false, empty string is returned.
// If 'withZero' is false, zero triples are omitted.
func NInWordsMicroseconds(n int64, showZero, withZero bool) string {
	return IntInWordsAndItems(n, showZero, withZero, Feminine, Microseconds(n))
}

// Nanoseconds returns russian for "nanosecond" corresponding to 'n'.
func Nanoseconds(n int64) string {
	return nanosecondNumberCase[getNumeralNumberCase(n)]
}

// NNanoseconds returns string containing 'n' and corresponding russian for "nanosecond".
// If 'n' is 0 and 'showZero' is false, empty string is returned.
func NNanoseconds(n int64, showZero bool) string {
	return IntAndItems(n, showZero, Nanoseconds(n))
}

// NInWordsNanoseconds returns string containing 'n' in russian words and corresponding russian for "nanosecond".
// If 'n' is 0 and 'showZero' is false, empty string is returned.
// If 'withZero' is false, zero triples are omitted.
func NInWordsNanoseconds(n int64, showZero, withZero bool) string {
	return IntInWordsAndItems(n, showZero, withZero, Feminine, Nanoseconds(n))
}

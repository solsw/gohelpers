package stringhelper

import (
	"sort"
)

// Unique returns unique strings from 'ss', preserving order of strings in 'ss'.
func Unique(ss []string) []string {
	if len(ss) < 2 {
		return ss
	}
	var res []string
	var m = make(map[string]interface{})
	for i := range ss {
		_, ok := m[ss[i]]
		if !ok {
			res = append(res, ss[i])
			m[ss[i]] = nil
		}
	}
	return res
}

// UniqueSorted returns sorted unique strings from 'ss'.
// (May be up to two times faster than Unique. Subject for benchmarking.)
func UniqueSorted(ss []string) []string {
	if len(ss) < 2 {
		return ss
	}
	sort.Strings(ss)
	res := []string{ss[0]}
	for i := 1; i < len(ss); i++ {
		if ss[i] != ss[i-1] {
			res = append(res, ss[i])
		}
	}
	return res
}

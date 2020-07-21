package combinatorics

import (
	"errors"

	"github.com/solsw/gohelpers/slicehelper"
)

// Permutations returns all permutations (http://en.wikipedia.org/wiki/Permutations) of the slice's elements.
func Permutations(sl []interface{}) ([][]interface{}, error) {
	if len(sl) == 0 {
		return nil, errors.New("empty slice")
	}
	if len(sl) == 1 {
		return [][]interface{}{sl}, nil
	}
	var r [][]interface{}
	for i := 0; i < len(sl); i++ {
		// element at i position
		eI := sl[i]
		// slice without element at i position
		withoutI, _ := slicehelper.RemoveAt(sl, i)
		perms, _ := Permutations(withoutI)
		for j := 0; j < len(perms); j++ {
			es := slicehelper.ConcatElSl(eI, perms[j])
			r = append(r, es)
		}
	}
	return r, nil
}

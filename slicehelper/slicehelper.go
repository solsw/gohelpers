// Package slicehelper contains various slice-related helpers.
package slicehelper

import (
	cryptoRand "crypto/rand"
	"errors"
	"math/big"
	"reflect"
	"sort"
	"unsafe"
)

// ConcatElSl returns new slice concatenating the element and the slice.
func ConcatElSl(el interface{}, sl []interface{}) []interface{} {
	return append(append(make([]interface{}, 0, 1+len(sl)), el), sl...)
}

// ContainsEq reports whether 'sl' contains 'el' in sense of 'eq'.
func ContainsEq(sl []interface{}, el interface{}, eq func(interface{}, interface{}) bool) (bool, error) {
	if eq == nil {
		return false, errors.New("'eq' is nil")
	}
	for _, e := range sl {
		if eq(e, el) {
			return true, nil
		}
	}
	return false, nil
}

// ContainsEqMust calls ContainsEq, but panics in case of error.
func ContainsEqMust(sl []interface{}, el interface{}, eq func(interface{}, interface{}) bool) bool {
	r, err := ContainsEq(sl, el, eq)
	if err != nil {
		panic(err)
	}
	return r
}

// Contains reports whether 'sl' contains 'el' using reflect.DeepEqual as equality comparer.
func Contains(sl []interface{}, el interface{}) bool {
	r, _ := ContainsEq(sl, el, reflect.DeepEqual)
	return r
}

// ContainsCmp reports whether *sorted* 'sl' contains 'el' using 'cmp'.
//
// 'cmp' compares two elements and returns negative if the first one is less than the second,
// zero if the first one is equal to the second and positive if the first one is greater than the second.
func ContainsCmp(sl []interface{}, el interface{}, cmp func(interface{}, interface{}) int) (bool, error) {
	if cmp == nil {
		return false, errors.New("'cmp' is nil")
	}
	idx := sort.Search(len(sl), func(i int) bool {
		return cmp(el, sl[i]) <= 0
	})
	if idx == len(sl) {
		return false, nil
	}
	return cmp(el, sl[idx]) == 0, nil
}

// ContainsCmpMust calls ContainsCmp, but panics in case of error.
func ContainsCmpMust(sl []interface{}, el interface{}, cmp func(interface{}, interface{}) int) bool {
	r, err := ContainsCmp(sl, el, cmp)
	if err != nil {
		panic(err)
	}
	return r
}

// RemoveAt removes element at 'idx' position from the slice returning new slice.
func RemoveAt(sl []interface{}, idx int) ([]interface{}, error) {
	if len(sl) == 0 {
		return nil, errors.New("empty slice")
	}
	if idx < 0 || idx >= len(sl) {
		return nil, errors.New("wrong index")
	}
	var r []interface{}
	r = make([]interface{}, 0, len(sl)-1)
	if idx > 0 {
		// removing not first element
		r = append(r, sl[:idx]...)
	}
	if idx < len(sl)-1 {
		// removing not last element
		r = append(r, sl[idx+1:]...)
	}
	return r, nil
}

// RemoveAtInPlace removes element at 'idx' position from the slice in place.
func RemoveAtInPlace(sl *[]interface{}, idx int) (*[]interface{}, error) {
	if *sl == nil || len(*sl) == 0 {
		return nil, errors.New("empty slice")
	}
	if idx < 0 || idx >= len(*sl) {
		return nil, errors.New("wrong index")
	}
	if len(*sl) == 1 {
		return &[]interface{}{}, nil
	}
	copy((*sl)[idx:], (*sl)[idx+1:])
	hdr := (*reflect.SliceHeader)(unsafe.Pointer(sl))
	hdr.Len = len(*sl) - 1
	return sl, nil
}

// ShuffleCr returns input slice shuffled the in place using crypto/rand package.
// ShuffleCr panics (by means of reflect.Swapper), if the provided interface is not a slice.
func ShuffleCr(sl interface{}) interface{} {
	sw := reflect.Swapper(sl)
	len := reflect.ValueOf(sl).Len()
	if len == 0 {
		return sl
	}
	bigLen := big.NewInt(int64(len))
	for i := 0; i < len; i++ {
		j, _ := cryptoRand.Int(cryptoRand.Reader, bigLen)
		sw(i, int(j.Int64()))
	}
	return sl
}

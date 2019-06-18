package slicehelper

import (
	cryptoRand "crypto/rand"
	"errors"
	"math/big"
	"reflect"
	"unsafe"
)

// ConcatEls concatenates elements into the slice.
func ConcatEls(els ...interface{}) (*[]interface{}, error) {
	if els == nil {
		return nil, errors.New("'els' is nil")
	}
	r := append(make([]interface{}, 0, len(els)), els...)
	return &r, nil
}

// ConcatElSl concatenates the element with the slice.
func ConcatElSl(el interface{}, sl *[]interface{}) (*[]interface{}, error) {
	if el == nil && *sl == nil {
		return nil, errors.New("'el' and 'sl' are nils")
	}
	if el == nil {
		return sl, nil
	}
	var r []interface{}
	if *sl == nil || len(*sl) == 0 {
		r = []interface{}{el}
	} else {
		r = append(append(make([]interface{}, 0, 1+len(*sl)), el), *sl...)
	}
	return &r, nil
}

// ConcatSlEl concatenates the slice with the element.
func ConcatSlEl(sl *[]interface{}, el interface{}) (*[]interface{}, error) {
	if *sl == nil && el == nil {
		return nil, errors.New("'sl' and 'el' are nils")
	}
	if el == nil {
		return sl, nil
	}
	var r []interface{}
	if *sl == nil || len(*sl) == 0 {
		r = []interface{}{el}
	} else {
		r = append(*sl, el)
	}
	return &r, nil
}

// ConcatSlSl concatenates two slices.
func ConcatSlSl(sl1 *[]interface{}, sl2 *[]interface{}) (*[]interface{}, error) {
	if *sl1 == nil && *sl2 == nil {
		return nil, errors.New("'sl1' and 'sl2' are nils")
	}
	if *sl1 == nil {
		return sl2, nil
	}
	if *sl2 == nil {
		return sl1, nil
	}
	if len(*sl1) == 0 {
		return sl2, nil
	}
	if len(*sl2) == 0 {
		return sl1, nil
	}
	r := append(*sl1, *sl2)
	return &r, nil
}

// RemoveAt removes element at 'idx' position from the slice returning new slice.
func RemoveAt(sl *[]interface{}, idx int) (*[]interface{}, error) {
	if *sl == nil || len(*sl) == 0 {
		return nil, errors.New("'sl' is empty")
	}
	if idx < 0 || idx >= len(*sl) {
		return nil, errors.New("wrong index")
	}
	var r []interface{}
	r = make([]interface{}, 0, len(*sl)-1)
	if idx > 0 { // removing not first element
		r = append(r, (*sl)[:idx]...)
	}
	if idx < len(*sl)-1 { // removing not last element
		r = append(r, (*sl)[idx+1:]...)
	}
	return &r, nil
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

// ShuffleCr shuffles the slice in place using crypto/rand package.
// Shuffled slice is returned by the function.
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

package jsonhelper

import (
	"encoding/json"
)

// MarshalMust calls json.Marshal, but panics in case of error.
func MarshalMust(v interface{}) []byte {
	bb, err := json.Marshal(v)
	if err != nil {
		panic(err)
	}
	return bb
}

// MarshalIndentMust calls json.MarshalIndent, but panics in case of error.
func MarshalIndentMust(v interface{}, prefix, indent string) []byte {
	bb, err := json.MarshalIndent(v, prefix, indent)
	if err != nil {
		panic(err)
	}
	return bb
}

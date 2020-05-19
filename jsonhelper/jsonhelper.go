// Package jsonhelper contains various encoding/json helpers.
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

// MarshalIndentDef calls json.MarshalIndent with prefix="" and indent="  ".
func MarshalIndentDef(v interface{}) ([]byte, error) {
	return json.MarshalIndent(v, "", "  ")
}

// MarshalIndentMust calls json.MarshalIndent, but panics in case of error.
func MarshalIndentMust(v interface{}, prefix, indent string) []byte {
	bb, err := json.MarshalIndent(v, prefix, indent)
	if err != nil {
		panic(err)
	}
	return bb
}

// MarshalIndentMustDef calls json.MarshalIndent with prefix="" and indent="  ", but panics in case of error.
func MarshalIndentMustDef(v interface{}) []byte {
	return MarshalIndentMust(v, "", "  ")
}

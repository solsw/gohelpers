// Package jsonhelper contains various encoding/json helpers.
package jsonhelper

import (
	"encoding/json"
	"io"
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
// (See json.MarshalIndent for 'prefix' and 'indent' usage.)
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

// Format formats JSON-encoded data from 'r' to 'w'.
// (See json.MarshalIndent for 'prefix' and 'indent' usage.)
func Format(r io.Reader, w io.Writer, prefix, indent string) error {
	d := json.NewDecoder(r)
	v := make(map[string]interface{})
	if err := d.Decode(&v); err != nil {
		return err
	}
	e := json.NewEncoder(w)
	e.SetIndent(prefix, indent)
	if err := e.Encode(v); err != nil {
		return err
	}
	return nil
}

// FormatDef calls Format with prefix="" and indent="  ".
func FormatDef(r io.Reader, w io.Writer) error {
	return Format(r, w, "", "  ")
}

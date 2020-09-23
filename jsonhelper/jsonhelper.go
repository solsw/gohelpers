// Package jsonhelper contains encoding/json helpers.
package jsonhelper

import (
	"encoding/json"
	"io"
	"strings"
)

// MustMarshal is like json.Marshal but panics in case of error.
func MustMarshal(v interface{}) []byte {
	bb, err := json.Marshal(v)
	if err != nil {
		panic(err)
	}
	return bb
}

// MustMarshalIndent is like json.MarshalIndent but panics in case of error.
// (See json.MarshalIndent for 'prefix' and 'indent' usage.)
func MustMarshalIndent(v interface{}, prefix, indent string) []byte {
	bb, err := json.MarshalIndent(v, prefix, indent)
	if err != nil {
		panic(err)
	}
	return bb
}

// MarshalIndentDef calls json.MarshalIndent with prefix="" and indent="  ".
func MarshalIndentDef(v interface{}) ([]byte, error) {
	return json.MarshalIndent(v, "", "  ")
}

// MustMarshalIndentDef is like MarshalIndentDef but panics in case of error.
func MustMarshalIndentDef(v interface{}) []byte {
	return MustMarshalIndent(v, "", "  ")
}

// Format reads JSON-encoded data from 'r', then writes formatted data to 'w'.
// (See json.MarshalIndent for 'prefix' and 'indent' usage.)
func Format(r io.Reader, w io.Writer, prefix, indent string) error {
	d := json.NewDecoder(r)
	v := make(map[string]interface{})
	if err := d.Decode(&v); err != nil {
		return err
	}
	e := json.NewEncoder(w)
	e.SetIndent(prefix, indent)
	return e.Encode(v)
}

// FormatDef calls Format with prefix="" and indent="  ".
func FormatDef(r io.Reader, w io.Writer) error {
	return Format(r, w, "", "  ")
}

// FormatStrToStr formats JSON-encoded 'json' to string.
// (See json.MarshalIndent for 'prefix' and 'indent' usage.)
func FormatStrToStr(json, prefix, indent string) (string, error) {
	var b strings.Builder
	if err := Format(strings.NewReader(json), &b, prefix, indent); err != nil {
		return "", err
	}
	return b.String(), nil
}

// MustFormatStrToStr is like FormatStrToStr but panics in case of error.
func MustFormatStrToStr(json, prefix, indent string) string {
	s, err := FormatStrToStr(json, prefix, indent)
	if err != nil {
		panic(err)
	}
	return s
}

// FormatStrToStrDef calls FormatStrToStr with prefix="" and indent="  ".
func FormatStrToStrDef(json string) (string, error) {
	return FormatStrToStr(json, "", "  ")
}

// MustFormatStrToStrDef is like FormatStrToStrDef but panics in case of error.
func MustFormatStrToStrDef(json string) string {
	return MustFormatStrToStr(json, "", "  ")
}

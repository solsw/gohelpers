// Package sqlhelper contains database/sql helpers.
package sqlhelper

import (
	"database/sql"
)

// StrToNullStr converts string to sql.NullString.
// If 's' is empty and 'emptyIsNULL' is true, 'ns.Valid' is false.
func StrToNullStr(s string, emptyIsNULL bool) (ns sql.NullString) {
	ns = sql.NullString{String: s}
	if len(s) > 0 || !emptyIsNULL {
		ns.Valid = true
	}
	return
}

// NullStrToStr converts sql.NullString to string.
// If 'ns.Valid' is false, empty string is returned.
func NullStrToStr(ns sql.NullString) string {
	if !ns.Valid {
		return ""
	}
	return ns.String
}

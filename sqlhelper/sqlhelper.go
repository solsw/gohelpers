// Package sqlhelper contains various database/sql helpers.
package sqlhelper

import (
	"database/sql"
)

// StrToNullStr converts string to sql.NullString.
func StrToNullStr(s string, emptyIsNULL bool) *sql.NullString {
	ns := sql.NullString{String: s}
	if len(s) > 0 || !emptyIsNULL {
		ns.Valid = true
	}
	return &ns
}

// NullStrToStr converts sql.NullString to string.
func NullStrToStr(ns *sql.NullString) string {
	if !ns.Valid {
		return ""
	}
	return ns.String
}

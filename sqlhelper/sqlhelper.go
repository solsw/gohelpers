// Package sqlhelper contains database/sql helpers.
package sqlhelper

import (
	"database/sql"
)

// StrToNullStr converts string to sql.NullString.
// If 's' is empty and 'emptyIsNULL' is true, 'ns.Valid' is false.
func StrToNullStr(s string, emptyIsNULL bool) sql.NullString {
	return sql.NullString{
		String: s,
		Valid:  !(s == "" && emptyIsNULL),
	}
}

// NullStrToStr converts sql.NullString to string.
// If 'ns.Valid' is false, empty string is returned.
func NullStrToStr(ns sql.NullString) string {
	if !ns.Valid {
		return ""
	}
	return ns.String
}

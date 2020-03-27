package sqlhelper

import (
	"database/sql"
)

// StrToNullStr converts string to sql.NullString.
func StrToNullStr(s string) *sql.NullString {
	return &sql.NullString{String: s, Valid: true}
}

// NullStrToStr converts sql.NullString to string.
func NullStrToStr(ns *sql.NullString) string {
	if !ns.Valid {
		return ""
	}
	return ns.String
}

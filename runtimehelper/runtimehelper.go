// Package runtimehelper contains runtime helpers.
package runtimehelper

import (
	"runtime"
)

// nthCallerName returns name of n-th caller function of nthCallerName.
func nthCallerName(n int) (string, bool) {
	pc, _, _, ok := runtime.Caller(n)
	if !ok {
		return "", false
	}
	return runtime.FuncForPC(pc).Name(), true
}

// CallerName returns name of the function that called CallerName.
func CallerName() (string, bool) {
	return nthCallerName(2)
}

// MustCallerName returns name of the function that called MustCallerName.
// In case of a failure empty string is returned.
func MustCallerName() string {
	// do not call CallerName() because it adds another stack frame to the calling stack
	fn, ok := nthCallerName(2)
	if !ok {
		return ""
	}
	return fn
}

// CallerCallerName returns name of the function that called function that called CallerCallerName.
func CallerCallerName() (string, bool) {
	return nthCallerName(3)
}

// MustCallerCallerName returns name of the function that called MustCallerCallerName.
// In case of a failure empty string is returned.
func MustCallerCallerName() string {
	fn, ok := nthCallerName(3)
	if !ok {
		return ""
	}
	return fn
}

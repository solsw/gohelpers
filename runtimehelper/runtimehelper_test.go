// Package oshelper contains runtime helper functions.

package runtimehelper

import (
	"runtime"
	"testing"
)

func TestMustCallerName(t *testing.T) {
	t.Run("1", func(t *testing.T) {
		pc, _, _, _ := runtime.Caller(0)
		want := runtime.FuncForPC(pc).Name()
		got := MustCallerName()
		if got != want {
			t.Errorf("MustCallerName() got = %v, want %v", got, want)
		}
	})
}

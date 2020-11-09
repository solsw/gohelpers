// Package oshelper contains runtime helper functions.

package runtimehelper

import (
	"runtime"
	"testing"
)

func TestCallerNameMust(t *testing.T) {
	t.Run("1", func(t *testing.T) {
		pc, _, _, _ := runtime.Caller(0)
		want := runtime.FuncForPC(pc).Name()
		got := CallerNameMust()
		if got != want {
			t.Errorf("CallerNameMust() got = %v, want %v", got, want)
		}
	})
}

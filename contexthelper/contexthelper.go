// Package contexthelper contains context helpers.
package contexthelper

import (
	"context"
	"os"
	"os/signal"
)

// WithTerminate returns context.Context
// that is canceled when os.Interrupt or os.Kill signal is received.
func WithTerminate(ctx context.Context) context.Context {
	// inspired by https://gist.github.com/SpeedyCoder/59911301eea4d91e42feafccaa9bcaf7
	ctx, cancel := context.WithCancel(ctx)
	go func() {
		sigch := make(chan os.Signal, 1)
		signal.Notify(sigch, os.Interrupt, os.Kill)
		<-sigch
		cancel()
	}()
	return ctx
}

// BoolValue returns the bool value associated with this context for 'key'
// and a bool indicating whether the value exists and is of bool type.
func BoolValue(ctx context.Context, key interface{}) (bool, bool) {
	v, ok := ctx.Value(key).(bool)
	return v, ok
}

// Int64Value returns the int64 value associated with this context for 'key'
// and a bool indicating whether the value exists and is of int64 type.
func Int64Value(ctx context.Context, key interface{}) (int64, bool) {
	v, ok := ctx.Value(key).(int64)
	return v, ok
}

// Float64Value returns the float64 value associated with this context for 'key'
// and a bool indicating whether the value exists and is of float64 type.
func Float64Value(ctx context.Context, key interface{}) (float64, bool) {
	v, ok := ctx.Value(key).(float64)
	return v, ok
}

// StringValue returns the string value associated with this context for 'key'
// and a bool indicating whether the value exists and is of string type.
func StringValue(ctx context.Context, key interface{}) (string, bool) {
	v, ok := ctx.Value(key).(string)
	return v, ok
}

package contexthelper

import (
	"context"
	"os"
	"os/signal"
)

// WithTerminate returns context.Context
// that is canceled when os.Interrupt or os.Kill is received.
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

package contexthelper

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// NewCombinedContext combines two context.Contexts into one.
func NewCombinedContext(ctx1, ctx2 context.Context) context.Context {
	return &combinedContext{ctx1, ctx2, make(chan struct{})}
}

// combinedContext implements the context.Context interface.
type combinedContext struct {
	ctx1, ctx2 context.Context
	done       chan struct{}
}

// Deadline implements the context.Context.Deadline method.
//
// If both deadlines are set, the latest is returned.
func (cc *combinedContext) Deadline() (time.Time, bool) {
	d1, ok1 := cc.ctx1.Deadline()
	d2, ok2 := cc.ctx2.Deadline()
	if !ok1 {
		return d2, ok2
	}
	if !ok2 {
		return d1, ok1
	}
	if d1.After(d2) {
		return d1, true
	}
	return d2, true
}

func merge2Dones(done1, done2 <-chan struct{}, done chan<- struct{}) {
	defer close(done)
	for done1 != nil || done2 != nil {
		select {
		case v1, ok1 := <-done1:
			if ok1 {
				done <- v1
			} else {
				done1 = nil
			}
		case v2, ok2 := <-done2:
			if ok2 {
				done <- v2
			} else {
				done2 = nil
			}
		}
	}
}

var once sync.Once

// Done implements the context.Context.Done method.
//
// Done is closed when both Contexts' Dones are closed.
func (cc *combinedContext) Done() <-chan struct{} {
	once.Do(func() {
		go merge2Dones(cc.ctx1.Done(), cc.ctx2.Done(), cc.done)
	})
	return cc.done
}

// Err implements the context.Context.Err method.
func (cc *combinedContext) Err() error {
	select {
	case <-cc.Done():
		return fmt.Errorf("%v : %v", cc.ctx1.Err(), cc.ctx2.Err())
	default:
		return nil
	}
}

// Value implements the context.Context.Value method.
//
// If both Contexts contain the value, the one from the first Context is returned.
func (cc *combinedContext) Value(key interface{}) interface{} {
	if v1 := cc.ctx1.Value(key); v1 != nil {
		return v1
	}
	if v2 := cc.ctx2.Value(key); v2 != nil {
		return v2
	}
	return nil
}

package contexthelper

import (
	"context"
	"errors"
	"sync"
	"time"
)

// NewAndContext combines two context.Contexts into one.
func NewAndContext(ctx1, ctx2 context.Context) context.Context {
	return &andContext{ctx1: ctx1, ctx2: ctx2}
}

// andContext implements the context.Context interface.
type andContext struct {
	ctx1, ctx2        context.Context
	done              chan struct{}
	onceDone, onceErr sync.Once
}

// Deadline implements the context.Context.Deadline method.
//
// If both deadlines are set, the latest one is returned.
func (cc *andContext) Deadline() (time.Time, bool) {
	dl1, ok1 := cc.ctx1.Deadline()
	dl2, ok2 := cc.ctx2.Deadline()
	if !ok1 {
		return dl2, ok2
	}
	if !ok2 {
		return dl1, ok1
	}
	if dl1.After(dl2) {
		return dl1, true
	}
	return dl2, true
}

// Done implements the context.Context.Done method.
//
// The returned channel is closed when both Contexts' Done channels are closed.
func (cc *andContext) Done() <-chan struct{} {
	cc.onceDone.Do(func() {
		if cc.ctx1.Done() != nil || cc.ctx2.Done() != nil {
			cc.done = make(chan struct{})
			go andDone(cc.ctx1.Done(), cc.ctx2.Done(), cc.done)
		}
	})
	return cc.done
}

// Err implements the context.Context.Err method.
//
// The returned error contains errors from first and second Contexts (if any) separated by colon.
func (cc *andContext) Err() error {
	select {
	case <-cc.Done():
		var er string
		cc.onceErr.Do(func() {
			if cc.ctx1.Err() != nil {
				er += cc.ctx1.Err().Error()
			}
			er += ":"
			if cc.ctx2.Err() != nil {
				er += cc.ctx2.Err().Error()
			}
		})
		return errors.New(er)
	default:
		return nil
	}
}

// Value implements the context.Context.Value method.
//
// If both Contexts contain the value, the one from the first Context is returned.
func (cc *andContext) Value(key interface{}) interface{} {
	if v1 := cc.ctx1.Value(key); v1 != nil {
		return v1
	}
	if v2 := cc.ctx2.Value(key); v2 != nil {
		return v2
	}
	return nil
}

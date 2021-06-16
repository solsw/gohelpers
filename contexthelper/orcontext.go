package contexthelper

import (
	"context"
	"errors"
	"sync"
	"time"
)

// OrContext combines two context.Contexts into one and implements the context.Context interface.
type OrContext struct {
	Ctx1, Ctx2   context.Context
	onceDeadline sync.Once
	deadline     time.Time
	okDeadline   bool
	onceDone     sync.Once
	done         chan struct{}
	onceErr      sync.Once
	err          error
}

// NewOrContext combines two context.Contexts into one.
func NewOrContext(ctx1, ctx2 context.Context) *OrContext {
	return &OrContext{Ctx1: ctx1, Ctx2: ctx2}
}

// Deadline implements the context.Context.Deadline method.
//
// If both deadlines are set, the earliest one is returned.
func (cc *OrContext) Deadline() (time.Time, bool) {
	cc.onceDeadline.Do(func() {
		dl1, ok1 := cc.Ctx1.Deadline()
		dl2, ok2 := cc.Ctx2.Deadline()
		if !ok1 {
			cc.deadline, cc.okDeadline = dl2, ok2
			return
		}
		if !ok2 {
			cc.deadline, cc.okDeadline = dl1, ok1
			return
		}
		if dl1.Before(dl2) {
			cc.deadline, cc.okDeadline = dl1, true
			return
		}
		cc.deadline, cc.okDeadline = dl2, true
	})
	return cc.deadline, cc.okDeadline
}

func orDone(done1, done2 <-chan struct{}, done chan<- struct{}) {
	// done1 and done2 are not both nil here
	select {
	case <-done1:
	case <-done2:
	}
	close(done)
}

// Done implements the context.Context.Done method.
//
// The returned channel is closed when either one of Contexts' Done channels is closed.
func (cc *OrContext) Done() <-chan struct{} {
	cc.onceDone.Do(func() {
		if cc.Ctx1.Done() == nil && cc.Ctx2.Done() == nil {
			return
		}
		cc.done = make(chan struct{})
		go orDone(cc.Ctx1.Done(), cc.Ctx2.Done(), cc.done)
	})
	return cc.done
}

// Err implements the context.Context.Err method.
//
// If both Contexts' Errs are nil, nil is returned.
// Otherwise, the returned error contains errors (if any) from Contexts separated by colon.
func (cc *OrContext) Err() error {
	select {
	case <-cc.Done():
		cc.onceErr.Do(func() {
			var er string
			if cc.Ctx1.Err() != nil {
				er += cc.Ctx1.Err().Error()
			}
			er += ":"
			if cc.Ctx2.Err() != nil {
				er += cc.Ctx2.Err().Error()
			}
			cc.err = errors.New(er)
		})
		return cc.err
	default:
		return nil
	}
}

// Value implements the context.Context.Value method.
//
// The method returns nil. To get values from Contexts call corresponding Value methods directly.
func (*OrContext) Value(key interface{}) interface{} {
	return nil
}

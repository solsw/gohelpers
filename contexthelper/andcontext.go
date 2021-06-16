package contexthelper

import (
	"context"
	"errors"
	"sync"
	"time"
)

// AndContext implements the context.Context interface.
type AndContext struct {
	Ctx1, Ctx2   context.Context
	onceDeadline sync.Once
	deadline     time.Time
	okDeadline   bool
	onceDone     sync.Once
	done         chan struct{}
	onceErr      sync.Once
	err          error
}

// NewAndContext combines two context.Contexts into one.
func NewAndContext(ctx1, ctx2 context.Context) *AndContext {
	return &AndContext{Ctx1: ctx1, Ctx2: ctx2}
}

// Deadline implements the context.Context.Deadline method.
//
// If both deadlines are set, the latest one is returned.
func (cc *AndContext) Deadline() (time.Time, bool) {
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
		if dl1.After(dl2) {
			cc.deadline, cc.okDeadline = dl1, true
			return
		}
		cc.deadline, cc.okDeadline = dl2, true
	})
	return cc.deadline, cc.okDeadline
}

func andDone(done1, done2 <-chan struct{}, done chan<- struct{}) {
	// done1 and done2 are not both nil here
	for done1 != nil || done2 != nil {
		select {
		case _, ok1 := <-done1:
			if !ok1 {
				done1 = nil
			}
		case _, ok2 := <-done2:
			if !ok2 {
				done2 = nil
			}
		}
	}
	close(done)
}

// Done implements the context.Context.Done method.
//
// The returned channel is closed when both Contexts' Done channels are closed.
func (cc *AndContext) Done() <-chan struct{} {
	cc.onceDone.Do(func() {
		if cc.Ctx1.Done() == nil && cc.Ctx2.Done() == nil {
			return
		}
		cc.done = make(chan struct{})
		go andDone(cc.Ctx1.Done(), cc.Ctx2.Done(), cc.done)
	})
	return cc.done
}

// Err implements the context.Context.Err method.
//
// If both Contexts' Errs are nil, nil is returned.
// Otherwise, the returned error contains errors (if any) from Contexts separated by colon.
func (cc *AndContext) Err() error {
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
func (*AndContext) Value(key interface{}) interface{} {
	return nil
}

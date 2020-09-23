package nullable

// Bool represents a bool type that may have no value.
type Bool struct {
	val *bool
}

// NullBool creates a new Bool without value.
func NullBool() Bool {
	return Bool{nil}
}

// NewBool creates a new Bool with value 'v'.
func NewBool(v bool) Bool {
	return Bool{&v}
}

// Has reports whether n has a value.
func (n *Bool) Has() bool {
	return n.val != nil
}

// Get returns (n's value, true) if n has value.
// Returns (false, false) if n has no value.
func (n *Bool) Get() (bool, bool) {
	if n.val == nil {
		return false, false
	}
	return *n.val, true
}

// MustGet is like Get but returns 'false' if n has no value.
func (n *Bool) MustGet() bool {
	r, ok := n.Get()
	if !ok {
		return false
	}
	return r
}

// Set sets the n's value to 'v'.
func (n *Bool) Set(v bool) {
	n.val = &v
}

// Null sets n as having no value.
func (n *Bool) Null() {
	n.val = nil
}

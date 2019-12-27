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

// Has reports whether 'n' has a value.
func (n *Bool) Has() bool {
	return n.val != nil
}

// Get returns the 'n's value and 'true', if 'n' has value.
// If 'n' has no value 'false' and 'false' are returned.
func (n *Bool) Get() (bool, bool) {
	if n.val == nil {
		return false, false
	}
	return *n.val, true
}

// GetMust returns Get's result, if 'n' has a value.
// If 'n' has no value, 'false' is returned.
func (n *Bool) GetMust() bool {
	r, ok := n.Get()
	if !ok {
		return false
	}
	return r
}

// Set sets the 'n's value to 'v'.
func (n *Bool) Set(v bool) {
	n.val = &v
}

// Null sets 'n' as having no value.
func (n *Bool) Null() {
	n.val = nil
}

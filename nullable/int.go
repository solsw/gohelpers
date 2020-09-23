package nullable

// Int represents an int type that may have no value.
type Int struct {
	val *int
}

// NullInt creates a new Int without value.
func NullInt() Int {
	return Int{nil}
}

// NewInt creates a new Int with value 'v'.
func NewInt(v int) Int {
	return Int{&v}
}

// Has reports whether n has a value.
func (n *Int) Has() bool {
	return n.val != nil
}

// Get returns (n's value, true) if n has value.
// Returns (0, false) if n has no value.
func (n *Int) Get() (int, bool) {
	if n.val == nil {
		return 0, false
	}
	return *n.val, true
}

// MustGet is like Get but returns 0 if n has no value.
func (n *Int) MustGet() int {
	r, ok := n.Get()
	if !ok {
		return 0
	}
	return r
}

// Set sets the n's value to 'v'.
func (n *Int) Set(v int) {
	n.val = &v
}

// Null sets n as having no value.
func (n *Int) Null() {
	n.val = nil
}

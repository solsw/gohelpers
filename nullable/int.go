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

// Has reports whether 'n' has a value.
func (n *Int) Has() bool {
	return n.val != nil
}

// Get returns the 'n's value and 'true', if 'n' has value.
// If 'n' has no value 0 and 'false' are returned.
func (n *Int) Get() (int, bool) {
	if n.val == nil {
		return 0, false
	}
	return *n.val, true
}

// GetMust returns Get's result, if 'n' has a value.
// If 'n' has no value, 0 is returned.
func (n *Int) GetMust() int {
	r, ok := n.Get()
	if !ok {
		return 0
	}
	return r
}

// Set sets the 'n's value to 'v'.
func (n *Int) Set(v int) {
	n.val = &v
}

// Null sets 'n' as having no value.
func (n *Int) Null() {
	n.val = nil
}

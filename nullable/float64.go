package nullable

// Float64 represents a float64 type that may have no value.
type Float64 struct {
	val *float64
}

// NullFloat64 creates a new Float64 without value.
func NullFloat64() Float64 {
	return Float64{nil}
}

// NewFloat64 creates a new Float64 with value 'v'.
func NewFloat64(v float64) Float64 {
	return Float64{&v}
}

// Has reports whether n has a value.
func (n *Float64) Has() bool {
	return n.val != nil
}

// Get returns (n's value, true) if n has value.
// Returns (0.0, false) if n has no value.
func (n *Float64) Get() (float64, bool) {
	if n.val == nil {
		return 0.0, false
	}
	return *n.val, true
}

// MustGet is like Get but returns 0.0 if n has no value.
func (n *Float64) MustGet() float64 {
	r, ok := n.Get()
	if !ok {
		return 0.0
	}
	return r
}

// Set sets the n's value to 'v'.
func (n *Float64) Set(v float64) {
	n.val = &v
}

// Null sets n as having no value.
func (n *Float64) Null() {
	n.val = nil
}

package nullable

// String represents a string type that may have no value.
type String struct {
	val *string
}

// NilString creates a new String without value.
func NilString() String {
	return String{nil}
}

// NewString creates a new String with value 'v'.
func NewString(v string) String {
	return String{&v}
}

// Has reports whether n has a value.
func (n *String) Has() bool {
	return n.val != nil
}

// Get returns (n's value, true) if n has value.
// Returns (empty string, false) if n has no value.
func (n *String) Get() (string, bool) {
	if n.val == nil {
		return "", false
	}
	return *n.val, true
}

// MustGet is like Get but returns an empty string if n has no value.
func (n *String) MustGet() string {
	r, ok := n.Get()
	if !ok {
		return ""
	}
	return r
}

// Set sets the n's value to 'v'.
func (n *String) Set(v string) {
	n.val = &v
}

// Null sets n as having no value.
func (n *String) Null() {
	n.val = nil
}

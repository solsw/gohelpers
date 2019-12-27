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

// Has reports whether 'n' has a value.
func (n *String) Has() bool {
	return n.val != nil
}

// Get returns the 'n's value and 'true', if 'n' has value.
// If 'n' has no value empty string and 'false' are returned.
func (n *String) Get() (string, bool) {
	if n.val == nil {
		return "", false
	}
	return *n.val, true
}

// GetMust returns Get's result, if 'n' has a value.
// If 'n' has no value, empty string is returned.
func (n *String) GetMust() string {
	r, ok := n.Get()
	if !ok {
		return ""
	}
	return r
}

// Set sets the 'n's value to 'v'.
func (n *String) Set(v string) {
	n.val = &v
}

// Null sets 'n' as having no value.
func (n *String) Null() {
	n.val = nil
}

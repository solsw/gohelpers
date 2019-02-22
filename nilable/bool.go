package nilable

// Bool represents a bool type that can be assigned nil.
type Bool struct {
	val bool
	has bool
}

// NilBool creates a new Bool without value.
func NilBool() Bool {
	return Bool{}
}

// NewBool creates a new Bool with value 'v'.
func NewBool(v bool) Bool {
	return Bool{val: v, has: true}
}

// Has reports whether the Bool has a value.
func (n Bool) Has() bool {
	return n.has
}

// Get returns the Bool's value.
func (n Bool) Get() bool {
	return n.val
}

// Set sets the Bool's value to 'v'.
func (n Bool) Set(v bool) {
	n.val = v
	n.has = true
}

// Nil sets the Bool's value to nil.
func (n Bool) Nil() {
	n.val = false
	n.has = false
}

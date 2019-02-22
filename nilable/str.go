package nilable

// Str represents a string type that can be assigned nil.
type Str struct {
	val string
	has bool
}

// NilStr creates a new Str without value.
func NilStr() Str {
	return Str{}
}

// NewStr creates a new Str with value 'v'.
func NewStr(v string) Str {
	return Str{val: v, has: true}
}

// Has reports whether the Str has a value.
func (n Str) Has() bool {
	return n.has
}

// Get returns the Str's value.
func (n Str) Get() string {
	return n.val
}

// Set sets the Str's value to 'v'.
func (n Str) Set(v string) {
	n.val = v
	n.has = true
}

// Nil sets the Str's value to nil.
func (n Str) Nil() {
	n.val = ""
	n.has = false
}

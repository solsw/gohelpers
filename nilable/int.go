package nilable

// Int represents an int type that can be assigned nil.
type Int struct {
	val int
	has bool
}

// NilInt creates a new Int without value.
func NilInt() Int {
	return Int{}
}

// NewInt creates a new Int with value 'v'.
func NewInt(v int) Int {
	return Int{val: v, has: true}
}

// Has reports whether the Int has a value.
func (n Int) Has() bool {
	return n.has
}

// Get returns the Int's value.
func (n Int) Get() int {
	return n.val
}

// Set sets the Int's value to 'v'.
func (n Int) Set(v int) {
	n.val = v
	n.has = true
}

// Nil sets the Int's value to nil.
func (n Int) Nil() {
	n.val = 0
	n.has = false
}

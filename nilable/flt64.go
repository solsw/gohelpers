package nilable

// Flt64 represents a float64 type that can be assigned nil.
type Flt64 struct {
	val float64
	has bool
}

// NilFlt64 creates a new Flt64 without value.
func NilFlt64() Flt64 {
	return Flt64{}
}

// NewFlt64 creates a new Flt64 with value 'v'.
func NewFlt64(v float64) Flt64 {
	return Flt64{val: v, has: true}
}

// Has reports whether the Flt64 has a value.
func (n Flt64) Has() bool {
	return n.has
}

// Get returns the Flt64's value.
func (n Flt64) Get() float64 {
	return n.val
}

// Set sets the Flt64's value to 'v'.
func (n Flt64) Set(v float64) {
	n.val = v
	n.has = true
}

// Nil sets the Flt64's value to nil.
func (n Flt64) Nil() {
	n.val = 0
	n.has = false
}

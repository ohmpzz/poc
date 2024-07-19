package circle // Private package for Triangle (optional)

// Triangle struct implements the Shape interface
type Circle struct {
	Foo float64
	Bar float64
}

// Area method for Triangle
func (t *Circle) Area() float64 {
	return 0.5 * t.Foo * t.Bar
}

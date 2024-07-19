package geometry // Private package for Triangle (optional)

// Triangle struct implements the Shape interface
type Triangle struct {
	Base   float64
	Height float64
}

// Area method for Triangle
func (t *Triangle) Area() float64 {
	return 0.5 * t.Base * t.Height
}

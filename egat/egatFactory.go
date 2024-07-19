package egat

import "fmt"

type Foo struct {
	Load, Rate float64
	Models     string
}

func (e *Foo) Area() float64 {
	return 1.1
}

func (e *Foo) Masure() float64 {
	fmt.Println(e.Area())

	return 2
}

package main

import "fmt"
import "math"

type Shape interface {
    Area() float64
    Perimeter() float64
}

type Rectangle struct{ W, H float64 }
func (r Rectangle) Area() float64      { 
    return r.W * r.H 
}
func (r Rectangle) Perimeter() float64 { 
    return 2 * (r.W + r.H) 
}

type Circle struct{ R float64 }
func (c Circle) Area() float64      { 
    return math.Pi * c.R * c.R 
}
func (c Circle) Perimeter() float64 { 
    return 2 * math.Pi * c.R 
}

type Triangle struct{ A, B, C float64 }
func (t Triangle) Area() float64 {
    p := (t.A + t.B + t.C) / 2
    return math.Sqrt(p * (p - t.A) * (p - t.B) * (p - t.C))
}
func (t Triangle) Perimeter() float64 { return t.A + t.B + t.C }

func main() {
    var x Shape = Rectangle{W: 2, H: 7}
    fmt.Println(x.Area(), x.Perimeter())
    x = Circle{R: 8}
    fmt.Println(x.Area(), x.Perimeter())
    x = Triangle{A: 9, B: 6, C: 3}
    fmt.Println(x.Area(), x.Perimeter())
}
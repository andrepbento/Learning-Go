package main

import (
    "fmt"
    "math"
)

// define an interface
type Shape interface {
    area() float64
}

// define a circle
type Circle struct {
    x, y, radius float64
}

// define a rectangle
type Rectangle struct {
    width, height float64
}

// define a method for circle (implementation of Shape.area())
func(circle Circle) area() float64 {
    return math.Pi * circle.radius * circle.radius
}

// define a method for rectangl (implementation of Shape.area())
func(rect Rectangle) area() float64 {
    return rect.width * rect.height
}

// define a method for Shape
func getArea(shape Shape) float64 {
    return shape.area()
}

func main() {
    /*      interface definition
        type interface_name interface {
            method_name1 [return_type]
            method_name2 [return_type]
            method_name3 [return_type]
            ...
            method_nameN [return_type]
        }

            struct definition
        type struct_name struct {
            variables_xpto...
        }

            interface methods implementation
        func (struct_name_variable struct_name) method_name1() [return_type] {
            ...
        }
        ...
        func (struct_name_variable struct_name) method_namen() [return_type] {
            ...
        }
    */
    
    circle := Circle{x:0, y:0, radius:5}
    rectangle := Rectangle{width:10, height:5}

    fmt.Printf("Circle area : %.2f\n", getArea(circle))
    fmt.Printf("Rectangle area : %.2f\n", getArea(rectangle))
}
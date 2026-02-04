// module

package shape

import (
	"fmt"
	"math"
)

// interface
type ShapeWithArea interface {
	Area() float32
}

type ShapeWithPerimetr interface {
	Perimetr() float32
}

// композиция интерфесов
type Shape interface {
	ShapeWithArea
	ShapeWithPerimetr
}

// реализации
type Square struct {
	sideLength float32
}

func NewSquare(sideLength float32) *Square {
	return &Square{
		sideLength: sideLength,
	}
}

func (s Square) Area() float32 {
	return s.sideLength * s.sideLength
}
func (s Square) Perimetr() float32 {
	return s.sideLength * 2
}

func (s Square) GetSideLength() float32 {
	return s.sideLength
}

func (s *Square) SetSideLength(newSideLength float32) {
	s.sideLength = newSideLength
}

type Circle struct {
	radius float32
}

func NewCircle(radius float32) *Circle {
	return &Circle{
		radius: radius,
	}
}

func (c Circle) Area() float32 {
	return math.Pi * c.radius * c.radius
}
func (c Circle) Perimetr() float32 {
	return math.Pi * c.radius
}
func (s Circle) GetRadius() float32 {
	return s.radius
}

func (s *Circle) SetRadius(newRadius float32) {
	s.radius = newRadius
}

func main() {

	square := Square{4.1}
	circle := Circle{3}
	printShapeArea(square)
	printShapeArea(circle)

	// empty interface
	printInterface1(square)
	printInterface1(circle)
	printInterface2(square)
	printInterface2(circle)

	// приведение типов интерфесов
	printInterface3("asd")
	printInterface3(1)
	printInterface3(1.)
	printInterface3(3 / 1.)
	printInterface3(nil)
	printInterface3(circle)

	// приведение типов интерфесов с проверкой
	printInterface4(12313123)
	printInterface4("circle")

}

func printShapeArea(shape Shape) {
	fmt.Println(shape.Area())
}

// empty interface - у него нет "поведения", т.е. не декларирована ни одна функция. А значит любой тип реализует этот интерфейс. interface{} == any
func printInterface1(i interface{}) {
	fmt.Println(i)
}

// тоже empty interface
func printInterface2(i any) {
	fmt.Printf("%+v\n", i)
}

// тоже empty interface НО С ПРИВЕДЕНИЕМ ТИПОВ
func printInterface3(i any) {
	switch t := i.(type) {
	case int:
		fmt.Printf("Это int %+v\n", t)
	case string:
		fmt.Printf("Это string %+v\n", t)
	case float64:
		fmt.Printf("Это float64 %+v\n", t)
	case Shape:
		fmt.Printf("Это Shape %+v\n", t)
	default:
		fmt.Println("unknown type")
	}
}

// приведение типов
func printInterface4(i any) {

	//приведение к string
	// str := i.(string)

	//или приведение к string с проверкой
	str, ok := i.(string)
	if !ok {
		fmt.Println(ok, " : ", i)
		return
	}
	fmt.Println(ok, " : ", len(str))
}

//пакеты и модули

package main

import (
	"fmt"
	"modules/shape"
)

func main() {

	//объявление и инициализация с помощью конструктора
	circleMain := shape.NewCircle(5) // - получение указателя(ссылки(*)(адреса)) на структуру(объект) и заполнение объекта типа shape.Circle
	fmt.Println(circleMain)          // - печать ЗНАЧЕНИЙ структуры - &{5} - строковое представление указателя на структуру

	circle1 := &circleMain          // - получение указателя на ссылку circleMain (circle1 хранит адрес переменной circleMain, а не адрес самой структуры)
	fmt.Printf("%p\n", *circle1)    // - получаем ЗНАЧЕНИЕ указателя - ссылку на значение circleMain в куче, т.е. ссылка на адрес переменной circleMain - 0xc00010c060
	fmt.Printf("%p\n", &circle1)    // - получаем АДРЕС переменной circle1 - 0xc00010e040
	fmt.Printf("%p\n", circle1)     // - получаем ЗНАЧЕНИЕ указателя - адрес circleMain - 0xc00010e030
	fmt.Printf("%p\n", circleMain)  // - печать ЗНАЧЕНИЙ - адрес структуры Circle в куче - 0xc00010c060
	fmt.Printf("%p\n", &circleMain) // - печать АДРЕС - адрес circleMain - 0xc00010e030
	(*circle1).SetRadius(1)
	fmt.Println(circle1) // - получаем ЗНАЧЕНИЕ указателя - адрес circleMain - 0xc00010e030
	fmt.Println("<=======First_Example=======>")

	circle2 := *circleMain
	fmt.Println(circle2) //-радиус как и у circle1, так как ссылка circle1 == circleMain
	circle2.SetRadius(2) //-новый радиус
	fmt.Println(circle2)

	fmt.Println(circleMain == &circle2) //-ссылки разные
	fmt.Println(*circleMain == circle2) //-структуры одинаковые

	fmt.Println("<=======New_Example=======>")
	circleAnother := shape.NewCircle(5)
	fmt.Println(circleAnother)
	fmt.Println(*circleAnother == *circleMain)
	fmt.Printf("%p\n", circleMain)
	fmt.Printf("%p\n", circleAnother)

	//объявление и инициализация с помощью сеттера
	square := shape.Square{}
	square.SetSideLength(5)
	fmt.Println(square)
}

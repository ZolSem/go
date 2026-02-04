/*----- stuctures-----*/
package main

import "fmt"

func main() {

	user := struct {
		name   string
		age    int
		sex    string
		weight float32
		height float32
	}{"Vasya", 31, "male", 85.1, 178.2}

	fmt.Println(user)
	fmt.Printf("%+v\n", user)

	//or
	user1 := User{"Kolya", 45, "male", 81.7, 178.9}
	fmt.Println(user1)
	fmt.Printf("%+v\n", user1)

	user2 := NewUser("Zhanna", "female", 21, 55.3, 169)
	fmt.Println(user2)
	fmt.Printf("%+v\n", user2)

	//getter
	fmt.Println(user2.getName())

	//setter
	user2.setName("Ilona")
	fmt.Println(user2.getName())
	fmt.Printf("%+v\n", user2)

}

//or как класс - struct
type User struct {
	name   string
	age    int
	sex    string
	weight float32
	height float32
}

//конструктор
func NewUser(name, sex string, age int, weight, height float32) User {
	return User{
		name:   name,
		age:    age,
		sex:    sex,
		weight: weight,
		height: height}
}

//getter
func (u User) getName() string {
	return u.name
}

//setter СО ЗНАКОМ * ДЛЯ ОБРАЩЕНИЯ ПО ССЫЛКЕ
func (u *User) setName(newName string) {
	u.name = newName
}

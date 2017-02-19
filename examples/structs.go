package main

import "fmt"

type person struct {
	name string
	age int
}

func main() {
	fmt.Println(person{"Bell", 20}) // given name and age
	fmt.Println(person{ age: 20}) // only age
	fmt.Println(&person{"Bell", 20}) // address/pointer of the struct

	s := person{age: 20}
	fmt.Println(s.age) // this is okay
	fmt.Println("Name: ", s.name) // no error, but empty

	copyOfS := s
	fmt.Println(copyOfS.age)
	copyOfS.name = "Test"
	fmt.Println(s.name) // still empty

	ptrOfS := &s
	ptrOfS.name = "Test"
	fmt.Println(s.name) // wolla !

}

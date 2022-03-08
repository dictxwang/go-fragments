package _interface

import "fmt"

type Animal interface {
	Speak() string
}

// 实现了接口的所有方法，即认为是改接口的子类型
type Dog struct {
}

type Cat struct {
}

func (d *Dog)Speak() string {
	return "Woof"
}

func (c *Cat)Speak() string {
	return "Meow"
}

func f(animal Animal) {
	fmt.Println(animal.Speak())
}

func SampleMain()  {

	fmt.Println("\n[interface_sample]")

	d := &Dog{}
	f(d)

	c := Cat{}
	f(&c)

	animals := []Animal{&Dog{},&Cat{}}
	for _,animal := range animals {
		fmt.Println(animal.Speak())
	}
}

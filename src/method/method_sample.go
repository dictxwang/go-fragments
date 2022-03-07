package method

import "fmt"

type Person struct {
	name string
}

// 非指针传递，相当于复制了一个person，执行完后外层的person属性不会变化
func (p Person) ChangeName(name string) {
	p.name = name
}

// 指针传递，会引起外层的person属性变化
func (p *Person) ChangeNameWithPointer(name string) {
	p.name = name
}

func (p Person) GetName() string {
	return p.name
}

func SampleMain()  {

	fmt.Println("\n[method_sample]")
	person := Person{}
	person.ChangeName("liudehua")
	fmt.Println(person.GetName())
	person.ChangeNameWithPointer("liudehua")
	fmt.Println(person.GetName())
}
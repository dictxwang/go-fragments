package _reflect

import (
	"fmt"
	"math/rand"
	"reflect"
)

type t1 struct {
	A string
	B int
}

func (t t1) FnA() {
	fmt.Println("this is t1 FnA executing...")
}

func (t t1) FnB(name string, value int) {
	fmt.Printf("this is t1 FnB executing, %s=%d\n", name, value)
}

func (t t1) FnC() (bool, int) {
	return true, rand.Int()
}

func SampleMain()  {
	fmt.Println("\n[reflect_sample]")
	var num = 123
	var i interface{} = num
	fmt.Println("type:", reflect.TypeOf(i))
	fmt.Println("value:", reflect.ValueOf(i))

	var j = 233
	p := reflect.ValueOf(&j)
	v := reflect.ValueOf(j)
	fmt.Println(p, v)

	st := t1{
		A: "s",
		B: 123,
	}

	fmt.Println("st type:", reflect.TypeOf(st).Name())
	fmt.Println("st value:", reflect.ValueOf(st))

	/* 反射获取属性和方法时，或忽略非导出（私有）类型的属性和方法 */

	// 遍历属性
	stt := reflect.TypeOf(st)
	stv := reflect.ValueOf(st)
	for i := 0; i < stt.NumField(); i++ {
		field := stt.Field(i)
		value := stv.Field(i).Interface()
		fmt.Printf("field's name:%s, type:%s, value: %v\n", field.Name, field.Type, value)
	}

	// 遍历方法
	for i := 0; i < stt.NumMethod(); i++ {
		m := stt.Method(i)
		fmt.Printf("method's name:%s, type:%s\n", m.Name, m.Type)
	}

	// 修改属性
	// 需要通过指针实现，否则修改会panic
	stm := reflect.ValueOf(&st)
	// 调用Elem()获取指针包含的对象
	stm.Elem().FieldByName("B").SetInt(1000)
	fmt.Println("after modify B:", st.B)

	// 调用方法
	methodFnA := stv.MethodByName("FnA")
	// 无参数的方法
	argsFnA := make([]reflect.Value, 0)
	methodFnA.Call(argsFnA)

	methodFnB := stv.MethodByName("FnB")
	// 带参数的方法
	argsFnB := []reflect.Value{reflect.ValueOf("mathematics"), reflect.ValueOf(100)}
	methodFnB.Call(argsFnB)

	// 获取返回值
	methodFnC := stv.MethodByName("FnC")
	argsFnC := make([]reflect.Value, 0)
	results := methodFnC.Call(argsFnC)
	for i,r := range results {
		fmt.Printf("reflect method result: item%d = %v\n", i, r)
	}
}
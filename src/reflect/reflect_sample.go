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
	fmt.Println(p, v)  // 0xc000016058 233

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

	// 使用reflect进行深度比较
	a1 := []int{1,2,3}
	a2 := []int{1,2,3}
	fmt.Printf("a1 deepEqual a2: %v\n", reflect.DeepEqual(a1, a2))
}

func SampleMainType() {
	fmt.Println("\n[reflect_sample_type]")

	// 创建len和cap都是6的切片
	slice1 := make([]int, 6, 6)
	slice2 := []int {1,2,3}
	fmt.Println(reflect.TypeOf(slice1))  // []int
	fmt.Println(reflect.TypeOf(slice2))  // []int
	fmt.Println(reflect.ValueOf(slice1).Type())  // []int
	fmt.Println(reflect.ValueOf(slice1).Kind())  // slice
	// Elem的执行对象需要是指针类型
	//fmt.Println(reflect.ValueOf(slice2).Elem())  // panic: reflect: call of reflect.Value.Elem on slice Value
	fmt.Println(reflect.TypeOf(slice2).Elem())  // int

	// 转换成指针操作
	slicePoint1 := &slice1
	fmt.Println(reflect.TypeOf(slicePoint1))  // *[]int
	fmt.Println(reflect.TypeOf(slicePoint1).Elem())  // []int
	fmt.Println(reflect.TypeOf(slicePoint1).Elem().Elem())  // int

	array1 := [3]int {1,2,3}
	array2 := [...]int {1,2}
	array3 := [...]int {1,2,5:6} // len=6
	fmt.Println(reflect.TypeOf(array1))  // [3]int
	fmt.Println(reflect.ValueOf(array1).Type())  // [3]int
	fmt.Println(reflect.ValueOf(array1).Kind())  // array
	fmt.Println(reflect.TypeOf(array2))  // [2]int
	fmt.Println(reflect.TypeOf(array3))  // [6]int
	fmt.Println(reflect.TypeOf(array3).Elem())  // int

	arrayPoint1 := &array1
	arrayPoint1[1] = 100
	fmt.Println(reflect.TypeOf(arrayPoint1))  // *[3]int
	fmt.Println(*arrayPoint1)  // [1 100 3]

	t := t1{}
	fmt.Println(reflect.TypeOf(t))  // _reflect.t1
	fmt.Println(reflect.TypeOf(t).Kind())  // struct
	fmt.Println(reflect.ValueOf(t).Type())  // _reflect.t1
	fmt.Println(reflect.ValueOf(t).Kind())  // struct
}
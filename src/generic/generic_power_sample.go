package generic

import (
	"fmt"
	"reflect"
)

func verifyFuncSignature(fn reflect.Value, types ...reflect.Type) bool {

	// 检查是function
	if fn.Kind() != reflect.Func {
		return false
	}

	// 检查入参和出参数量
	if fn.Type().NumIn() != len(types) - 1 || fn.Type().NumOut() != 1 {
		return false
	}

	// 检查参数和方法参数类型是否能对上
	for i := 0; i < len(types) - 1; i++ {
		if fn.Type().In(i) != types[i] {
			return false
		}
	}

	outType := types[len(types) - 1]
	if outType != nil && fn.Type().Out(0) != outType {
		return false
	}
	return true
}

func transform(slice, function interface{}, inPlace bool) interface{} {

	sliceInType := reflect.ValueOf(slice)
	if sliceInType.Kind() != reflect.Slice {
		panic("transform: not slice")
	}

	// 检查方法签名
	fn := reflect.ValueOf(function)
	elemType := sliceInType.Type().Elem()
	if !verifyFuncSignature(fn, elemType, nil) {
		panic("transform: function must be of type func (" + sliceInType.Elem().String() + ") outputElemType")
	}

	sliceOutType := sliceInType
	// 如果是就地计算，则将计算结果存入原slice；否则就重新创建一个slice保存计算结果
	if !inPlace {
		sliceOutType = reflect.MakeSlice(reflect.SliceOf(fn.Type().Out(0)), sliceInType.Len(), sliceInType.Len())
	}
	for i := 0; i < sliceInType.Len(); i++ {
		sliceOutType.Index(i).Set(fn.Call([]reflect.Value{sliceInType.Index(i)})[0])
	}
	return sliceOutType.Interface()
}

func Transform(slice, function interface{}) interface{} {
	return transform(slice, function, false)
}

func TransformInPlace(slice, function interface{}) interface{} {
	return transform(slice, function, true)
}

func Reduce(slice, zero, pairFunc interface{}) interface{} {
	sliceInType := reflect.ValueOf(slice)
	if sliceInType.Kind() != reflect.Slice {
		panic("reduce: wrong type, not slice")
	}
	len := sliceInType.Len()
	if len == 0 {
		return zero
	} else if len == 1 {
		return sliceInType.Index(0)
	}

	elemType := sliceInType.Type().Elem()
	fn := reflect.ValueOf(pairFunc)
	if !verifyFuncSignature(fn, elemType, elemType, elemType) {
		t := elemType.String()
		panic("reduce: function must be of type func (" + t + "," + t + ")" + t)
	}

	var ins [2] reflect.Value
	ins[0] = sliceInType.Index(0)
	ins[1] = sliceInType.Index(1)
	out := fn.Call(ins[:])[0]

	for i := 2; i < len; i++ {
		ins[0] = out
		ins[1] = sliceInType.Index(i)
		out = fn.Call(ins[:])[0]
	}
	return out.Interface()
}



func SampleMainPower()  {

	fmt.Println("\n[generic_sample_power]")

	// 调用map
	list1 := []int{1,2,3,4,5,6,7,8,9}
	TransformInPlace(list1, func(a int) int {
		return a * 3
	})
	fmt.Println(list1)

	list2 := [] string {"a", "b", "c", "d", "e"}
	result2 := Transform(list2, func(a string) string {
		return a + a + a
	})
	fmt.Println(result2)

	// 调用reduce
	reduceResult1 := Reduce(list1, 0, func(a, b int) int {
		return a + b
	})
	fmt.Println("reduce result: ", reduceResult1)

	reduceResult2 := Reduce(list2, "", func(a, b string) string {
		return a + b
	})
	fmt.Println("reduce result2: ", reduceResult2)
}

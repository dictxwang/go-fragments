package generic

import (
	"fmt"
	"reflect"
)

// 实现泛型模式的map\reduce\filter，预计interface{}和reflect实现

func SimpleMap(data interface{}, fn interface{}) [] interface{} {

	vfn := reflect.ValueOf(fn)
	vdata := reflect.ValueOf(data)
	result := make([]interface{}, vdata.Len())

	for i := 0; i < vdata.Len(); i++ {
		result[i] = vfn.Call([]reflect.Value{vdata.Index(i)})[0].Interface()
	}
	return result
}

func SimpleReduce(data interface{}, fn interface{}) interface{} {

	vfn := reflect.ValueOf(fn)
	vdata := reflect.ValueOf(data)
	var ins [2]reflect.Value
	// 如果是数值型，需要Elem()获取值，如果是字符串则不需要
	ins[0] = vdata.Index(0).Elem()
	ins[1] = vdata.Index(1).Elem()
	result := vfn.Call(ins[:])[0]

	for i := 2; i < vdata.Len(); i++ {
		ins[0] = result
		ins[1] = vdata.Index(i).Elem()
		result = vfn.Call(ins[:])[0]
	}
	return result.Interface()
}

func SimpleFilter(data interface{}, fn interface{}) interface{} {

	vfn := reflect.ValueOf(fn)
	vdata := reflect.ValueOf(data)
	var result = []interface{}{}
	for i := 0; i < vdata.Len(); i++ {
		if vfn.Call([]reflect.Value{vdata.Index(i).Elem()})[0].Bool() {
			result = append(result, vdata.Index(i).Interface())
		}
	}
	return result
}

func SampleMainSimple() {
	fmt.Println("\n[generic_sample]")

	nums := []int{1,2,3,4,5}
	square := func(x int) int {
		return x * x
	}
	squareResult := SimpleMap(nums, square)
	fmt.Println("map result: ", squareResult)

	reduceResult := SimpleReduce(squareResult, func(a int, b int) int {
		return a + b
	})
	fmt.Println("reduce result: ", reduceResult)

	//reduceResult2 := SimpleReduce([]string{"a", "b", "c"}, func(a string, b string) string {
	//	return strings.ToUpper(a) + "_" + strings.ToUpper(b)
	//})
	//fmt.Println("reduce result: ", reduceResult2)

	filterResult := SimpleFilter(squareResult, func(a int) bool {
		return a > 10
	})
	fmt.Println("filter result: ", filterResult)
}

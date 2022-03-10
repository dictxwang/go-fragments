package generic

import (
	"fmt"
	"reflect"
	"strings"
)

// 实现泛型模式的map\reduce\filter，预计interface{}和reflect实现

func SimpleMap(data, fn interface{}) [] interface{} {

	vfn := reflect.ValueOf(fn)
	vdata := reflect.ValueOf(data)
	result := make([]interface{}, vdata.Len())

	for i := 0; i < vdata.Len(); i++ {
		result[i] = vfn.Call([]reflect.Value{vdata.Index(i)})[0].Interface()
	}
	return result
}

func SimpleReduce(data, fn interface{}) interface{} {

	vfn := reflect.ValueOf(fn)
	// 注意这里的data类型需要是slice，如果是array会在Call环节包类型异常
	vdata := reflect.ValueOf(data)
	var ins [2]reflect.Value
	ins[0] = vdata.Index(0)
	ins[1] = vdata.Index(1)
	result := vfn.Call(ins[:])[0]

	for i := 2; i < vdata.Len(); i++ {
		ins[0] = result
		ins[1] = vdata.Index(i)
		result = vfn.Call(ins[:])[0]
	}
	return result.Interface()
}

func SimpleFilter(data, fn interface{}) interface{} {

	vfn := reflect.ValueOf(fn)
	vdata := reflect.ValueOf(data)
	var result = []interface{}{}
	for i := 0; i < vdata.Len(); i++ {
		if vfn.Call([]reflect.Value{vdata.Index(i)})[0].Bool() {
			result = append(result, vdata.Index(i).Interface())
		}
	}
	return result
}

func SampleMainSimple() {
	fmt.Println("\n[generic_sample]")

	nums := []int{1,2,3,4,5,6,7,8}
	texts := []string{"a", "b", "c"}
	square := func(x int) int {
		return x * x
	}
	squareResult := SimpleMap(nums, square)
	fmt.Println("map result: ", squareResult)

	multiResult := SimpleMap(texts, func(x string) string {
		return x + x
	})
	fmt.Println("map result: ", multiResult)

	// nums需要是slice，不能是array
	reduceResult := SimpleReduce(nums, func(a int, b int) int {
		return a + b
	})
	fmt.Println("reduce result: ", reduceResult)

	// texts需要是slice，不能是array
	reduceResult2 := SimpleReduce(texts, func(a string, b string) string {
		return strings.ToUpper(a) + "_" + strings.ToUpper(b)
	})
	fmt.Println("reduce result: ", reduceResult2)

	// nums需要是slice，不能是array
	filterResult := SimpleFilter(nums, func(a int) bool {
		return a > 3
	})
	fmt.Println("filter result: ", filterResult)

	// nums需要是slice，不能是array
	filterResult2 := SimpleFilter(texts, func(a string) bool {
		return a != "a"
	})
	fmt.Println("filter result2: ", filterResult2)
}

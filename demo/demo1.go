package demo

import (
	"encoding/json"
	"fmt"
)

var n = 1

const (
	N1 = iota
	N2
)

func Demo1(f string, args ...interface{}) {
	fmt.Printf("Demo1 ......")
	fmt.Printf(f, args)
}

func Demo2() {
	arr := make([]int, 10, 10)
	for i := 0; i < len(arr); i++ {
		arr[i] = i
	}
	for index, value := range arr {
		fmt.Println(index)
		fmt.Println(value)
	}

	// 解析json数据
	str := `{
			"name":"test",
			"product_id":"1",
			"number":"110011",
			"price":"0.01",
			"is_on_sale":"true"
		}`
	var p interface{}
	err := json.Unmarshal([]byte(str), &p)
	if err != nil {
		return
	}
	// 现在interface{}解析出里面的数据
	m := p.(map[string]interface{})
	for k, v := range m {
		switch vv := v.(type) {
		case string:
			fmt.Printf("%s is string, value: %s\n", k, vv)
		case int:
			fmt.Printf("%s is int, value: %d\n", k, vv)
		case int64:
			fmt.Printf("%s is int64, value: %d\n", k, vv)
		case bool:
			fmt.Printf("%s is bool, vaule: %v", k, vv)
		default:
			fmt.Printf("%s is unknow type\n", k)
		}
	}
}

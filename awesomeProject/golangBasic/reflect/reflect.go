package main

import (
	"fmt"
	"reflect"
)

/*
   反射demo
**/

type Student struct{
	Name string `json:"name" ini:"name"`
	Age int `json:"age" ini:"age"`
}

func (s *Student) Say(){
	fmt.Println("say......")
}

func (s Student) Hello(){
	fmt.Println("hello.....")
}
func main(){
	s := Student{
		Name: "xiaoming",
		Age: 18,
	}
	// 获取方法
	t := reflect.TypeOf(s)
	vPtr := reflect.ValueOf(&s)
	v := reflect.ValueOf(s)
	fmt.Println(vPtr.NumMethod())
	for i := 0; i < vPtr.NumMethod(); i++ {
		methodType := vPtr.Method(i).Type()
		fmt.Printf("method:%s\n", methodType)
		// 通过反射调用方法传递的参数必须是 []reflect.Value 类型
		var args []reflect.Value
		vPtr.Method(i).Call(args)
	}
	m1, _ := t.MethodByName("Hello")
	fmt.Printf("type:%T, name:%v\n", m1.Type, m1.Name)
	//
	m2 := vPtr.MethodByName("Say")
	fmt.Println(m2)

	// 获取Field
	for i := 0; i < v.NumField();i++ {
		f := t.Field(i)
		fValue := v.Field(i)
		fmt.Printf("type:%v, name:%v, value:%v\n", f.Type.Name(), f.Name, fValue)
	}
}

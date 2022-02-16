package main

import (
	"fmt"
	"strconv"
)

func main(){
	// ->str
	// Itoa
	fmt.Println("conv to str.....")
	i2 := 200
	s00 := strconv.Itoa(i2)
	fmt.Printf("type:%T value:%#v\n", s00, s00) //type:string value:"200"

	// Format
	s1 := strconv.FormatBool(true)
	fmt.Printf("type:%T value:%v\n", s1, s1)
	s2 := strconv.FormatFloat(3.1415, 'E', -1, 64)
	fmt.Printf("type:%T value:%v\n", s2, s2)
	s3 := strconv.FormatInt(-2, 16)
	fmt.Printf("type:%T value:%v\n", s3, s3)
	s4 := strconv.FormatUint(2, 16)
	fmt.Printf("type:%T value:%v\n", s4, s4)

	// str->
	fmt.Println("conv str to .....")
	// Atoi
	s0 := "100"
	i1, err := strconv.Atoi(s0)
	if err != nil {
		fmt.Println("can't convert to int")
	} else {
		fmt.Printf("type:%T value:%#v\n", i1, i1) //type:int value:100
	}

	// Parse
	b, err := strconv.ParseBool("true")
	fmt.Printf("type:%T value:%v\n", b, b)
	f, err := strconv.ParseFloat("3.1415", 64)
	fmt.Printf("type:%T value:%v\n", f, f)
	i, err := strconv.ParseInt("-2", 10, 64)
	fmt.Printf("type:%T value:%v\n", i, i)
	u, err := strconv.ParseUint("2", 10, 64)
	fmt.Printf("type:%T value:%v\n", u, u)
}

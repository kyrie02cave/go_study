package snow

import (
	"fmt"
	"practice/package/calc" // module name:go mod init  ---> go mod edit -replace
)

func CallMathOps(){
	a := calc.Add(1.0, 2.0)
	b := calc.Sub(2.0, 1.0)
	c, err := calc.Multi(1.0, 2.0)
	if err != nil{
		return
	}
	d, err := calc.Div(2.0, 1.0)
	if err != nil{
		return
	}
	fmt.Printf("add:%v sub:%v multi:%v div:%v\n", a, b, c, d)
}

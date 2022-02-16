package process_control

import "fmt"

// 编写代码打印9*9乘法表

func PrintMultiTable(){
	const n = 9
	for i := 1; i <= n; i++{
		for j := 1; j <= i; j++{
			fmt.Printf("|%d * %d|", j, i)
		}
		fmt.Println()
	}
}
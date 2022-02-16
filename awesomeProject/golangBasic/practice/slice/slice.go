package slice

import "fmt"

// AppendSlice 请写出下面代码的输出结果。
func AppendSlice() {
	var a = make([]string, 5, 10)// [0 0 0 0 0 | 0 0 0 0 0]
	var aP *[]string
	for i := 0; i < 10; i++ {
		aP = &a
		a = append(a, fmt.Sprintf("%v", i)) // [0 0 0 0 0 | 0 1 2 3 4] --> [0 0 0 0 0 0 1 2 3 4 5 6 7 8 9 0 0 0 0 0]
	}
	fmt.Println(a)
	fmt.Printf("%p\n", aP)
}



// 请使用内置的sort包对数组var a = [...]int{3, 7, 8, 9, 1}进行排序（附加题，自行查资料解答）


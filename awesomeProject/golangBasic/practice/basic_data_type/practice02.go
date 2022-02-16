package basic_data_type

import (
	"fmt"
	"unicode"
)

// CountChineseWord： 统计字符串中汉字的数量
func CountChineseWord(s string) (n int){
	b := []byte(s)
	num := len(s)
	for _, c1 := range s{
		if unicode.Is(unicode.Han,c1) {
			n++
		}
	}
	fmt.Printf("num:%v []byte%v\n", num, b)
	return
}

func PrintRangeString(strs []string){
	for _, s := range strs{
		for i, c := range s{
			fmt.Printf("char %v:%v\n", i, c)
		}
	}
}

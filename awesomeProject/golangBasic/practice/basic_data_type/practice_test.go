package basic_data_type

import (
	"fmt"
	"testing"
)

func TestPractice(t *testing.T){
	tests := map[string]interface{}{
		"int":10,
		"float":10.5,
		"bool":true,
	}
	for _, i := range tests{
		printTypeAndValue(i)
	}
}

func TestCountChineseWord(t *testing.T){
	fmt.Printf("count:%v\n", CountChineseWord("hello沙河小王子"))
}

func TestPrintRangeString(t *testing.T) {
	tests := []string{"hello沙河小王子", "hello沙河小王子"}
	PrintRangeString(tests)
}

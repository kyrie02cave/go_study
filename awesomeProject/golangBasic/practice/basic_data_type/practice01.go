package basic_data_type

import "fmt"

func printTypeAndValue(para interface{}){
	// Some types can be done without reflection.
	switch f := para.(type) {
	case bool:
		fmt.Printf("bool\n")
	case float32:
		fmt.Printf("float32\n")
	case float64:
		fmt.Printf("float64\n")
	case int:
		fmt.Printf("int\n")
	case string:
		fmt.Printf("string\n")
	case []byte:
		fmt.Printf("[]byte\n")
	default:
		fmt.Printf("other type:%v\n", f)
	}
	fmt.Printf("type:%T value:%v\n", para, para)
}

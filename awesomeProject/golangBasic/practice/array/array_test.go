package array

import (
	"fmt"
	"testing"
)

func getMax(left int, right int) int{
	if left > right{
		return left
	}else{
		return right
	}
}

func getMin(left int, right int) int{
	if left <= right{
		return left
	}else{
		return right
	}
}

func TestFindTarget(t *testing.T) {
	nums := []int{1, 5, 2, 9, 4}
	leftWant := 1
	rightWant := 4
	target := 9
	left, right := FindTarget(nums, target)
	if getMax(leftWant, rightWant) != getMax(left, right) || getMin(leftWant, rightWant) != getMin(left, right){
		t.Errorf("expected:%#v, %#v got:%#v, %#v", leftWant, rightWant, left, right) // %#v与%#v区别
	}
}

func TestArraySum(t *testing.T) {
	intSlice := []int{1,2,3,4}
	intGot, err := ArraySum(intSlice)
	if err != nil{
		return
	}
	if float64(getIntSum(intSlice)) != intGot{
		fmt.Println("int failed............")
	}else{
		fmt.Println(" int succeed............")
	}
	float32Slice := []float32{1.1, 1.2, 1.3}
	float32Got, err := ArraySum(float32Slice)
	if err != nil{
		return
	}
	v := float64(getFloat32Sum(float32Slice))
	if  v != float32Got{
		fmt.Println("float32 failed............") // 失败。。。。。。。。。浮点型精度问题 bug!!!!!!!
	}else{
		fmt.Println("float32 succeed............")
	}

}

func getIntSum(nums []int) (sum int){
	for _, n := range nums{
		sum += n
	}
	return sum
}

// 重复代码！！！！！！！！！！！
func getFloat32Sum(nums []float32) (sum float32){
	for _, n := range nums{
		sum += n
	}
	return sum
}
package basic_ops

import (
	"testing"
)

func TestFindNumAppearOnce(t *testing.T) {
	type test struct{
		Nums []int
		Want int
	}
	tests := map[string]test{
		"test0":test{Nums:[]int{1,2,2,1,2,2,3}, Want: 3},
		"test1":test{Nums:[]int{1,2,2,3,2,2,3}, Want: 1},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) { // 使用t.Run()执行子测试
			got := FindNumAppearOnce(tc.Nums)
			if got != tc.Want { // slice不可直接比较
				t.Errorf("expected:%#v, got:%#v", tc.Want, got) // %#v与%#v区别
			}
		})
	}
}

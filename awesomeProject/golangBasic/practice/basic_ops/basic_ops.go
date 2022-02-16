package basic_ops

func FindNumAppearOnce(nums []int) (once int){
	for _,n := range nums{
		once ^= n
	}
	return
}

package array

import (
	"errors"
	"reflect"
	"sort"
	"strconv"
)

type Slice struct {
	sort.IntSlice
	idx []int
}
func (s Slice) Swap(i, j int) {
	s.IntSlice.Swap(i, j)
	s.idx[i], s.idx[j] = s.idx[j], s.idx[i]
}

func FindTarget(nums []int, target int) (idx1 int, idx2 int){
	length := len(nums)
	idx := make([]int, length)
	for i := 0; i < length;i++{
		idx[i] = i
	}
	slice := Slice{nums, idx}

	idx1 = -1
	idx2 = -1
	sort.Sort(slice)
	left := 0
	right := length - 1
	for{
		if left >= right{
			break
		}
		if slice.IntSlice[left] + slice.IntSlice[right] == target{
			return slice.idx[left], slice.idx[right]
		}
		if slice.IntSlice[left] + slice.IntSlice[right] > target{
			right--
		}else if slice.IntSlice[left] + slice.IntSlice[right] < target{
			left++
		}
	}
	return
}



// 数组切片求和
func ArraySum(input interface{}) (sum float64, err error) {
	sum = 0
	list := reflect.ValueOf(input)
	switch reflect.TypeOf(input).Kind() {
	case reflect.Slice, reflect.Array:
		for i := 0; i < list.Len(); i++ {
			val := list.Index(i)
			v, err := toFloat64(val.Interface())
			if err != nil {
				return 0, err
			}
			sum += v
		}
	default:
		return 0, errors.New("input must be slice or array")
	}
	return
}

// 统一转化为float64
func toFloat64(in interface{}) (f64 float64, err error) {
	switch val := in.(type) {
	case float64:
		return val, nil
	case float32:
		return float64(val), nil
	case uint8:
		return float64(val), nil
	case uint16:
		return float64(val), nil
	case uint32:
		return float64(val), nil
	case uint64:
		return float64(val), nil
	case uint:
		if strconv.IntSize == 32 || strconv.IntSize == 64 {
			return float64(val), nil
		}
		return 0, errors.New("convert uint to float64 failed")
	case int8:
		return float64(val), nil
	case int16:
		return float64(val), nil
	case int32:
		return float64(val), nil
	case int64:
		return float64(val), nil
	case int:
		if strconv.IntSize == 32 || strconv.IntSize == 64 {
			return float64(val), nil
		}
		return 0, errors.New("convert int to float64 failed")
	case bool:
		if val {
			f64 = 1
		}
		return
	case string:
		f64, err = strconv.ParseFloat(val, 64)
		if err == nil {
			return
		}
		return 0, errors.New("convert string to float64 failed")
	default:
		return 0, errors.New("convert to float64 failed")
	}
}

package calc

import (
	"errors"
	"fmt"
)

func Add(x ...float64) float64{
	sum := 0.0
	for _, n := range x{
		sum += n
	}
	return sum
}

func Sub(subed float64, sub ...float64) float64{
	for _, s := range sub{
		subed -= s
	}
	return subed
}

func Multi(m ...float64) (float64, error){
	length := len(m)
	if length < 2{
		err := errors.New("参数过少。。。")
		fmt.Errorf("param error:%v\n", err)
		return 0, err
	}
	multi := 1.0
	for _, n := range m{
		multi *= n
	}
	return multi, nil
}

func Div(dived float64, div...float64) (float64, error){
	for _, d := range div{
		if d == 0{
			err := errors.New("除数为零。。。")
			fmt.Errorf("param error:%v\n", err)
			return 0, err
		}
		dived /= d
	}
	return dived, nil
}
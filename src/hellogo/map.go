package main

import (
	"fmt"
	"math"
)


type Vertex3 struct {
	Lat, Long float64
}

var m = map[string]Vertex3 {
	"google": {
		40.31331,
		-74.12244,
	},
	"baidu": {
		37.23123,
		-32.9854,
	},
}

func compute(fn func(float64,float64) float64) float64{
	return fn(3,4)
}

func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func fibonacci() func() int{
	step := 0
	before1 := 0
	before2 := 1
	return func() int {
		if(step == 0) {
			step += 1
			return 0
		}
		if(step == 1) {
			step += 1
			return 1
		}

		cur := before2 + before1
		before1 = before2
		before2 = cur
		return cur
	}
}

func main() {
	//m = make(map[string]Vertex3)
	//m["Bell Labs"] = Vertex3{
	//	40.68433, -74.39967,
	//}
	//fmt.Println(m)
	//test1 := map[string]int {}
	//test2 := make(map[string]int)
	//fmt.Println(test1)
	//fmt.Println(test2)

	//m := make(map[string]int,1)
	//
	//m["test"] = 11
	//m["test2"] = 22
	//fmt.Println(m)
	//
	//getvalue := m["test2"]
	//delete(m,"test")
	//
	//fmt.Println(m)
	//fmt.Println(getvalue)
	//
	//elm, ok := m["test"]
	//fmt.Println(elm)
	//fmt.Println(ok)

	//test1["add"] = 32
	//
	//fmt.Println(test1)

	hypot := func(x, y float64) float64{
		return math.Sqrt(x*x + y*y)
	}

	fmt.Println(hypot(5, 12))

	fmt.Println(compute(hypot))

	fmt.Println(compute(math.Pow))

	pos, neg := adder(), adder()

	for i:=0; i<10 ;i++ {
		fmt.Println(
			pos(i),
			neg(-2*i),
		)
	}

	f := fibonacci()
	for i:=0;i<10;i++ {
		fmt.Println(f())
	}




}

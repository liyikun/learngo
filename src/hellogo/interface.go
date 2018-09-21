package main

import (
	"math"
	"fmt"
)

type Myfloat float64

func (f Myfloat) Abs() float64{
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

type Abser interface {
	Abs() float64
}

type I interface {
	hello()
}

type person struct {
	name string
	sex string
}

func (p person) hello() {
	fmt.Println(p.name,p.sex)
}

func main() {
	var a Abser
	f := Myfloat(-math.Sqrt(2))
	//v := verTex{3,4}

	a = f
	fmt.Println(a.Abs())


	var i I = person{"lwm","girl"}
	i.hello()

}


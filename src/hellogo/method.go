package main

import (
	"math"
	"fmt"
)

type verTex struct {
	X, Y float64
}

func (v *verTex) Abs() float64{
	return math.Sqrt(v.X*v.X+v.Y*v.Y)
}


type Myfloat float64

func (f Myfloat) Abs() float64{
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func Scala(v verTex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func Scala2(v verTex,f float64) verTex {
	v.X = v.X * f
	v.Y = v.Y * f
	return v
}

func (v *verTex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func ScaleFunc(v *verTex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main()  {
	//v := verTex{3, 4}
	////Scala(v, 10)
	//v = Scala2(v,10)
	//fmt.Println(v.Abs())
	//
	//t := Myfloat(-32)
	//fmt.Println(t.Abs())
	//
	//
	//var testv verTex
	//
	//testv.Scale(10)
	//
	//p := &testv
	//
	//p.Scale(20)
	//
	//ScaleFunc(p,10)

	v := &verTex{3, 4}
	fmt.Printf("Before scaling: %+v, Abs: %v\n", v, v.Abs())
	v.Scale(5)
	fmt.Printf("After scaling: %+v, Abs: %v\n", v, v.Abs())


}

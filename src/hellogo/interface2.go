package main

import (
	"fmt"

)

type I interface {
	M()
}

type T struct {
	S string
}

func (t *T) M() {
	if t == nil {
		fmt.Println("args is nil")
		return
	}
	fmt.Println(t.S)
}

type F float64

func (f F) M() {
	fmt.Println(f)
}

func do(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("Twice %v is %v\n", v, v*2)
	case string:
		fmt.Printf("%q is %v bytes long\n", v, len(v))
	default:
		fmt.Printf("I don't know about type %T!\n", v)
	}
}

type person struct {
	Name string
	Age int
}

func (p person) String() string {
	return fmt.Sprintf("%v (%v years)",p.Name,p.Age)
}

func main()  {

	a := person{"any", 22}
	z := person{"lyk",31}

	fmt.Println(a,z)

	//do(32)
	//do("hha")
	//do(true)
	//
	//var i I
	//
	//var t *T
	//i = t
	//descibe(i)
	//i.M()
	//
	//i = &T{"Hello"}
	//descibe(i)
	//i.M()
	//
	////i = F(math.Pi)
	////descibe(i)
	////i.M()
	//
	//
	////descibe(i)
	////i.M()
	//
	//var in interface{} = "hello"
	//describe2(in)
	//
	//s := in.(string)
	//fmt.Println(s)
	//
	//s, ok := in.(string)
	//fmt.Println(s, ok)
	//
	//f, ok := in.(float64)
	//fmt.Println(f, ok)
	//
	////f = in.(float64)
	////fmt.Println(f)



}

func describe2(i interface{})  {
	fmt.Printf("(%v,%T)\n",i ,i)
}

func descibe(i I)  {
	fmt.Printf("(%v，%T)\n",i ,i)
}
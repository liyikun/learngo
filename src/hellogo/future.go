package main

import (
	"fmt"
)

import (
	"golang.org/x/tour/pic"
	"go/ast"
)


type Vertex struct {
	X int
	Y int
}


type Vertex2 struct {
	x,y int
}

var (
	v1 = Vertex2{1,2}
	v2 = Vertex2{x: 1}
	v3 = Vertex{}
	p = &Vertex{1,2}
)

func printSlice(s []int)  {
	fmt.Printf("len=%d cap=%d %v\n",len(s),cap(s),s)
}

func Pic(dx, dy int) [][]uint8 {
	pacy := make([][]uint8,dy)
	for i := range pacy {
		pacy[i] = make([]uint8,dx)
	}
	return pacy
}


//var pow = []int{1,2,4,8,16,32,64,128}

func main() {
	//fmt.Println("haha")
	//
	//var p *int
	//i := 42
	//p = &i
	//fmt.Println(*p)
	//*p = 21

	//i, j := 42, 2701

	//p := &i
	//fmt.Println(*p)
	//*p = 21
	//fmt.Println(i)
	//
	//p = &j
	//*p = *p /37
	//fmt.Println(j)
	//
	//fmt.Println(Vertex{1,2})

	//v:=Vertex{1,2}
	//v.X = 4
	//fmt.Println(v)
	//
	//p := &v
	//p.X = 1e9
	//fmt.Println(v)
	//fmt.Println(v1,v2,v3,*p)
	//
	//var a [2]string
	//a[0] = "Hello"
	//a[1] = "World"
	//fmt.Println(a[0],a[1])
	//fmt.Println(a)
	//
	//primes := [6]int{2,3,4,5,6,7}
	//fmt.Println(primes)
	//s := primes[1:4]
	//fmt.Println(s)

	// q := []int{2,3,4,535,52,32,12}
	// fmt.Println(q)

	// s := []struct {
	// 	i int
	// 	b bool
	// }{
	// 	{1,true},
	// 	{2,false},
	// 	{3,true},
	// }

	// fmt.Println(s)

	// q = q[2:6]
	// //fmt.Println(q)
	// printSlice(q)

	// q = q[:3]
	// //fmt.Println(q)
	// printSlice(q)

	// q= q[2:]

	// //fmt.Println(q)
	// printSlice(q)

	// var tests [] int
	// fmt.Println(tests, len(tests), cap(tests))
	// if tests == nil {
	// 	fmt.Println("nil")
	// }

	// a := make([]int, 5)
	// printSlice(a)
	// b := make([]int, 0, 5)
	// printSlice(b)

	// b = b[:cap(b)]
	// printSlice(b)

	// board := [][]string {
	// 	[]string{"_","_","_"},
	// 	[]string{"_","_","_"},
	// 	[]string{"_","_","_"},
	// }
	
	// board[0][0] = "X"
	// board[2][2] = "O"
	// board[1][2] = "X"
	// board[1][0] = "0"
	// board[0][2] = "X"

	// for i :=0; i < len(board); i++ {
	// 	fmt.Printf("%s\n", strings.Join(board[i], ""))
	// }


		//var s []int
		//s = append(s, 0)
		//printSlice(s)
		//
		//s =append(s, 1)
		//printSlice(s)
		//
		//s = append(s,1,2,3,4)
		//printSlice(s)
		//
		//for i, v:= range pow {
		//	fmt.Printf("2**%d = %d\n",i,v)
		//}
		//
		//


		pow := make([]int, 10)
		for i := range pow {
			pow[i] = 1 << uint(i) // == 2**i
		}
		for _,value := range pow {
			fmt.Printf("%d\n", value)
		}

		fmt.Println(2 << 10)

	pic.Show(Pic)
	}

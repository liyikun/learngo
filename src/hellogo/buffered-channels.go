package main

import "fmt"

//func fibonacci(n int,c chan int)  {
//	x, y := 0, 1
//	for i :=0 ; i < n ; i++ {
//		c <-x
//		fmt.Println("channel",x)
//		x, y = y, x+y
//	}
//	close(c)
//}
//
//func getFib(c chan int)  {
//	for i := range c {
//		fmt.Println("main",i)
//	}
//}

func fibonacci(c, quit chan int)  {
	x, y := 0, 1
	for {
		select {
		case c <-x:
			x, y =y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

func main()  {
	c := make(chan int)
	quit := make(chan int)
	go func() {
		for i :=0;i < 10; i++ {
			fmt.Println(<-c)
		}
		quit <-0
	}()
	fibonacci(c, quit)
	//c := make(chan int, 10)
	//go fibonacci(cap(c), c)
	//go getFib(c)
	//fmt.Println("hello")

	//ch := make(chan int, 2)
	//ch <- 1
	//ch <- 2
	//ch <- 3
	//fmt.Println(<-ch)
	//fmt.Println(<-ch)
	//fmt.Println(<-ch)
}

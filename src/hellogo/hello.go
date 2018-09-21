package main

import (
	"fmt"
	"math"
	"runtime"
	"time"
)

func add(x, y int) int {
	return x + y
}

func swap(x, y string) (string, string) {
	return y,x
}

func split(sum int) (x, y int) {
	// x = sum * 4 / 9
	// y = sum - x
	return
}

//var c, python, java bool


//var i, j int = 1, 2

func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

func pow(x, n, lim float64) float64 {
	if v:= math.Pow(x, n); v < lim {
		return v
	}
	return lim
}

func main() {
	//
	//var i, j int = 1,2
	//k := 3
	//c, python, java := true, false, "no！"
	//
	//
	//fmt.Println("hello world!")
	//fmt.Println(add(42, 13))
	//a,b := swap("hello","world!")
	//fmt.Println(a, b)
	//fmt.Println(split(17))
	//fmt.Println(i, c, python, java)
	//fmt.Println(i, j, c, java, k)

	var i int = 42
	var f float64 = float64(i)
	var u uint = uint(f)
	fmt.Println(u)


	//sum := 0
	//for i :=0;i < 10; i ++ {
	//	sum += i
	//}
	//fmt.Println(sum)

	sum := 1
	for ;sum < 1000; {
		sum += sum
	}
	fmt.Println(sum)
	fmt.Println(sqrt(2),sqrt(-4))
	fmt.Println(pow(2,3,10))
	fmt.Println(pow(2,4,10))

	fmt.Println("go is run")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux")
	default:
		fmt.Printf("%s",os)
	}

	fmt.Println("When's Saturday?")
	today := time.Now().Weekday()
	switch time.Saturday {
	case today + 0:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tomorrow.")
	case today + 2:
		fmt.Println("In two days.")
	default:
		fmt.Println("Too far away.")
	}

	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon!")
	default:
		fmt.Println("Good evening.")
	}

	defer fmt.Println(("world!"))
	fmt.Println(("hello"))

	fmt.Println("counting")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("done")

}



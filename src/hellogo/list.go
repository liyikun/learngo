package main

import (
	"fmt"
	"encoding/json"
)

func maxSequence(list[]int) int {
	suma := list[0]
	sumb := 0
	max := list[0]
	for _, value:= range list{
		suma +=	value
		sumb += value

		if(suma > max) {
			max = suma
		}
		if(sumb > max) {
			max = sumb
		}
		if(suma<=sumb) {
			suma = 0
		} else {
			sumb = 0
		}
 	}
	return max
}

func main() {
	var list string
	var outlist []int
	fmt.Scan(&list)
	json.Unmarshal([]byte(list), &outlist)
	fmt.Println(maxSequence(outlist))
}
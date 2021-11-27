package random

import (
	"fmt"
)

func RandomLoopMain() {
	randomLoop1()
	randomLoop2()
	randomArray()
	randomSliceAppended()
	randomSliceLiteral()
}

func randomLoop1() {
	fmt.Println("### randomLoop1")
	for i := 0; i < 3; i++ {
		fmt.Println(i)
	}
	fmt.Println("")
}

func randomLoop2() {
	fmt.Println("### randomLoop2")
	i := 0
	isLessThanFive := true
	for isLessThanFive {
		if i >= 3 {
			isLessThanFive = false
		}
		fmt.Println(i)
		i++
	}
	fmt.Println("")
}

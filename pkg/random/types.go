package random

import (
	"fmt"
)

func randomArray() {
	fmt.Println("### randomArray")
	// NOTE: arrays are fixed length
	var langs [3]string
	langs[0] = "Go"
	langs[1] = "Python"
	langs[2] = "Ruby"
	fmt.Println(langs)
	fmt.Println(langs[0])
	fmt.Println("")
}

func randomSliceAppended() {
	fmt.Println("### randomSliceAppended")
	var langs []string
	langs = append(langs, "Go")
	langs = append(langs, "Python")
	langs = append(langs, "Ruby")
	langs = append(langs, "Java")
	fmt.Println(langs)
	fmt.Println(langs[0])
	fmt.Println("")
}

func randomSliceLiteral() {
	fmt.Println("### randomSliceLiteral")
	langs := []string{"Go", "Python", "Ruby"}
	fmt.Println(langs)
	fmt.Println(langs[0])
	fmt.Println("")
}

package main

import (
	"bytes"
	"fmt"
)

func main() {

	str0 := "Welcome Home"
	str1 := "Ganpati Bappa"

	fmt.Println(str0 + str1)

	//using bytes
	var b bytes.Buffer

	b.WriteString("T")
	b.WriteString("H")
	b.WriteString("I")
	b.WriteString("S")
	b.WriteString("Y")
	b.WriteString("E")
	b.WriteString("A")
	b.WriteString("R")
	fmt.Println(b.String())

}

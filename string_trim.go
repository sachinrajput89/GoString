package main

import (
	"fmt"
	"strings"
)

func main() {

	str0 := "!! Welcome Home !!"

	str1 := strings.Trim(str0, "!")

	fmt.Println(str1)

}

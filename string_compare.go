package main

import "fmt"

func main() {

	str0 := "Welcome boss"

	//	str1 := 'Welcome Boss'

	str2 := `Welcome Boss`

	//	str3 := `Welcome Boss'

	fmt.Println(str0, str2)

	//==========================================================

	//iterate using range

	for index, s := range "main value" {
		fmt.Println(index, s)
	}

}

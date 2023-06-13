//assignment2
//author - meet mehta

package main

import (
	"fmt"
	"strings" //since we are using strings inbuilt function we will need to import its library
)

func main() {
	var str string
	fmt.Print("enter the string to covert into UPPERCASE:\n")
	fmt.Scan(&str)
	str_upper := strings.ToUpper(str)

	//The syntax of ToUpper() function is
	//strings.ToUpper(str)

	fmt.Printf(str_upper)
}

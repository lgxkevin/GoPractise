package main

import (
	"fmt"
	"strconv"
)

var k int = 99

//var(
//	firstName string = "Kevin"
//	lastName string = "Kevin"
//	age int = 3
//)
func main() {
	//fmt.Println("Message: ", testMessage.TextMessage)

	//var i int = 42
	var h int = 30
	var w string
	w = strconv.Itoa(h)
	fmt.Printf("%v, %T\n", w, w)
	fmt.Printf("%v, %T\n", k, k)
}

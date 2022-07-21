// You can edit this code!
// Click here and start typing.
package main

import "fmt"

func test() {
	mySlice := []string{"Hey", "There", "HO", "are", "you"}
	updateSlice(mySlice)
	fmt.Println(mySlice)
}

func updateSlice(s []string) {
	s[0] = "Bye"
}

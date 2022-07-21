package main

import "fmt"

type contactInfo struct{
	email string
	zipCode int
}

type person struct{
	firstName string
	lastName string
	contact contactInfo
}

func main(){
// alex := person{"Aviral","Garg" }

// fmt.Println(alex)
// fmt.Printf("%+v",alex)

jim :=person{
	firstName:"Aviral",
	lastName:"Garg",
	contact:contactInfo{
		email:"hey@gmail.com",
		zipCode:12345,
	},
}
// jimPointer := &jim
// jimPointer.updateName("heyyy")
//we can write this directly as a shortcut
jim.updateName("Hey")
jim.print()
}
func (pointerToPerson *person) updateName(newFirstName string){
(*pointerToPerson).firstName=newFirstName
}
func (p person) print(){
	fmt.Printf("%+v",p)

}
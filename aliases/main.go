package main

import "fmt"

//all alisases have receiver functions

type Code struct {
	time float64
	name string
}

type str_my_type string

//constructor function for struct type
func newCode(time float64, name string) Code {
	return Code{time, name}
}

//receiver function for type Code

func (c Code) OutputValues() {
	fmt.Println("NAME IS :", c.name, "and time taken is :", c.time)
}

//receiver func for type my_type_str
func (s str_my_type) display() {
	fmt.Println(" my NAME IS :", s)
}

func main() {
	codeObj := newCode(32.5, "one plus")
	codeObj.OutputValues()

	var my_name str_my_type = "Vibha"
	my_name.display()

}

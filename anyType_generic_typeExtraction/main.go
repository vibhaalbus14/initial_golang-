package main

import "fmt"

//pass in any type by creating an empty interface
//interface{}

//generic syntax
//func func_name[Variable_rep allowed_type1 | allowed_type2  |.. | allowed_typen] (arg variable_rep) (return_type){}

//extract value types
//1.using switch =>
// switch val.(type){
// case int:...
// case float64:...
// default:..,}

//2.plane extraction
//val.(int) => will check if int
//val.(float64)=> will check if val is of type float64

func Display(n interface{}) {
	fmt.Println("Val is , ", n)
}

//restrict to certain types
func restriction[T int | string | float64](val T) {
	fmt.Println("Heyy ! Im a generic , and i allow only certain types of inputs to be processed")
	fmt.Println("Value is ,", val)
}
func main() {

	Display(5)
	Display(2.3)
	Display("Im a  string")
	Display('r')

	//calling generic function
	restriction(5)
	restriction(2.3)
	restriction("Im a  string")
	//restriction('r')

}

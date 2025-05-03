package main

import (
	"fmt"
	"struct_example/basic_interface"
	"struct_example/note"
	"struct_example/todo"
)

func main() {

	// //calling for note
	// noteObj, err := note.Input()
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	// noteObj.Display()
	// err1 := noteObj.Save()
	// if err1 != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	// //calling for todo
	// todoObj, err := todo.Input()
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	// todoObj.Display()
	// err3 := todoObj.Save()
	// if err3 != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	//calling the same save and display using interface
	//interface has only the function signatures in it
	//anything that has these functions in it is automatically a type of this interface

	//interface receiver function

	noteObj, err := note.Input()
	if err != nil {
		fmt.Println(err)
		return
	}

	todoObj, err1 := todo.Input()
	if err1 != nil {
		fmt.Println(err1)
		return
	}

	//use interface function to save
	//since todo and NOte objects implement both display and save functions , they automatically are caller type
	basic_interface.CallInterface(noteObj)
	basic_interface.CallInterface(todoObj)
}

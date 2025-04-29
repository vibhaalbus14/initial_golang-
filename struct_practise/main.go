package main

import (
	"fmt"
	"struct_example/note"
)

func main() {

	noteObj, err := note.Input()
	if err != nil {
		fmt.Println(err)
		return
	}

	noteObj.Display()
	err1 := noteObj.Save()
	if err1 != nil {
		fmt.Println(err)
		return
	}
}

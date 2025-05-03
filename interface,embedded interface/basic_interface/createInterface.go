package basic_interface

import (
	"fmt"
)

type Caller interface {
	Display()
	Save() error
}

func CallInterface(c Caller) {
	c.Display()
	err := c.Save()
	if err != nil {
		fmt.Println(err)
		return
	}
}

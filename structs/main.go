package main

import (
	"fmt"
	"struct_creation/user"
)

func main() {
	//user.Checkuser()
	u1 := user.NewUser("Vibha", 23)
	u1.Printvalues()
	u1.Overwrite_values("vishu", 24)
	u1.Printvalues()
	fmt.Println("-------------------------------")
	u2 := user.NewStudent("Maths", "hARI", 5, 21)
	u2.PrintStuValues()

}

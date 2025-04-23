package user

import (
	"fmt"
)

// custom datatype called "User " is craeted and struct defines its structure
type User struct {
	name string // fields cannot be exported since the begin with small letters => encapsulation is achieved
	age  int
}

type Student struct {
	subject  string
	semester int
	*User
}

func NewStudent(userSub, userName string, userSem, userAge int) *Student {
	return &Student{userSub, userSem, NewUser(userName, userAge)}

}

func Checkuser() {
	//instantiation
	u1 := User{name: "Vibha", age: 22} //iitializes all fields to its type def values

	fmt.Println(u1.name, u1.age)
}

//normal function
//its a constructor used to initialse instance of struct
//reduces repeated code to create multiple instances

func NewUser(userName string, userAge int) *User {
	return &User{userName, userAge} // return pointer pointing to new instance
	//scope of the pointer is amde available throughout by go
	//it pushes pointers or the instance to heap instead of stack
}

func (u *User) Printvalues() {
	fmt.Println((*u).name, (*u).age)
}

// receiver function
func (u *User) Overwrite_values(newName string, newAge int) {
	(*u).name = newName
	(*u).age = newAge
}

func (s *Student) PrintStuValues() {
	fmt.Println(s.semester, s.subject, s.name, s.age)
}

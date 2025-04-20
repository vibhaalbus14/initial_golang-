//null pointer=> not initialised and has a default val <nil>
//dangling pointer=> pointer to a memory space that is freed/ pointer to invalid memory

//merits of pointers
//1.due to direct manipulation, less code to write to make changes to variables
//2.less memory space by avoiding making variable copies

package main

import "fmt"

func main() {
	x := 10
	var p *int
	fmt.Println("value of x is before calling passByVal is ,", x)
	fmt.Println("value of pointer before initialisation => null pointer is", p)
	fmt.Println("-----------------------------------------------------------------")

	p = &x //initialise p with address of x
	fmt.Println("p gives me address", p)
	fmt.Println("*p defereneces the value stored at address => dereferencing operation", *p)
	fmt.Println("-----------------------------------------------------------------")

	passByVal(x) //call to function by passing valueconst
	// here, go makes a copy of x in a diff address with same value and passes this copy to func
	//so changes made to this in function scope is not refelected in original variable
	//as function works with a  diff copy
	fmt.Println("after calling passByValue, x is changed to ", x)
	passByRef(p) // call to function by passing the address of original variable => passByRef
	//function works with the same original variable without making a copy
	//thus less memory space and less code
	//changes made are reflected in original variable
	fmt.Println("after calling passByRef, x is changed to", x)

}

func passByVal(a int) {
	a -= 8
}

func passByRef(a *int) {
	*a -= 8
}

package main

import "fmt"

func double(ptr *[]int) {
	for index, val := range *ptr {
		(*ptr)[index] = val * 2
	}
}

func triple(ptr *[]int) {
	for index, num := range *ptr {
		(*ptr)[index] = num * 3
	}
}

func variadic(nums ...int) { //expects "n" number of individual numbers, but does not know the count
	//here all integers are collected and go internally compiles them into a slice called num of type int
	fmt.Println("the values in slice created are:")
	for _, val := range nums {
		fmt.Println(val)
	}
}

//passing functions as values
func transform(num *[]int, transformFunc func(*[]int)) {
	transformFunc(num)
}

//return function as values
func chooseFunc(num int) func(*[]int) { // return function itself
	if num == 2 {
		return double
	} else {
		return triple
	}
}

func main() {
	//initial slices
	firstSlice := []int{1, 2, 3, 4, 5}
	secondSlice := []int{6, 7, 8, 9, 10}
	thirdSlice := []int{11, 12, 13, 14, 15}

	fmt.Println("initially: ", firstSlice)
	fmt.Println(secondSlice)

	//which function to choose=> return a function
	trfFunc1 := chooseFunc(2)
	trfFunc2 := chooseFunc(3)

	//calling a facade that implements given function on given slice
	transform(&firstSlice, trfFunc1)
	transform(&secondSlice, trfFunc2)
	fmt.Println("after calling: ", firstSlice)
	fmt.Println(secondSlice)

	//implementing spread and group operator
	variadic(thirdSlice...) //spreads thrirdSlice and passes individual numbers

}

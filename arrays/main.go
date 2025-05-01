package main

import "fmt"

//creating ana array
//syntax: arrayName:= [size]data_type{vals}

//make
//it will initally allocate specified mem space and saves frequent memory reallocation
//syntax: slice_name=make([]type , length,cap)
//in the above syntax:
//1.a dynamic array / slice of given type must be specified
//2.length => initial empty slots for which default value of given type is initialised
//3.cap => total cap of array behind the scenes
//eg: arr:= make([]int, 2,5)
//array generated for the slice with name arr is [0,0]
//this means out of 5= capacity, 2 will be filled with null values and thus length=2
func main() {
	arr1 := [3]int{5, 6, 7}
	fmt.Println("arr1", arr1)
	fmt.Println("length", len(arr1))
	fmt.Println("capacity", cap(arr1))

	fmt.Println("creating an array with partial initialization")
	arr2 := [3]int{5} //partial initialisation will fill in the empty
	// slots with def value of the data type, in this case 0 because data type is int
	fmt.Println("arr2", arr2)
	fmt.Println("length", len(arr2))
	fmt.Println("capacity", cap(arr2))

	fmt.Println("appending to partial initialised array")
	fmt.Println("slicing the entire array")
	//arr2:=append(arr2) => append to fixed array cannot be done directly
	//1.use indices to initialize
	//2.use a slice to append values to array
	s := arr2[:] //when entire array is sliced, append will increase the capacity of slice
	//thus, it is no more backed by arr2 under the hood , but by diff array
	fmt.Println("before appending => arr2,slice", arr2, s)
	s = append(s, 7)
	fmt.Println("after appending => arr2,slice", arr2, s)

	fmt.Println("slicing the array partially")
	//arr2:=append(arr2) => append to fixed array cannot be done directly
	//1.use indices to initialize
	//2.use a slice to append values to array
	s1 := arr2[:1] //when entire array is sliced, append will increase the capacity of slice
	//thus, it is no more backed by arr2 under the hood , but by diff array
	fmt.Println("before appending => arr2,slice", arr2, s1)
	s1 = append(s1, 7)
	fmt.Println("after appending => arr2,slice", arr2, s1)

	//adding elements by accessing
	arr2[0] = 9
	fmt.Println("after changing by accessing index => arr2,slice", arr2, s)

	//unpcaking list values=> extend list with another list
	//names := []string{"vibha", "vijaya", "vinutha"}
	names := make([]string, 1, 5)
	names[0] = "vibha"
	names = append(names, "vijaya", "vinutha")
	newNames := []string{"vishu", "vasudevan", "vamika"}
	fmt.Println("initial names arr : ", names)
	fmt.Println("newNames arr", newNames)

	//extend names to include newNames
	names = append(names, newNames...) //... => spread operator, spreds out the elments in slice
	fmt.Println("after extending")
	fmt.Println("names arr : ", names)
	fmt.Println("newNames arr", newNames)
	fmt.Println()

	//for loop
	fmt.Println("Printing the values of names using for loop")
	for i := 0; i < len(names); i++ {
		fmt.Println(names[i])
	}
	fmt.Println()
	fmt.Println("Printing the values of names using for-range loop")
	for _, val := range names {
		fmt.Println(val)
	}

}

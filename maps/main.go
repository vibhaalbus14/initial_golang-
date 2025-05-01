package main

import "fmt"

//map syntax
//variable_name := map[key data_type]values_data_type{key1:val1,key2:val2...}
//append new values=> variable_name[key_name]=val

//make
//it will initally allocate spcifies mem space and saves frequeen memory reallocation
//syntax:; hashMap:=make(map[key data_type]values_data_type,initial_length)

func main() {

	//hashMap := map[int]string{1: "Vibha", 2: "Vishu"}
	hashMap := make(map[int]string, 3)
	fmt.Print("hahsmAp, ", hashMap)

	hashMap[3] = "Bhuvana"
	fmt.Print("hahsmAp after adding new val, ", hashMap)

	fmt.Println("accessing through key 2", hashMap[2])

	fmt.Println("printing values in hahsMap using for-range loop")
	for key, val := range hashMap {
		fmt.Println(key, val, hashMap[key])
	}
}

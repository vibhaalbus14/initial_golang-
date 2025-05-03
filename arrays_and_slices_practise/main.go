package main

import "fmt"

type Product struct {
	title string
	id    int
	price float64
}

//constructor to product
func newProduct(currId int, currTitle string, currPrice float64) Product {
	return Product{id: currId, price: currPrice, title: currTitle}
}

func main() {

	//1.create an array of 3 hobbies
	hobbies := [3]string{"physical excercise", "code", "eat"}
	fmt.Println("hobbies array : ", hobbies)

	//2
	fmt.Println("firstElement only, by accessing via index  : ", hobbies[0])
	s1 := hobbies[1:]
	fmt.Println("last 2 elements only, by creating new slice s1  : ", s1)

	//3
	s2 := hobbies[:2] //one way by creating direct slice
	s3 := hobbies[:1]
	s3 = append(s3, hobbies[1]) //extending the prev slice within its capacity
	fmt.Println(`slice of first 2 elements `, s2, s3)

	//4
	s4 := s2[1:]
	s4 = append(s4, hobbies[2])
	fmt.Println(`slice of last 2 elements by reslicing the prev slice`, s4)

	//5
	goals := []string{"learn go lang basics", "build backend "}
	fmt.Println("Goals dynamic array/slice is :", goals)

	//6
	goals[1] = "rest backend"
	goals = append(goals, "learn additional functionality")
	fmt.Println("Goals after changing 2 nd ele and appending 3rd ele :", goals)

	//7

	//create a slice of type product
	productSlice := []Product{newProduct(1, "Pen", 5.5), newProduct(2, "Pencil", 6.7)}
	fmt.Println("dynamic array of products : ", productSlice)
	productSlice = append(productSlice, newProduct(3, "Eraser", 2))
	fmt.Println("dynamic array of products after appending 3rd product : ", productSlice)
	fmt.Println("acessing 3 rd product title from productSlice ", productSlice[2].title)

}

package main

/*
1.data types
2.variable declaration
3.types of print and scan statements
4.sprint stmt
5.func def, returning multiple values
6.scanf reads only the formatted chars and leaves new line in buffer, this gets picked up by next scan and assumes this to
be the inp=> soln make the first scanf-< scanln, this reads till new line and discards new line and does not affect next
scanf
7.main function must have the package main
*/
import (
	"fmt"
	"pranave/gugu"
)

func main() {
	//declare variables
	//earnings before tax(ebt), earnings after tax(eat)
	var revenue, expense, taxRate, ebt, eat, ratio float32

	fmt.Print("enter revenue,expense and tax rate")
	//fmt.Scanf("%f %f %f", &revenue, &expense, &taxRate)
	fmt.Scanln(&revenue, &expense, &taxRate)

	ebt = revenue - expense
	eat = revenue - (revenue * (taxRate / 100)) - expense
	ratio = ebt / eat

	fmt.Printf("earnings before tax is : %.2f \nearnings after tax is : %.2f \nratio is : %.2f ", eat, ebt, ratio)
	gugu.Initialize()
	fmt.Println(`i'm a multiline string.
	i 
	can
	span
	across
	various
	lines.
	YOU cant use \n in me, i wont give u next line :)`)
	fmt.Println("-----------------------------------------------------------------")
	var l int
	var b int
	fmt.Print("here lets call a function to calc area, now gimme the length and breadth values")
	fmt.Scanf("%d %d", &l, &b)
	rect, sq := calcArea(l, b)
	fmt.Printf("area of rect is %d and sq is %d ", rect, sq)
	fmt.Println("there we go ;)")

}

func calcArea(l int, b int) (int, int) {
	fmt.Println("area is : ", int(l*b))
	return int(l * b), l * l
}

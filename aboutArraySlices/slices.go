package main

import (
	"fmt"
)

var gGroceries[]string

func addGrocery(items ...string) {
	fmt.Println("Capacity: ", cap(gGroceries))
	for _,item:=range items{
		gGroceries=append(gGroceries, item)
	}
}

func listGroceries() {
	fmt.Println("Grocery list is as follows:")
	// for i := 0; i < len(gGroceries); i++ {
	// 	fmt.Println(gGroceries[i])
	// }

	// advance version
	for index,item:=range gGroceries{
		fmt.Println(index,item)
	}
}

func main() {
	listGroceries()
	addGrocery("Apple")
	listGroceries()
	addGrocery("Banana")
	addGrocery("Banana2")
	addGrocery("Banana3")
	addGrocery("Banana4")
	addGrocery("Banana5")
	addGrocery("Banana6")
	addGrocery("Banana", "Pineapple", "Meat")
	listGroceries()
}

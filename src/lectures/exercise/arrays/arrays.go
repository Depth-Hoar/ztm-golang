//--Summary:
//  Create a program that can store a shopping list and print
//  out information about the list.
//
//--Requirements:
//* Using an array, create a shopping list with enough room
//  for 4 products
//  - Products must include the price and the name
//* Insert 3 products into the array
//* Print to the terminal:
//  - The last item on the list
//  - The total number of items
//  - The total cost of the items
//* Add a fourth product to the list and print out the
//  information again

package main

import "fmt"

type Product struct {
	price int
	name  string
}

//   - Products must include the price and the name
//
// * Insert 3 products into the array
// * Print to the terminal:
//   - The last item on the list
//   - The total number of items
//   - The total cost of the items
func printStats(list [4]Product) {
	var cost, totalItems int

	for i := 0; i < len(list); i++ {
		item := list[i]
		cost += item.price

		if item.name != "" {
			totalItems += 1
		}
	}
	fmt.Println("last item on the list:", list[totalItems-1])
	fmt.Println("total items", totalItems)
	fmt.Println("Total cost", cost)
}

func main() {
	//   - Using an array, create a shopping list with enough room
	//     for 4 products
	shoppingList := [4]Product{
		{1, "Banana"},
		{6, "meat"},
		{3, "salad"},
	}

	printStats(shoppingList)

	//* Add a fourth product to the list and print out the
	//  information again
	shoppingList[3] = Product{4, "Bread"}

	printStats(shoppingList)
}

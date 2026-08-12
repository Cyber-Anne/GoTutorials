package main

import "fmt"

func main() {
	fruits := []string{"Apple", "Grapes", "Banana", "Strawberry", "Mango", "Watermelon"}
	fmt.Println("My favourite fruits are: ")

	for i := 0; i < len(fruits); i++ {
		fmt.Println(fruits[i])
	}

	shoppingList := []string{"Milk", "Bread", "Sausage", "Eggs", "Spices", "Flour"}
	for index, value := range shoppingList {
		fmt.Printf("%d. %v\n", index+1, value)
	}

}

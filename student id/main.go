package main

import "fmt"

func main() {
	var name string = "Marie-Anne"
	var age int = 20
	University := "University of Ghana"
	var Course string = "Computer Science"
	var Level int = 300

	fmt.Println("============================")
	fmt.Println("         STUDENT ID         ")
	fmt.Println("============================")
	fmt.Println("Name: ",name)
	fmt.Printf("Age: %v \n", age)
	fmt.Printf("University: %v \n", University)
	fmt.Printf("Course: %q \n", Course)
	fmt.Printf("Level: %v \n", Level)
	fmt.Println("============================")


}


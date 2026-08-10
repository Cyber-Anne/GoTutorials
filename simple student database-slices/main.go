package main

import "fmt"

func main() {
	names := []string{"Marie", "Anne", "Dzifa", "Afia", "Naa"}
	ages := []int{18, 23, 22, 21, 26, 19, 32}
	courses := []string{"Archaeology", "Engineering", "Medicine", "Architecture", "Cosmetology"}

	fmt.Println("STUDENT 1")
	fmt.Println("Name: ", names[0])
	fmt.Println("Age: ", ages[1])
	fmt.Println("Course: ", courses[1])
	fmt.Println()
	fmt.Println("STUDENT 2")
	fmt.Println("Name: ", names[1])
	fmt.Println("Age: ", ages[3])
	fmt.Println("Course: ", courses[2])
	fmt.Println()
	fmt.Println("STUDENT 3")
	fmt.Println("Name: ", names[3])
	fmt.Println("Age: ", ages[5])
	fmt.Println("Course: ", courses[4])

}

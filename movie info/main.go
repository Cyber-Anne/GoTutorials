package main

import "fmt"

func main(){
	var movie string = "Blended"
	year:= 2020
	var rating = 9.5
	director := "Christopher Nolan"

	fmt.Printf("%v was released in %v \n", movie, year)
	fmt.Printf("It has a rating of %.3f \n", rating)
	fmt.Println("Directed by ",director)

}

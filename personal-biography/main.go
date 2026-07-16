package main

import "fmt"

func main(){
	name:= "Marie-Anne"
	age:= 20
	country:="Ghana"
	occupation:="Student"
	hobby:="reading"
	dreamJob:="engineer"

	fmt.Print("HELLO!")
	fmt.Print("\n")
	fmt.Print("\n")
	fmt.Println("My name is ", name)
	fmt.Printf("I am %v years old.\n", age)
	fmt.Println("I live in ", country)
	fmt.Println("My hobby is ", hobby)
	fmt.Printf("I am currently a %v but I want to be an %v.", occupation,dreamJob)

}
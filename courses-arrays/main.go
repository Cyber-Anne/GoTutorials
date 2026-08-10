package main
import "fmt"

func main(){
	name := "Marie-Anne"
	age:= 20
	university := "University of Ghana"
	programme := "Computer Science"

	courses := [7]string{"DCIT 304","DCIT 306","DCIT 308","DCIT 310","DCIT 312","DCIT 318","DCIT 322"}

	fmt.Println("Student: ",name)
	fmt.Println("Age: ",age)
	fmt.Println("University: ",university)
	fmt.Println("Programme: ",programme)
	fmt.Println()
	fmt.Println()
	fmt.Println("Courses Registered")
	fmt.Println(courses[0])
	fmt.Println(courses[1])
	fmt.Println(courses[2])
	fmt.Println(courses[3])
	fmt.Println(courses[4])
	fmt.Println(courses[5])
	fmt.Println(courses[6])
}
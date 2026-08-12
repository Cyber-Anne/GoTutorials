package main
import "fmt"

func main(){
	i:=1
	
	for i<= 12 {
		y:=5*i
		fmt.Printf("5 x %v = %v\n",i,y)
		i++
	}
	for i=2; i<=20; i+=2 {
		fmt.Println(i)
	}
	for i=5; i<=50; i+=5 {
		fmt.Println(i)
	}
}
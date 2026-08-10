package main
import "fmt"

func main(){
foods := []string{"indomie","waakye","kenkey","banku","rice"}

fmt.Println("All foods: ",foods)
fmt.Println()
fmt.Println("First food: ",foods[0])
fmt.Println("Third food: ",foods[2])
fmt.Println()
fmt.Println()

foods = append(foods, "fries","plantain")
fmt.Println("New list of foods: ",foods)
}
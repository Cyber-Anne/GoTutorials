package main
import "fmt"

func main(){
	movies := []string{"Blended","After","Reality High","He's all that","Venom","Office romance","The notebook","Kissing booth"}
    rangeOne:= movies[ :3]
	rangeTwo:= movies[3:6]
	rangeThree:= movies[2:7]

	fmt.Println("These are the movies i have watched recently")
	fmt.Println(movies)
	fmt.Println()
	fmt.Println("The most romantic among them were ",rangeOne)
	fmt.Println()
	fmt.Println("The most captivating among them were ",rangeTwo)
	fmt.Println()
	fmt.Println("The most tearful among them were ",rangeThree)
	


}
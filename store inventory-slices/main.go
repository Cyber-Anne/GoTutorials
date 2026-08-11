package main
import "fmt"

func main(){
	products:=[]string{"Glycolic acid", "Vitamin C serum", "Niacinamide serum", "Aloe vera gel", "Moisturizer", "Sunscreen"}
	prices:=[]float32{20.54, 60, 120,240,80,220,32.46,150,230}

	fmt.Println("============Store Inventory================")
	fmt.Println("Product: ",products[3])
	fmt.Println("Price: ",prices[1])
    fmt.Println()
	fmt.Println("Product: ",products[5])
	fmt.Println("Price: ",prices[0])
    fmt.Println()
	fmt.Println("Product: ",products[2])
	fmt.Println("Price: ",prices[5])
   


}
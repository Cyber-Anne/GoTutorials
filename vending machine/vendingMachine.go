package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
)

var scanner = bufio.NewReader(os.Stdin)


var (
	A1 = 1.50
	B2 = 2.00
	C3 = 1.25
	D4 = 1.00
)


var (
	stockA1 = 5
	stockB2 = 5
	stockC3 = 5
	stockD4 = 5
)


var currentBalance = 0.0

func main() {
	exitProgram := false

	for !exitProgram {
		fmt.Println(" ======WELCOME======")
		fmt.Println("Choose an option: ")
		fmt.Println("1. Use vending machine")
		fmt.Println(" 2. Admin access")
		fmt.Println("3. Exit program")
		fmt.Println(" ===================")
		fmt.Println("Enter  Option: ")

		mainChoice := getInput()

		switch mainChoice {
		case "1":
			VendingMachine()
		case "2":
			adminAccess()
		case "3":
			fmt.Println("Exiting program")
			fmt.Println("========HAVE A WONDERFUL DAY========")
			exitProgram = true
		default:
			fmt.Println("Invalid input! Please enter 1, 2 or 3")
		}
	}
}



func getInput() string {
	input, _ := scanner.ReadString('\n')
	return strings.TrimSpace(input)
}


func getValidAmount() float64 {
	var amount float64
	valid := false

	for !valid {
		fmt.Print("Enter the amount you want you want to insert in GHS: ")
		input := getInput()

		isValid := true
		dotCount := 0
		for _, c := range input {
			if c == '.' {
				dotCount++
			} else if c < '0' || c > '9' {
				isValid = false
				break
			}
		}

		if !isValid || dotCount > 1 || input == "" {
			fmt.Println("Invalid input! Please enter numbers only.")
			continue
		}

		integerPart := 0.0
		decimalPart := 0.0
		i := 0

		for i < len(input) && input[i] != '.' {
			integerPart = integerPart*10 + float64(input[i]-'0')
			i++
		}

		if i < len(input) && input[i] == '.' {
			i++
			factor := 0.1
			for i < len(input) {
				decimalPart += float64(input[i]-'0') * factor
				factor /= 10
				i++
			}
		}

		amount = integerPart + decimalPart

		if amount <= 0 {
			fmt.Println("Amount must be greater than 0. Try again.")
		} else {
			valid = true
		}
	}

	return math.Round(amount*100) / 100
}


func getValidInteger(prompt string) int {
	var number int
	valid := false

	for !valid {
		fmt.Print(prompt)
		input := getInput()
		isValid := true

		for _, c := range input {
			if c < '0' || c > '9' {
				isValid = false
				break
			}
		}

		if !isValid || input == "" {
			fmt.Println("Invalid input! Enter a number.")
			continue
		}

		number = 0
		for _, c := range input {
			number = number*10 + int(c-'0')
		}

		if number <= 0 {
			fmt.Println("Number must be greater than 0.")
		} else {
			valid = true
		}
	}

	return number
}

func VendingMachine() {
mainLoop:
	for {
		fmt.Println("======THIS IS THE MENU======")
		fmt.Println("| CODE | ITEM NAME   | PRICE |")
		fmt.Println("|------|-------------|-------|")
		fmt.Printf(" | A1   | COKE        | %.2f  |\n", A1)
		fmt.Printf(" | B2   | FANTA       | %.2f  |\n", B2)
		fmt.Printf(" | C3   | COOKIE      | %.2f  |\n", C3)
		fmt.Printf(" | D4   | Water Bottle| %.2f  |\n", D4)

		if currentBalance <= 0 {
			currentBalance = getValidAmount()
		}

		cedis := int(currentBalance)
		pesewas := int(math.Round((currentBalance - float64(cedis)) * 100))
		fmt.Printf("Current balance: GHS %.2f\n", currentBalance)
		fmt.Printf("%d cedis\n", cedis)
		fmt.Printf("%d pesewas\n", pesewas)

		
		var code string
		validCode := false
		for !validCode {
			fmt.Print("Enter product code (A1, B2, C3, D4): ")
			code = getInput()
			if code == "A1" || code == "B2" || code == "C3" || code == "D4" {
				validCode = true
			} else {
				fmt.Println("Invalid code! Please try again.")
			}
		}

		
		var itemPrice float64
		var itemStock int
		var itemName string

		switch code {
		case "A1":
			itemPrice = A1
			itemStock = stockA1
			itemName = "COKE"
		case "B2":
			itemPrice = B2
			itemStock = stockB2
			itemName = "FANTA"
		case "C3":
			itemPrice = C3
			itemStock = stockC3
			itemName = "COOKIE"
		case "D4":
			itemPrice = D4
			itemStock = stockD4
			itemName = "WATER BOTTLE"
		}

		
		if itemStock <= 0 {
			fmt.Println("OUT OF STOCK!")
			continue
		}

		
		quantity := getValidInteger("Enter quantity: ")
		for quantity > 5 || quantity > itemStock {
			if quantity > 5 {
				fmt.Println("You cannot purchase more than 5 of this item.")
			} else {
				fmt.Println("Not enough stock available!")
			}
			quantity = getValidInteger("Enter quantity: ")
		}

		totalPrice := itemPrice * float64(quantity)

		
		for currentBalance < totalPrice {
			fmt.Printf("Insufficient funds! You need GHS %.2f more to purchase.\n", totalPrice-currentBalance)
			fmt.Println("======================")
			fmt.Println("1. Insert more money")
			fmt.Println("2. Choose a different item")
			fmt.Print("Enter choice: ")
			choiceInput := getInput()

			if choiceInput == "1" {
				addedAmount := getValidAmount()
				currentBalance += addedAmount
				currentBalance = math.Round(currentBalance*100) / 100
			} else if choiceInput == "2" {
				continue mainLoop
			} else {
				fmt.Println("Invalid choice!")
			}
		}

		
		fmt.Println("========PURCHASE SUCCESSFUL========")
		fmt.Printf("===Dispensing %d piece(s) of %s===\n", quantity, itemName)
		currentBalance -= totalPrice
		currentBalance = math.Round(currentBalance*100) / 100

		
		switch code {
		case "A1":
			stockA1 -= quantity
		case "B2":
			stockB2 -= quantity
		case "C3":
			stockC3 -= quantity
		case "D4":
			stockD4 -= quantity
		}

		cedis = int(currentBalance)
		pesewas = int(math.Round((currentBalance - float64(cedis)) * 100))
		fmt.Printf("Remaining balance: GHS %.2f\n", currentBalance)
		fmt.Printf("%d cedis\n", cedis)
		fmt.Printf("%d pesewas\n", pesewas)

	label:
		for {
			fmt.Println("=======What would you like to do next?=======")
			fmt.Println("1) Buy another item")
			fmt.Println("2) Insert more money")
			fmt.Println("3) Return change and exit")
			fmt.Print("Enter choice: ")
			menuChoiceInput := getInput()

			switch menuChoiceInput {
			case "1":
				break label
			case "2":
				add := getValidAmount()
				currentBalance += add
				currentBalance = math.Round(currentBalance*100) / 100

				cedis := int(currentBalance)
				pesewas := int(math.Round((currentBalance - float64(cedis)) * 100))
				fmt.Printf("New balance: GHS %.2f\n", currentBalance)
				fmt.Printf("%d cedis\n", cedis)
				fmt.Printf("%d pesewas\n", pesewas)
			case "3":
				cedis := int(currentBalance)
				pesewas := int(math.Round((currentBalance - float64(cedis)) * 100))
				fmt.Printf("Returning change: GHS %.2f\n", currentBalance)
				fmt.Printf("%d cedis\n", cedis)
				fmt.Printf("%d pesewas\n", pesewas)
				fmt.Println("Thank you for using the vending machine!")
				fmt.Println("Have a wonderful day!")
				fmt.Println("===========================================")
				currentBalance = 0
				return
			default:
				fmt.Println("Invalid choice!")
			}
		}
	}
}

func adminAccess() {
	attempts := 0

	for attempts < 3 {
		fmt.Print("Enter admin password: ")
		password := getInput()

		if password == "ADMIN" {
			adminAccessMenu()
			return
		}

		attempts++
		if attempts == 3 {
			fmt.Println("====Maximum attempts reached. Returning to main menu====")
		} else {
			fmt.Println("Wrong password. Try again.")
		}
	}
}

func adminAccessMenu() {
	adminMenuActive := true

	for adminMenuActive {
		fmt.Println("=======ADMIN MENU=========")
		fmt.Println("1. Restock items")
		fmt.Println("2. Change prices")
		fmt.Println("3. Exit to main menu")
		fmt.Print("Enter choice: ")
		choice := getInput()

		switch choice {
		case "1":
			stockA1 = getValidInteger("Enter quantity for COKE: ")
			stockB2 = getValidInteger("Enter quantity for FANTA: ")
			stockC3 = getValidInteger("Enter quantity for COOKIE: ")
			stockD4 = getValidInteger("Enter quantity for WATER BOTTLE: ")
			fmt.Println("======Restocked Successfully.====")
		case "2":
			fmt.Print("Enter new price for COKE: ")
			A1 = getValidAmount()
			fmt.Print("Enter new price for FANTA: ")
			B2 = getValidAmount()
			fmt.Print("Enter new price for COOKIE: ")
			C3 = getValidAmount()
			fmt.Print("Enter new price for WATER BOTTLE: ")
			D4 = getValidAmount()
			fmt.Println("========Prices Updated Successfully======")
		case "3":
			adminMenuActive = false
		default:
			fmt.Println("Invalid choice!")
		}
	}
}

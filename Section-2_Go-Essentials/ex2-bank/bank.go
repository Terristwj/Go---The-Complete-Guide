package main

import "fmt"

func main(){
	// Init
	var accountBalance float64 = 1000
	fmt.Println("Welcome to Go Bank!")

	for {
		// Loop
		printEntryMessage()
		uChoice := getUserChoice()

		// Options
		isCheckBalance := uChoice == 1
		isDeposit := uChoice == 2
		isWithdraw := uChoice == 3

		// Routes
		if isCheckBalance {
			fmt.Println("Your current balance is:", accountBalance)
		} else if isDeposit {
			accountBalance = handleDeposit(accountBalance)
		} else if isWithdraw {
			accountBalance = handleWithdraw(accountBalance)
		} else {
			break
		}
	}

	// Exit
	fmt.Println("Existing bank...")
}

func printEntryMessage(){
	fmt.Println("\nWhat do you want to do?")
	fmt.Println("1. Check balance")
	fmt.Println("2. Deposit money")
	fmt.Println("3. Withdraw money")
	fmt.Println("4. Exit")
}

func getUserChoice() int64{
	var uChoice int64
	fmt.Scan(&uChoice)
	fmt.Println("Your choice:", uChoice)
	return uChoice
}

func getUserAmount() float64{
	var uAmt float64
	fmt.Print("Enter Amount: ")
	fmt.Scan(&uAmt)
	return uAmt
}

// Route Handlers
func handleDeposit(accountBalance float64) float64{
	uAmt := getUserAmount()
	if uAmt <= 0 {
		fmt.Println("Invalid amount!")
	} else {
		accountBalance += uAmt
		fmt.Println("Your new balance is:", accountBalance)
	}
	return accountBalance
}

func handleWithdraw(accountBalance float64) float64{
	uAmt := getUserAmount()
	if uAmt <= 0 {
		fmt.Println("Invalid amount!")
	} else if uAmt > accountBalance {
		fmt.Println("Not enough money!")
	} else {
		accountBalance -= uAmt
		fmt.Println("Your new balance is:", accountBalance)
	}
	return accountBalance
}
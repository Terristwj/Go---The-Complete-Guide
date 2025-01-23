// (2) Must be in its own folder
package communication

import "fmt"

// (3) Must be capitalized
func PresentOptions() {
	fmt.Println("What do you want to do?")
	fmt.Println("1. Check balance")
	fmt.Println("2. Deposit money")
	fmt.Println("3. Withdraw money")
	fmt.Println("4. Exit")
}
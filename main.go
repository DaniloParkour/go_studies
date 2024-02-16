package main

import "fmt"

type Account struct {
	owner      string
	numAgence  int
	numAccount int
	balance    float32
}

func main() {

	account1 := Account{owner: "Danilo", numAgence: 589, numAccount: 123456, balance: 125.50}
	fmt.Println("\n", account1)

	var account2 *Account
	account2 = new(Account)
	account2.owner = "Cris"
	account2.numAccount = 112233
	account2.numAgence = 5090
	account2.balance = 9600.75

	fmt.Println("\nShow account2 first var value and second the content of pointer")
	fmt.Println(account2)
	fmt.Println(*account2)

}

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

	fmt.Println("\nShow account2: first var value and the second is the content of pointer")
	fmt.Println(account2)
	fmt.Println(*account2)

	account3 := Account{owner: "Heitor", numAgence: 589, numAccount: 123456, balance: 125500.50}
	account4 := Account{owner: "Heitor", numAgence: 589, numAccount: 123456, balance: 125500.50}
	fmt.Println("\nWe working with values and they're has the same values. Is equals =", account3 == account4)
	account4.owner = "Heitor Vidal"
	fmt.Println("\nWe working with values and they're has one different valuee. Is equals =", account3 == account4)

	//Usint refference addres, que equals'll check the addres and not the content
	var account5 *Account
	var account6 *Account
	account5 = new(Account)
	account6 = new(Account)

	account5.owner = "Helena"
	account5.numAccount = 112233
	account5.numAgence = 5090
	account5.balance = 9600.75

	account6.owner = "Helena"
	account6.numAccount = 112233
	account6.numAgence = 5090
	account6.balance = 9600.75

	fmt.Println("\nWe working with reffecrence address and they're has the same values on fields but we're comparing the refference address and the're differente. Is equals =", account5 == account6)
	fmt.Println("\nCheck below the addres of account5 and account6.")
	fmt.Println(&account5)
	fmt.Println(&account6)
	fmt.Println("\nNow we keep working with 2 diffentent refference address with the same values on fields but now we're comparing values on each addres. Is equals =", *account5 == *account6)

}

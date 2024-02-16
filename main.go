package main

import (
	"fmt"
)

type Account struct {
	owner      string
	numAgence  int
	numAccount int
	balance    float32
}

// With (c *Account) we're implementin a function that the structs Account can call
func (a *Account) withdrawn(value float32) string {
	canWithdrawn := value > 0 && value <= a.balance
	if canWithdrawn {
		a.balance -= value
		return "Withdrawn With SUCCSESS"
	} else {
		return "Invalid Value Or Balance Not Enough"
	}
}

func (a *Account) deposit(value float32) (string, float32) {
	if value > 0 {
		a.balance += value
		return "Deposit SUCCESS. Result balance is", a.balance
	}
	return "Value can't be negative. Balance keep", a.balance
}

func (fromAccount *Account) transfer(value float32, toAccount *Account) string {
	if value > 0 && fromAccount.balance >= value {
		fromAccount.withdrawn(value)
		toAccount.deposit(value)
		return "SUCCESS!"
	}
	return "ERROR!"
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

	fmt.Println(account1.withdrawn(45.))

	fmt.Println("\nNow We Use A Deposit Function That Returns 2 Values")
	message, value := account2.deposit(250)
	fmt.Println(message, value)

	fmt.Println("\nFunction Transfer")
	fmt.Println("Balance account 1 is ", account1.balance)
	fmt.Println("Balance account 2 is ", account2.balance)
	fmt.Println("Transfer 400 from account 2 to account 1")
	fmt.Println("\n\n", account2.transfer(400, &account1))
	fmt.Println("\nNew balance of account are")
	fmt.Println("Balance account 1 is ", account1.balance)
	fmt.Println("Balance account 2 is ", account2.balance)

}

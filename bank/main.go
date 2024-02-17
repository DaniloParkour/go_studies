package main

import (
	"fmt"
	"my_projects/alura/studies/bank/accounts"
	"my_projects/alura/studies/bank/clients"
)

func main() {

	account1 := accounts.Account{Owner: clients.Owner{Name: "Danilo", CPF: "111.222.333-44", Profession: "Engenheiro de Software"}, NumAgence: 589, NumAccount: 123456, Balance: 125.50}
	fmt.Println("\n", account1)

	var account2 *accounts.Account
	account2 = new(accounts.Account)
	account2.Owner = clients.Owner{Name: "Cris", CPF: "111.222.333-44", Profession: "Engenheiro de Software"}
	account2.NumAccount = 112233
	account2.NumAgence = 5090
	account2.Balance = 9600.75

	fmt.Println("\nShow account2: first var value and the second is the content of pointer")
	fmt.Println(account2)
	fmt.Println(*account2)

	account3 := accounts.Account{Owner: clients.Owner{Name: "Cris", CPF: "555.777.777-88", Profession: "UI/UX Designer"}, NumAgence: 589, NumAccount: 123456, Balance: 125500.50}
	account4 := accounts.Account{Owner: clients.Owner{Name: "Heitor", CPF: "999.999.999-00", Profession: "Estudante"}, NumAgence: 589, NumAccount: 123456, Balance: 125500.50}
	fmt.Println("\nWe working with values and they're has the same values. Is equals =", account3 == account4)
	account4.Owner = clients.Owner{Name: "Heitor Vidal", CPF: "111.222.333-44", Profession: "Engenheiro de Software"}
	fmt.Println("\nWe working with values and they're has one different valuee. Is equals =", account3 == account4)

	//Usint refference addres, que equals'll check the addres and not the content
	var account5 *accounts.Account
	var account6 *accounts.Account
	account5 = new(accounts.Account)
	account6 = new(accounts.Account)

	account5.Owner = clients.Owner{Name: "Helena", CPF: "111.222.333-44", Profession: "Engenheiro de Software"}
	account5.NumAccount = 112233
	account5.NumAgence = 5090
	account5.Balance = 9600.75

	account6.Owner = clients.Owner{Name: "Helena", CPF: "111.222.333-44", Profession: "Engenheiro de Software"}
	account6.NumAccount = 112233
	account6.NumAgence = 5090
	account6.Balance = 9600.75

	fmt.Println("\nWe working with reffecrence address and they're has the same values on fields but we're comparing the refference address and the're differente. Is equals =", account5 == account6)
	fmt.Println("\nCheck below the addres of account5 and account6.")
	fmt.Println(&account5)
	fmt.Println(&account6)
	fmt.Println("\nNow we keep working with 2 diffentent refference address with the same values on fields but now we're comparing values on each addres. Is equals =", *account5 == *account6)

	fmt.Println(account1.Withdrawn(45.))

	fmt.Println("\nNow We Use A Deposit Function That Returns 2 Values")
	message, value := account2.Deposit(250)
	fmt.Println(message, value)

	fmt.Println("\nFunction Transfer")
	fmt.Println("Balance account 1 is ", account1.Balance)
	fmt.Println("Balance account 2 is ", account2.Balance)
	fmt.Println("Transfer 400 from account 2 to account 1")
	fmt.Println("\n\n", account2.Transfer(400, &account1))
	fmt.Println("\nNew Balance of account are")
	fmt.Println("Balance account 1 is ", account1.Balance)
	fmt.Println("Balance account 2 is ", account2.Balance)

}

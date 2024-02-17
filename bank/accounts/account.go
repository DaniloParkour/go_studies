package accounts

import "my_projects/alura/studies/bank/clients"

type Account struct {
	//Upcase First Letter For Public Visibility
	Owner      clients.Owner
	NumAgence  int
	NumAccount int
	balance    float32
}

// With (c *Account) we're implementin a function that the structs Account can call
func (a *Account) Withdrawn(value float32) string {
	canWithdrawn := value > 0 && value <= a.balance
	if canWithdrawn {
		a.balance -= value
		return "Withdrawn With SUCCSESS"
	} else {
		return "Invalid Value Or Balance Not Enough"
	}
}

func (a *Account) Deposit(value float32) (string, float32) {
	if value > 0 {
		a.balance += value
		return "Deposit SUCCESS. Result balance is", a.balance
	}
	return "Value can't be negative. Balance keep", a.balance
}

func (fromAccount *Account) Transfer(value float32, toAccount *Account) string {
	if value > 0 && fromAccount.balance >= value {
		fromAccount.Withdrawn(value)
		toAccount.Deposit(value)
		return "SUCCESS!"
	}
	return "ERROR!"
}

func (a *Account) GetBalance() float32 {
	return a.balance
}

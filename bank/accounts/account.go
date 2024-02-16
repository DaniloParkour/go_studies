package accounts

type Account struct {
	//Upcase First Letter For Public Visibility
	Owner      string
	NumAgence  int
	NumAccount int
	Balance    float32
}

// With (c *Account) we're implementin a function that the structs Account can call
func (a *Account) Withdrawn(value float32) string {
	canWithdrawn := value > 0 && value <= a.Balance
	if canWithdrawn {
		a.Balance -= value
		return "Withdrawn With SUCCSESS"
	} else {
		return "Invalid Value Or Balance Not Enough"
	}
}

func (a *Account) Deposit(value float32) (string, float32) {
	if value > 0 {
		a.Balance += value
		return "Deposit SUCCESS. Result balance is", a.Balance
	}
	return "Value can't be negative. Balance keep", a.Balance
}

func (fromAccount *Account) Transfer(value float32, toAccount *Account) string {
	if value > 0 && fromAccount.Balance >= value {
		fromAccount.Withdrawn(value)
		toAccount.Deposit(value)
		return "SUCCESS!"
	}
	return "ERROR!"
}

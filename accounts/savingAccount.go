package accounts

//é possivel definir um apelido escrevendo qualquer coisa na frente do import
import "go_projects/banking_go/customers"

type SavingAccount struct {
	Holder                                 customers.Holder
	AgencyNumber, AccountNumber, Operation int
	balance                                float64
}

func (c *SavingAccount) Withdraw(withdrawalAmount float64) (string, float64) {
	if c.balance >= withdrawalAmount && withdrawalAmount > 0 {
		c.balance -= withdrawalAmount
		return "Saque feito com sucesso!", c.balance
	} else {
		return "Saldo insuficiente!", c.balance
	}
}

func (c *SavingAccount) DepositAmount(depositAmount float64) (string, float64) {
	if depositAmount > 0 {
		c.balance += depositAmount
		return "Depósito efetuado com sucesso!", c.balance
	} else {
		return "Não foi possível depositar!", c.balance
	}
}

func (c *SavingAccount) GetBalance() float64 {
	return c.balance
}

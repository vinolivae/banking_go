package accounts

import "go_projects/banking_go/customers"

//Letras maiusculas também definem campos publicos de uma struct
type CurrentAccount struct {
	Holder                      customers.Holder
	AgencyNumber, AccountNumber int
	balance                     float64
}

/*
A notação (c *CurrentAccount) define um método em uma estrutura.
(c *CurrentAccount) define um receptor de método. "c" é uma variável que representa uma instância
da estrutura "CurrentAccount" e "*CurrentAccount" significa que "c" é um ponteiro para uma
instância de "CurrentAccount". Permitindo, assim, que a instância da estrutura seja modificada.
*/
func (c *CurrentAccount) Withdraw(withdrawalAmount float64) (string, float64) {
	if c.balance >= withdrawalAmount && withdrawalAmount > 0 {
		c.balance -= withdrawalAmount
		return "Saque feito com sucesso!", c.balance
	} else {
		return "Saldo insuficiente!", c.balance
	}
}

func (c *CurrentAccount) DepositAmount(depositAmount float64) (string, float64) {
	if depositAmount > 0 {
		c.balance += depositAmount
		return "Depósito efetuado com sucesso!", c.balance
	} else {
		return "Não foi possível depositar!", c.balance
	}
}

/*
O toAccount *CurrentAccount aponta para a conta que recebe a transferência, porém, na decalaração
da função é necessário utilizar o "&" para expor o endereço da conta.
*/
func (c *CurrentAccount) Transfer(transferAmount float64, toAccount *CurrentAccount) bool {
	if transferAmount <= c.balance && transferAmount > 0 {
		c.balance -= transferAmount
		toAccount.balance += transferAmount
		return true
	} else {
		return false
	}
}

func (c *CurrentAccount) GetBalance() float64 {
	return c.balance
}

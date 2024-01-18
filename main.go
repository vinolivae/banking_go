package main

import (
	"fmt"
	"strconv"
)

type CurrentAccount struct {
	holder        string
	agencyNumber  int
	accountNumber int
	balance       float64
}

func main() {
	viniciusCurrentAccount := CurrentAccount{holder: "vinicius", agencyNumber: 589, accountNumber: 123456, balance: 125.50}

	fmt.Println("Saldo da conta é", viniciusCurrentAccount.balance)
	fmt.Println(viniciusCurrentAccount.withdraw(20))
	fmt.Println("O saldo agora é", viniciusCurrentAccount.balance)
}

/*
A notação (c *CurrentAccount) define um método em uma estrutura.
(c *CurrentAccount) define um receptor de método. "c" é uma variável que representa uma instância
da estrutura "CurrentAccount" e "*CurrentAccount" significa que "c" é um ponteiro para uma
instância de "CurrentAccount". Permitindo, assim, que a instância da estrutura seja modificada.
*/
func (c *CurrentAccount) withdraw(withdrawalAmount float64) string {
	if validateWithdraw(c.balance, withdrawalAmount) {
		c.balance -= withdrawalAmount
		return "Saque de " + strconv.FormatFloat(withdrawalAmount, 'f', 2, 64) + " feito com sucesso!"
	} else {
		return "Saldo insuficiente!"
	}
}

func validateWithdraw(balance, withdrawalAmount float64) bool {
	canWithdraw := balance >= withdrawalAmount && withdrawalAmount > 0
	return canWithdraw
}

package main

import (
	"fmt"
	"go_projects/banking_go/accountGenerator"
)

/*
Módulos que possuem a função "withdraw" que recebe um float64 e retorna uma string
assina um contrato com essa interface.

O uso dessa interface contempla as regras das contas que possuem a função withdraw.

OBS..: As interfaces são implementadas de forma implicita. E o objetivo é reaproveitar
métoros ou funções entre tipos diferentes.
*/
type VerifyAccount interface {
	Withdraw(amount float64) (string, float64)
}

func main() {
	joaoCurrentAccount := accountGenerator.CreateCurrentAccount("João")
	joaoCurrentAccount.DepositAmount(100)

	mariaCurrentAccount := accountGenerator.CreateCurrentAccount("Maria")
	mariaCurrentAccount.DepositAmount(200)

	luizSavingAccount := accountGenerator.CreateSavingAccount("Luiz")
	luizSavingAccount.DepositAmount(300)

	anaSavingAccount := accountGenerator.CreateSavingAccount("Ana")
	anaSavingAccount.DepositAmount(400)

	//210
	//90
	joaoCurrentAccount.Transfer(10, &mariaCurrentAccount)
	fmt.Println("Saldo da Maria:", mariaCurrentAccount.GetBalance())
	fmt.Println("Saldo do João:", joaoCurrentAccount.GetBalance())
	fmt.Println()

	//100
	//200
	mariaCurrentAccount.Transfer(10, &joaoCurrentAccount)
	fmt.Println("Saldo do João:", joaoCurrentAccount.GetBalance())
	fmt.Println("Saldo da Maria:", mariaCurrentAccount.GetBalance())
	fmt.Println()

	//90
	joaoCurrentAccount.Withdraw(10)
	fmt.Println("Saldo do João:", joaoCurrentAccount.GetBalance())
	fmt.Println()

	//190
	mariaCurrentAccount.Withdraw(10)
	fmt.Println("Saldo da Maria:", mariaCurrentAccount.GetBalance())
	fmt.Println()

	//290
	luizSavingAccount.Withdraw(10)
	fmt.Println("Saldo do Luiz:", luizSavingAccount.GetBalance())
	fmt.Println()

	//390
	anaSavingAccount.Withdraw(10)
	fmt.Println("Saldo da Ana:", anaSavingAccount.GetBalance())
	fmt.Println()

	//70
	payInvoice(&joaoCurrentAccount, 20)
	fmt.Println("Saldo do João:", joaoCurrentAccount.GetBalance())
	fmt.Println()

	// 160
	payInvoice(&mariaCurrentAccount, 30)
	fmt.Println("Saldo da Maria:", mariaCurrentAccount.GetBalance())
	fmt.Println()

	// 250
	payInvoice(&luizSavingAccount, 40)
	fmt.Println("Saldo do Luiz:", luizSavingAccount.GetBalance())
	fmt.Println()

	// 340
	payInvoice(&anaSavingAccount, 50)
	fmt.Println("Saldo da Ana:", anaSavingAccount.GetBalance())
	fmt.Println()
}

/*
Para pagar boleto não precisamos identificar uma conta.
A interface "verifyAccount" nos ajuda a executar essa ação para qualquer tipo de conta, contanto
que ela possua a função "withdraw".

A conta passada no parametro da função precisa conter o "&", assim o endereço será exposto para
o ponteiro na função "withdraw". Ou definir um ponteiro aqui. e.g.: (v *VerifyAccount)
*/
func payInvoice(account VerifyAccount, invoiceAmount float64) {
	account.Withdraw(invoiceAmount)
}

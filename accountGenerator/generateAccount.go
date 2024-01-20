package accountGenerator

import (
	a "go_projects/banking_go/accounts"
	"go_projects/banking_go/customers"
	"math/rand"
	"time"
)

func CreateSavingAccount(holderName string) a.SavingAccount {
	holderData := customers.Holder{Name: holderName, Document: "123.456.789-01", Profession: profession()}
	holderAccount := a.SavingAccount{Holder: holderData, AgencyNumber: agencyNumber(), AccountNumber: accountNumber()}
	return holderAccount
}

func CreateCurrentAccount(holderName string) a.CurrentAccount {
	holderData := customers.Holder{Name: holderName, Document: "123.456.789-01", Profession: profession()}
	holderAccount := a.CurrentAccount{Holder: holderData, AgencyNumber: agencyNumber(), AccountNumber: accountNumber()}
	return holderAccount
}

func profession() string {
	rand.Seed(time.Now().UnixNano())
	professions := []string{"carpinteiro", "pedreiro", "advogado", "programador", "juiz", "streamer"}
	random_index := rand.Intn(len(professions))
	random_profession := professions[random_index]

	return random_profession
}

func accountNumber() int {
	return pickANumber(1234, 9876)
}

func agencyNumber() int {
	return pickANumber(123456, 987654)
}

func pickANumber(min int, max int) int {
	rand.Seed(time.Now().UnixNano())
	random_number := rand.Intn(max-min+1) + min
	return random_number
}

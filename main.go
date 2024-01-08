package main

import (
	"fmt"
	"os"
	"time"

	"github.com/mitch000001/go-hbci"
	"github.com/mitch000001/go-hbci/client"
	"github.com/mitch000001/go-hbci/domain"
	"github.com/mitch000001/go-hbci/iban"
)

var testAccount domain.AccountConnection
var sepaTestAccount domain.InternationalAccountConnection

func main() {

	hbci.SetDebugMode(true)
	hbci.SetInfoLog(true)

	config := client.Config{
		AccountID:          os.Getenv("GOHBCI_USERID"),
		BankID:             os.Getenv("GOHBCI_BLZ"),
		PIN:                os.Getenv("GOHBCI_PIN"),
		EnableDebugLogging: true,
	}

	testAccount = domain.AccountConnection{AccountID: config.AccountID, CountryCode: 280, BankID: config.BankID}

	i, err := iban.NewGerman(config.BankID, config.AccountID)
	if err != nil {
		panic(err)
	}

	sepaTestAccount = domain.InternationalAccountConnection{
		IBAN:      string(i),
		AccountID: config.AccountID,
		BankID:    domain.BankID{CountryCode: 280, ID: config.BankID},
	}

	c, err := client.New(config)
	if err != nil {
		panic(err)
	}

	timeframe := domain.Timeframe{
		StartDate: domain.NewShortDate(time.Now().AddDate(0, 0, -90)),
	}

	// Get Accounts
	accounts, _ := c.Accounts()
	fmt.Printf("Accounts: %s\n", accounts)

	// Get Transactions
	if len(accounts) > 0 {
		transactions, err := c.AccountTransactions(accounts[0].AccountConnection, timeframe, false, "")

		if err != nil {
			fmt.Printf("AccountTransactions: %T:%v\n", err, err)
		}

		for _, tr := range transactions {
			fmt.Printf("Transaction: %s\n", tr)
		}
	}

	// Get SepaTransactions
	sepaTransactions, err := c.SepaAccountTransactions(sepaTestAccount, timeframe, false, "")

	fmt.Println("-------------------------------- SepaAccountTransactions -----------------------------")

	if err != nil {
		fmt.Printf("SepaAccountTransactions: %T:%v\n", err, err)
	}

	for _, tr := range sepaTransactions {
		fmt.Printf("Transaction: %s\n", tr)
	}

	// Get SepaAccountBalances
	balances, err := c.SepaAccountBalances(sepaTestAccount, false, "")

	fmt.Println("-------------------------------- SepaAccountBalances -----------------------------")

	if err != nil {
		fmt.Printf("SepaAccountBalances: %T:%v\n", err, err)
	}
	for _, tr := range balances {
		fmt.Printf("Balances: %s\n", tr)
	}

}

package main

import (
	"fmt"
	"github.com/Lavoaster/monzo-go"
)

func main() {
	// TODO: Get this from somewhere else like ENV
	accessToken := ""

	client := monzo.New(accessToken)

	metadata := make(map[string]string)
	metadata["notes"] = "#dinner"

	//client.AnnotateTransaction("", metadata)

	fmt.Println(client.WhoAmI())

	//
	//accounts := client.GetAccounts()
	//
	//for _, account := range accounts {
	//	if account.Closed {
	//		continue
	//	}
	//
	//	//balance := client.GetBalance(account.Id)
	//	//fmt.Println("")
	//	//fmt.Println("------------")
	//	//fmt.Println("Account ID: " + account.Id)
	//	//fmt.Println("Account Description: " + account.Description)
	//	//fmt.Println("Account Creation Date: " + account.Created.Format(time.RFC3339))
	//	//fmt.Println("Account Balance: " + strconv.Itoa(int(balance.Balance)))
	//	//fmt.Println("Account Currency: " + strconv.Itoa(int(balance.Balance)))
	//	//fmt.Println("Account Spend Today: " + strconv.Itoa(int(balance.SpendToday)))
	//	//fmt.Println("------------")
	//
	//
	//}
}

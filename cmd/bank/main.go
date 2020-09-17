package main

import (
	"bank/pkg/bank/card"
	"bank/pkg/bank/types"
	"fmt"
)

func main() {
	result := types.Card{Balance: 0, Active: false}
	card.Withdraw(&result, 10000)
	fmt.Printf("%+v", result)
	fmt.Printf("%+v", card.IssueCard(types.TJS, "black", "bla bla bla"))
}
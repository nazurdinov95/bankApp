package card

import (
	"bank/pkg/bank/types"
	"fmt"
)

func ExampleWithdraw_positive() {
	card := types.Card{Balance: 20_000_00, Active: true}
	Withdraw(&card, 10_000_00)

	fmt.Println(card.Balance)
	// Output: 1000000
}

func ExampleWithdraw_noMoney() {
	card := types.Card{Balance: 0, Active: true}
	Withdraw(&card, 10_000_00)
	
	fmt.Println(card.Balance)
	// Output: 0
}

func ExampleWithdraw_inactive() {
	card := types.Card{Balance: 0, Active: false}
	Withdraw(&card, 10_000_00)
	
	fmt.Println(card.Balance)
	// Output: 0
}

func ExampleWithdraw_limit() {
	card := types.Card{Balance: 100000, Active: false}
	Withdraw(&card, 2000005)
	
	fmt.Println(card.Balance)
	// Output: 100000
}

func ExampleDeposit_normal() {
	card := types.Card{Balance: 100000, Active: true}
	Deposit(&card, 100000)

	fmt.Println(card.Balance)
	// Output: 200000
}

func ExampleDeposit_noActive() {
	card := types.Card{Balance: 100000, Active: false}
	Deposit(&card, 100000)

	fmt.Println(card.Balance)
	// Output: 100000
}

func ExampleDeposit_moThenLimit() {
	card := types.Card{Balance: 100000, Active: true}
	Deposit(&card, 6000000)

	fmt.Println(card.Balance)
	// Output: 100000
}


func ExampleAddBonus_normal() {
	card := types.Card{Balance: 1000,  MinBalance: 1000000, Active: true}
	AddBonus(&card, 3, 30, 365)

	fmt.Println(card.Balance)
	// Output: 3465
}

func ExampleAddBonus_noActive() {
	card := types.Card{Balance: 1000, MinBalance: 1000000, Active: false}
	AddBonus(&card, 3, 30, 365)

	fmt.Println(card.Balance)
	// Output: 1000
}

func ExampleAddBonus_negative() {
	card := types.Card{Balance: -1000, MinBalance: 1000000, Active: true}
	AddBonus(&card, 3, 30, 365)

	fmt.Println(card.Balance)
	// Output: -1000
}

func ExampleAddBonus_moreThenLimit() {
	card := types.Card{Balance: 1000, MinBalance: 1000000000, Active: true}
	AddBonus(&card, 3, 30, 365)

	fmt.Println(card.Balance)
	// Output: 1000
}

func ExampleTotal()  {
	cards := []types.Card{
		{
			Balance: 10_000_00,
			Active: true,
		},
		{
			Balance: 10_000_00,
			Active: true,
		},
		{
			Balance: 10_000_00,
			Active: true,
		},
	}

	fmt.Println(Total(cards))
	// Output: 3000000
}
package card

import (
	"bank/pkg/bank/types"
)

const withdrawLimit = 2000000
const depositLimit = 5000000
const bonusLimit = 500000

// IssueCard creates card
func IssueCard(currency types.Currency, color string, name string) types.Card {
	return types.Card{
		ID:       1000,
		PAN:      "5058 xxxx xxxx 0001",
		Balance:  0,
		Currency: currency,
		Color:    color,
		Name:     name,
		Active:   true,
	}
}

// Withdraw withdrawins an amount from a card
func Withdraw(card *types.Card, amount types.Money) {
	if amount < 0 {
		return
	}

	if amount > withdrawLimit {
		return
	}

	if !card.Active {
		return
	}

	if card.Balance < amount {
		return
	}

	card.Balance -= amount
}

// Deposit сredits funds to the card.
func Deposit(card *types.Card, amount types.Money) {
	if amount < 0 {
		return
	}

	if amount > depositLimit {
		return
	}

	if !card.Active {
		return
	}

	card.Balance += amount
}

// AddBonus add bonus to card
func AddBonus(card *types.Card, percent int, daysInMonth int, daysInYear int) {
	if !card.Active {
		return
	}

	if card.Balance <= 0 {
		return
	}

	bonus := int(card.MinBalance) * percent / 100 * daysInMonth / daysInYear

	if bonus > bonusLimit || bonus < 0 {
		return
	}

	card.Balance += types.Money(bonus)
}

// Total calculates the total on all cards.
func Total(cards []types.Card) types.Money {
	total := types.Money(0)
	for _, card := range cards {
		if card.Active && card.Balance >= 0 {
			total += card.Balance
		}
	}

	return total
}


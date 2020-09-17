package types

// Money represents a monetary amount in minimum units (cents, kopecks, diramas, etc.).
type Money int64

// Currency presents currency code.
type Currency string

// Currency codes
const (
	TJS Currency = "TJS"
	RUB Currency = "RUB"
	USD Currency = "USD"
)

// PAN presents number of card.
type PAN string

// Card presents information about card.
type Card struct {
	ID       int
	PAN      PAN
	Balance  Money
	MinBalance Money
	Currency Currency
	Color    string
	Name     string
	Active   bool
}

// Payment presents information about the payment.
type Payment struct {
	ID int
	Amount Money
}
package payment

import (
	"bank/pkg/bank/types"
	"fmt"
)

func ExampleMax()  {
	payments := []types.Payment {
		{
			ID: 1,
			Amount: 1000,
		},
		{
			ID: 2,
			Amount: 3000,
		},
		{
			ID: 3,
			Amount: 1000,
		},
	}

	max := Max(payments)
	fmt.Println(max.ID)
	// Output: 2
}
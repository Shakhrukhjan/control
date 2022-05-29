package card

import (
	"control/pkg/types"
	"fmt"
)

func ExampleChangeActiveCard() {
	temps := []*types.Card{
		{
			Balance:  16_000,
			Currency: "TJS",
			Activity: false,
		},
		{
			Balance:  17_000,
			Currency: "RUB",
			Activity: true,
		},
		{
			Balance:  18_000,
			Currency: "EUR",
			Activity: false,
		},
	}
	ChangeActiveCard(temps)
	for i := 0; i < len(temps); i++ {
		fmt.Println(*&temps[i].Activity)
	}

	// Output:
	// true
	// true
	// true
}

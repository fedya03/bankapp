package card

import (
	"bank/pkg/bank/types"
	"fmt"
	//"fmt"
)

func ExamplePaymentSources() {
	cards := []types.Card{
		{
			Balance:  888_88,
			Active:   true,
			PAN:      "5058 xxxx xxxx 8888",
			Currency: types.TJS,
		},
		{
			Balance:  999_99,
			Active:   true,
			PAN:      "5058 xxxx xxxx 9999",
			Currency: types.TJS,
		},
	}
	sources := PaymentSources(cards)
	for _, card := range sources {
		fmt.Println(card.Number)
	}

	// Output:
	// 5058 xxxx xxxx 8888
	// 5058 xxxx xxxx 9999
}

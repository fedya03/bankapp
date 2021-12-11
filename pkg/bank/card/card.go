package card

import "bank/pkg/bank/types"

func PaymentSources(cards []types.Card) []types.PaymentSource {
	arrBalance := []types.PaymentSource{}
	for _, card := range cards {
		if card.Balance > 0 && card.Active {
			arrBalance = append(arrBalance, types.PaymentSource{
				Balance: card.Balance,
				Number:  string(card.PAN),
				Type:    "card",
			})
		}
	}
	return arrBalance
}

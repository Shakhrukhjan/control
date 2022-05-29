package card

import (
	"control/pkg/types"
)

func ChangeActiveCard(cards []*types.Card) {
	for _, card := range cards {
		if card.Activity == false {
			card.Activity = true 
		}
	}
}

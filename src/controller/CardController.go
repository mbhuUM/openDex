package CardController

import (
    "dex/src/model" 
)
type cardInterface interface {
	addCard(*Card.card) (void)
	removeCard(*Card.card) (void)
	searchCard(string) (*Card.card)

}
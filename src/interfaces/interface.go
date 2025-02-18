package interfaces

import (
    "dex/src/model" 
)

type CardInterface interface {
	CreateCard(card card.Card)
	RemoveCard(card card.Card)
	SearchCardByPokemonName(query string) (card card.Card)
	DeleteCard(card card.Card)
}
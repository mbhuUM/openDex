package card


//              INSERT INTO pokemon_cards (name, set_name, card_number, rarity, card_type)
type Card struct{
    Name        string    `json:"name"`
    Set         string    `json:"set"`
	CardNumber 	string 	  `json:"card_number"`
    Rarity      string    `json:"rarity"`
	Type 		string    `json:"card_type"`
	
}
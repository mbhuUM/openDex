package card

type card struct{
	ID          string    `json:"id"`
    Name        string    `json:"name"`
    Set         string    `json:"set"`
    Number      string    `json:"number"`
    ImageURL    string    `json:"image_url"`
    owner       string    // private field
}
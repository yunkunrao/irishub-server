package document

type Coins []Coin

type Coin struct {
	Denom  string  `bson:"denom"`
	Amount float64 `bson:"amount"`
}

// validator description
type ValDescription struct {
	Moniker  string `json:"moniker" bson:"moniker"`
	Identity string `json:"identity" bson:"identity"`
	Website  string `json:"website" bson:"website"`
	Details  string `json:"details" bson:"details"`
}

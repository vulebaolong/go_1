package model

type Expense struct {
	Id     int    `json:"id"`
	Amount int    `json:"amount"`
	Note   string `json:"note"`
}

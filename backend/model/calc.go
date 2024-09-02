package model

type OldCalc struct {
	AccountId      int         `json:"account_id" binding:"required"`
	EventId        int         `json:"event_id" binding:"required"`
	AccountAmounts map[int]int `json:"account_amounts" binding:"required"`
}
type Calc struct {
	EventId        int             `json:"event_id" binding:"required"`
	AccountAmounts []AccountAmount `json:"account_amounts" binding:"required"`
}
type AccountAmount struct {
	AccountId int         `json:"account_id" binding:"required"`
	Amount    map[int]int `json:"account_amount" binding:"required"`
}

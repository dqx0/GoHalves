package model

type Calc struct {
	AccountId      int         `json:"account_id" binding:"required"`
	EventId        int         `json:"event_id" binding:"required"`
	AccountAmounts map[int]int `json:"account_amounts" binding:"required"`
}

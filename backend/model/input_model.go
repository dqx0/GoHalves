package model

type InputAccount struct {
	UserID   string `json:"user_id" binding:"required"`
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email"`
	Password string `json:"password" binding:"required"`
}
type InputEvent struct {
	UserId      int    `json:"user_id" binding:"required"`
	Title       string `json:"title" binding:"required"`
	Description string `json:"description" binding:"reqiured"`
}
type InputPay struct {
	PaidUserID uint `json:"paid_user_id" binding:"required"`
	EventID    uint `json:"event_id" binding:"required"`
	Amount     uint `json:"amount" binding:"required"`
}
type InputAccountEvent struct {
	AccountID   uint `json:"account_id" binding:"required"`
	EventID     uint `json:"event_id" binding:"required"`
	AuthorityID uint `json:"authority_id" binding:"required"`
}
type InputAccountPay struct {
	AccountID uint `json:"account_id" binding:"required"`
	PayID     uint `json:"pay_id" binding:"required"`
}
type InputFriend struct {
	SendAccountID     string `json:"send_account_id" binding:"required"`
	ReceivedAccountID string `json:"received_account_id" binding:"required"`
}
